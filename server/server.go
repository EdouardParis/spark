package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/edouardparis/spark/views"
)

type key int

var healthy int32

const (
	requestIDKey key = 0
)

type Server struct {
	http.Server
}

func (s *Server) Run() {
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		s.ErrorLog.Println("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		s.SetKeepAlivesEnabled(false)
		err := s.Shutdown(ctx)
		if err != nil {
			s.ErrorLog.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	s.ErrorLog.Println("Server is ready to handle requests at", s.Addr)
	atomic.StoreInt32(&healthy, 1)
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.ErrorLog.Fatalf("Could not listen on %s: %v\n", s.Addr, err)
	}

	<-done
	s.ErrorLog.Println("Server stopped")

}

func New(listenAddr string) *Server {
	router := http.NewServeMux()
	router.Handle("/", views.Index())
	router.Handle("/healthcheck", healthcheck())
	router.Handle("/api/v1/charges", views.Charges())

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	return &Server{
		http.Server{
			Addr:     listenAddr,
			Handler:  tracing(nextRequestID)(logging(logger)(router)),
			ErrorLog: logger,
		},
	}
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func healthcheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

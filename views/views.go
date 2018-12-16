package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edouardparis/spark/payloads"
	"github.com/edouardparis/spark/resources"
)

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello, World!")
	})
}

func Charges() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createCharge(w, r)
			return
		case http.MethodGet:
			listCharge(w, r)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
	})
}

func createCharge(w http.ResponseWriter, r *http.Request) {
	payload := &payloads.Charge{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !payload.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	charge := newCharge(payload)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(charge)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func listCharge(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, you!")
}

func newCharge(payload *payloads.Charge) *resources.Charge {
	return &resources.Charge{}
}

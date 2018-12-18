package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/edouardparis/spark/payloads"
	"github.com/edouardparis/spark/resources"
	"github.com/edouardparis/spark/store"
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payload := payloads.Charge{}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !payload.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	charge := resources.NewCharge(&payload)
	store.InsertCharge(charge)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(charge)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func listCharge(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		page   = 0
		size   = 30
		params = r.URL.Query()
	)

	pageStr := params.Get("page")
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	sizeStr := params.Get("size")
	if sizeStr != "" {
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	list := store.ListCharges(page, size)
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

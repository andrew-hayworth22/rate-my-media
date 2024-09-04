package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrew-hayworth22/rate-my-media/business"
)

func ValidationError(w http.ResponseWriter, validationError business.ValidationError) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(validationError)
	if err != nil {
		InternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(data)
}

func InternalError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	w.Write([]byte("An unexpected error occurred"))
	fmt.Printf("internal error: %s", err)
}

func Success(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		InternalError(w, err)
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}

package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Encode[T any](w http.ResponseWriter, status int, data T) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		EncodeInternalError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
}

func EncodeInternalError(w http.ResponseWriter) {
	Encode(w, http.StatusInternalServerError, "An internal error has occurred")
}

func EncodeValidationError(w http.ResponseWriter, problems map[string]string) {
	Encode(w, http.StatusBadRequest, problems)
}

func Decode[T any](r *http.Request) (T, error) {
	var result T
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("decode error: %w", err)
	}

	return result, nil
}

func DecodeValid[T Validator](r *http.Request, data T) (T, map[string]string, error) {
	var result T
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		return result, nil, fmt.Errorf("decode error: %w", err)
	}
	if problems := result.Valid(r.Context()); len(problems) > 0 {
		return result, problems, fmt.Errorf("invalid %T: %d problems", result, len(problems))
	}
	return result, nil, nil
}

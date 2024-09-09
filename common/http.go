package common

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct{
	Status int `json:"status"`
	Error error `json:"error"`
}
func WriteJson(w http.ResponseWriter, v any) {
	err := json.NewEncoder(w).Encode(v)

	if err != nil {
		WriteStatus(w, 500)
	}
}

func WriteStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteStatus(w, status)
	e := &errorResponse{
		status, err,
	}

	json.NewEncoder(w).Encode(e)
}

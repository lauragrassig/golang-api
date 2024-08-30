package common

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v any) {
	err := json.NewEncoder(w).Encode(v)

	if err != nil {
		WriteStatus(w, 500)
	}
}

func WriteStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
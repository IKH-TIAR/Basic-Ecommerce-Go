package utils

import "net/http"

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, status int, error string) {
	WriteJSON(w, status, ErrorResponse{Error: error})
}
package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type HttpRequestHandler interface {
	RegisterRoutes(router *mux.Router)
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func Error(w http.ResponseWriter, error *ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(error.Code)
	json.NewEncoder(w).Encode(error)
}

func NotFound(w http.ResponseWriter, subject string) {
	Error(w, &ErrorResponse{Code: http.StatusNotFound, Error: fmt.Sprintf("%v not found", subject)})
}

func InternalServerError(w http.ResponseWriter) {
	Error(w, &ErrorResponse{Code: http.StatusInternalServerError, Error: http.StatusText(http.StatusInternalServerError)})
}
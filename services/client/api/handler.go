package api

import "github.com/gorilla/mux"

type HttpRequestHandler interface {
	RegisterRoutes(router *mux.Router)
}
package api

import (
	"client/services"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ClientServer struct {
	port     int
	server   *http.Server
	handlers []HttpRequestHandler
}

func NewClientServer(config *services.Config, handlers []HttpRequestHandler) *ClientServer {
	server := &ClientServer{
		port:     config.ServerPort,
		handlers: handlers,
	}
	return server
}

func (s *ClientServer) HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	for _, handler := range s.handlers {
		handler.RegisterRoutes(router)
	}

	s.server = &http.Server{Addr: fmt.Sprintf(":%v", s.port), Handler: router}
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start up: %s\n", err)
		}
	}()
}

func (s *ClientServer) Shutdown(ctx context.Context) {
	// Shutdown the server gracefully
	log.Print("Server is shutting down")
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}
	log.Print("Server was shut down")
}
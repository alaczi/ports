package main

import (
	"context"
	"encoding/json"
	"fmt"
	repo "github.com/alaczi/ports/repository"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	ServerPort      int    `default:"8080"`
	PortServiceAddr string `default:"localhost:50051"`
	DataFile        string `default:"./data/ports.json"`
}

type ClientService struct {
	port       int
	repository repo.PortRepository
	server     *http.Server
}

func newClientService(config *Config) *ClientService {
	service := &ClientService{
		repository: newGRPCPortRepository(&config.PortServiceAddr, &config.DataFile),
		port:       config.ServerPort,
	}
	return service
}

func (s ClientService) HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/ports/{id}", s.getPort)
	//	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", s.port), myRouter))
	s.server = &http.Server{Addr: fmt.Sprintf(":%v", s.port), Handler: myRouter}
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start up: %s\n", err)
		}
	}()
}

func (s ClientService) Shutdown(ctx context.Context) {
	// Shutdown the server gracefully
	if s.server != nil {
		if err := s.server.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown failed: %v\n", err)
		}
	}
	log.Print("Server was shut down")
}

func (s ClientService) getPort(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	port, err := s.repository.GetPort(key)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if port == nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(port)
}

func main() {
	var c Config
	err := envconfig.Process("client", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	s := newClientService(&c)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	s.HandleRequests()
	sig := <-sigCh
	log.Printf("Received signal: %v\n", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}

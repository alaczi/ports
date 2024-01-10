package main

import (
	repo "github.com/alaczi/ports/repository"
	"go.uber.org/dig"
	"log"
	"os"
	"os/signal"
	"port_domain_service/api"
	"port_domain_service/domain/port"
	"port_domain_service/services"
	"syscall"
)

func main() {
	container := dig.New()
	container.Provide(services.NewConfig)
	container.Provide(provideRepository)
	container.Provide(providePortService)
	container.Provide(provideServer)

	startup := func(server *api.PortServer) {
		server.Serve()
	}
	if err := container.Invoke(startup); err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh
	log.Printf("Received signal: %v\n", sig)
	shutdown := func(server *api.PortServer) {
		server.Shutdown()
	}
	if err := container.Invoke(shutdown); err != nil {
		log.Printf("Error during shutdown %v", err)
	}
}

func provideRepository() repo.PortRepository {
	return port.NewInMemoryPortRepository()
}

func providePortService(repository repo.PortRepository) services.PortServiceInterface {
	return services.NewPortService(repository)
}

func provideServer(config *services.Config, postService services.PortServiceInterface) *api.PortServer {
	return api.NewPortServer(config, postService)
}

package main

import (
	"github.com/alaczi/ports/logger"
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
	if err := container.Provide(services.NewConfig); err != nil {
		log.Fatalf("Failed to read configration %v", err)
	}
	container.Provide(provideLogger)
	container.Provide(provideRepository)
	container.Provide(providePortService)
	container.Provide(provideServer)

	startup := func(server *api.PortServer) error {
		return server.Serve()
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

func provideLogger() logger.Logger {
	return &logger.ConsoleLogger{}
}
func provideRepository(log logger.Logger) repo.PortRepository {
	return port.NewInMemoryPortRepository(log)
}

func providePortService(repository repo.PortRepository) services.PortServiceInterface {
	return services.NewPortService(repository)
}

func provideServer(config *services.Config, log logger.Logger, postService services.PortServiceInterface) *api.PortServer {
	return api.NewPortServer(config, log, postService)
}

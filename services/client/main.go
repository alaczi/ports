package main

import (
	"client/api"
	"client/domain/port"
	"client/services"
	"context"
	repo "github.com/alaczi/ports/repository"
	"go.uber.org/dig"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	container := dig.New()

	ctx := context.Background()
	provideStartupContext := func() *context.Context {
		return &ctx
	}
	container.Provide(provideStartupContext)
	container.Provide(services.NewConfig)
	container.Provide(provideRepository)
	container.Provide(providePortService)
	container.Provide(provideHandlers)
	container.Provide(provideServer)

	startup := func(server *api.ClientServer) {
		server.HandleRequests()
	}
	if err := container.Invoke(startup); err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh
	log.Printf("Received signal: %v\n", sig)
	shutdown := func(server *api.ClientServer, ctx *context.Context) {
		contextWithTimeout, cancel := context.WithTimeout(*ctx, 3*time.Second)
		defer cancel()
		server.Shutdown(contextWithTimeout)
	}
	if err := container.Invoke(shutdown); err != nil {
		log.Printf("Error during shutdown %v", err)
	}
}

func provideRepository(config *services.Config, ctx *context.Context) repo.PortRepository {
	return port.NewGRPCPortRepository(*ctx, &config.PortServiceAddr, &config.DataFile)
}

func providePortService(repository repo.PortRepository) services.PortServiceInterface {
	return services.NewPortService(repository)
}

func provideHandlers(portService services.PortServiceInterface) []api.HttpRequestHandler {
	return []api.HttpRequestHandler{api.NewPortHandler(portService)}
}

func provideServer(config *services.Config, handlers []api.HttpRequestHandler) *api.ClientServer {
	return api.NewClientServer(config, handlers)
}

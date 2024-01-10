package port

import (
	"context"
	"github.com/alaczi/ports/logger"
	repo "github.com/alaczi/ports/repository"
	"sync"
)

type InMemoryPortRepository struct {
	mu     sync.Mutex
	ports  map[string]*repo.Port
	logger logger.Logger
}

func NewInMemoryPortRepository(log logger.Logger) *InMemoryPortRepository {
	r := &InMemoryPortRepository{
		ports:  make(map[string]*repo.Port),
		logger: log,
	}
	return r
}

func (r *InMemoryPortRepository) UpsertPort(ctx context.Context, port *repo.Port) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ports[port.Id] = port
	return nil
}

func (r *InMemoryPortRepository) GetPort(ctx context.Context, id string) (*repo.Port, error) {
	ch := make(chan *repo.Port)
	go func(ch chan *repo.Port, id string) {
		if port, ok := r.ports[id]; ok {
			ch <- port
		}
		ch <- nil
	}(ch, id)

	select {
	case <-ctx.Done():
		{
			r.logger.Logf("Context returned error before the port was retrieved: %v", ctx.Err())
			return nil, ctx.Err()
		}
	case port := <-ch:
		return port, nil
	}
}
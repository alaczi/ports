package port

import (
	"context"
	repo "github.com/alaczi/ports/repository"
	"sync"
)

type InMemoryPortRepository struct {
	mu    sync.Mutex
	ports map[string]*repo.Port
}

func NewInMemoryPortRepository() *InMemoryPortRepository {
	s := &InMemoryPortRepository{
		ports: make(map[string]*repo.Port),
	}
	return s
}

func (s *InMemoryPortRepository) UpsertPort(ctx context.Context, port *repo.Port) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ports[port.Id] = port
	return nil
}

func (s *InMemoryPortRepository) GetPort(ctx context.Context, id string) (*repo.Port, error) {
	ch := make(chan *repo.Port)
	go func(ch chan *repo.Port, id string) {
		if port, ok := s.ports[id]; ok {
			ch <- port
		}
		ch <- nil
	}(ch, id)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case port := <-ch:
		return port, nil
	}
}
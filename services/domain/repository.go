package main

import (
	repo "github.com/alaczi/ports/repository"
	"sync"
)

type InMemoryPortRepository struct {
	mu    sync.Mutex
	ports map[string]*repo.Port
}

func newInMemoryPortRepository() *InMemoryPortRepository {
	s := &InMemoryPortRepository{
		ports: make(map[string]*repo.Port),
	}
	return s
}

func (s *InMemoryPortRepository) UpsertPort(port *repo.Port) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ports[port.Id] = port
	return nil
}

func (s *InMemoryPortRepository) GetPort(id string) (*repo.Port, error) {
	if port, ok := s.ports[id]; ok {
		return port, nil
	}
	return nil, nil
}

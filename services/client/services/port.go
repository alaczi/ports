package services

import (
	"context"
	repo "github.com/alaczi/ports/repository"
)

type PortServiceInterface interface {
	GetPort(ctx context.Context, id string) (*repo.Port, error)
	NewPort(ctx context.Context, port *repo.Port) error
}

type PortService struct {
	repository repo.PortRepository
}

func NewPortService(repository repo.PortRepository) *PortService {
	service := &PortService{
		repository: repository,
	}
	return service
}

func (s *PortService) GetPort(ctx context.Context, id string) (*repo.Port, error) {
	return s.repository.GetPort(ctx, id)
}

func (s *PortService) NewPort(ctx context.Context, port *repo.Port) error {
	return s.repository.UpsertPort(ctx, port)
}
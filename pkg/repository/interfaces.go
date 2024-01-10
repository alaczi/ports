package repository

import "context"

type PortRepository interface {
	UpsertPort(ctx context.Context, port *Port) error

	GetPort(ctx context.Context, id string) (*Port, error)
}

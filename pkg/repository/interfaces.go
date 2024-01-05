package repository

type PortRepository interface {
	UpsertPort(port *Port) error

	GetPort(id string) (*Port, error)
}

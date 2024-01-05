package ports

import (
	repo "github.com/alaczi/ports/repository"
)

func ToProtoPort(port *repo.Port) *Port {
	return &Port{
		Id:          port.Id,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
}

func ToPort(port *Port) *repo.Port {
	if port == nil {
		return nil
	}
	return &repo.Port{
		Id:          port.Id,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
}

func ToPorts(ports []*Port) []*repo.Port {
	var results = make([]*repo.Port, len(ports))
	for key, value := range ports {
		results[key] = ToPort(value)
	}
	return results
}

func ToProtoPorts(ports []*repo.Port) *Ports {
	var results = make([]*Port, len(ports))
	for key, value := range ports {
		results[key] = ToProtoPort(value)
	}
	return &Ports{Ports: results}
}

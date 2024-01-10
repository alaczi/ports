package services

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int `default:"50051"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := envconfig.Process("domain", config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
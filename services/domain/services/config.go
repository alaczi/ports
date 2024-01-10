package services

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port int `default:"50051"`
}

func NewConfig() *Config {
	config := &Config{}
	err := envconfig.Process("domain", config)
	if err != nil {
		log.Fatalf("Failed to read configuration %v", err.Error())
	}
	return config
}
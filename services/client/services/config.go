package services

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	ServerPort      int    `default:"8080"`
	PortServiceAddr string `default:"localhost:50051"`
	DataFile        string `default:"./data/ports.json"`
}

func NewConfig() *Config {
	config := &Config{}
	err := envconfig.Process("client", config)
	if err != nil {
		log.Fatalf("Failed to read configuration %v", err.Error())
	}
	return config
}
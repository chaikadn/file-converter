package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string `env:"SERVER_HOST"`
	Port string `env:"SERVER_PORT"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	// for dev only
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	// load env

	// parse flags

	return nil
}

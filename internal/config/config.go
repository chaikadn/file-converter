package config

import (
	"fmt"
	"log/slog"

	"github.com/caarlos0/env/v6"
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
		return fmt.Errorf("failed to load env: %w", err)
	}

	if err := c.parseEnv(); err != nil {
		return fmt.Errorf("failed to parse env: %w", err)
	}
	slog.Debug("env parsed")

	// parse flags

	return nil
}

func (c *Config) parseEnv() error {
	return env.Parse(c)
}

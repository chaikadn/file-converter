package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/chaikadn/file-converter/internal/config"
	"github.com/chaikadn/file-converter/internal/server"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("Logger initialized")

	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}
	slog.Info("Config loaded")

	h := server.NewHandler(nil, nil)
	s := server.NewServer(net.JoinHostPort(cfg.Host, cfg.Port), h)

	slog.Info("Starting server", "host", cfg.Host, "port", cfg.Port)
	if err := s.Run(); err != nil {
		slog.Error("fatal server error", "error", err)
		os.Exit(1)
	}
}

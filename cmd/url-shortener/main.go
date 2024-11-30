package main

import (
	"log/slog"

	"github.com/jdaniecki/url-shortener/internal/persistence"
	"github.com/jdaniecki/url-shortener/internal/server"
)

var version string
var host = "localhost:8080"

func startServer(host string) error {
	storage := persistence.NewInMemoryStorage()
	s := server.New(storage, host)
	return s.Serve()
}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("Starting url-shortener version", "version", version)
	if err := startServer(host); err != nil {
		slog.Error("could not start servern", "error", err)
	}
}

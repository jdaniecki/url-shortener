package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/jdaniecki/url-shortener/internal/persistence"
	"github.com/jdaniecki/url-shortener/internal/server"
)

var version string
var fqdn string
var port uint

func startServer(host string) error {
	storage := persistence.NewInMemoryStorage()
	s := server.New(storage, host)
	return s.Serve()
}

func main() {
	// setup logger
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("Starting url-shortener version", "version", version)

	// setup flags
	flag.StringVar(&fqdn, "fqdn", "localhost", "fqdn to serve on")
	flag.UintVar(&port, "port", 8080, "port to listen on")
	flag.Parse()

	// start the http server
	host := fmt.Sprintf("%s:%d", fqdn, port)
	if err := startServer(host); err != nil {
		slog.Error("could not start servern", "error", err)
	}
}

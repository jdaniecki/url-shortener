package main

import (
	"log"

	"github.com/jdaniecki/url-shortener/internal/persistence"
	"github.com/jdaniecki/url-shortener/internal/server"
)

var version string
var host = "localhost:8080"

func main() {
	log.Printf("Starting url-shortener version %s\n", version)

	log.Printf("Starting HTTP server on %v", host)
	storage := persistence.NewInMemoryStorage()
	s := server.New(storage, host)
	if err := s.Serve(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}

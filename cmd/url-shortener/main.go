package main

import (
	"log"
	"net/http"

	"github.com/jdaniecki/url-shortener/internal/api"
	"github.com/jdaniecki/url-shortener/internal/persistence"
	"github.com/jdaniecki/url-shortener/internal/server"
)

var version string

func main() {
	log.Printf("Starting url-shortener version %s\n", version)

	storage := persistence.NewInMemoryStorage()
	s := server.New(storage)

	strictHandler := api.NewStrictHandler(s, nil)
	r := http.NewServeMux()
	handler := api.HandlerFromMux(strictHandler, r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}

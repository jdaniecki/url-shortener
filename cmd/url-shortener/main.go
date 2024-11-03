package main

import (
	"log"
	"net/http"

	"github.com/jdaniecki/url-shortener/internal/api"
	"github.com/jdaniecki/url-shortener/internal/server"
)

func main() {
	s := &server.Server{}
	handler := api.Handler(s)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}

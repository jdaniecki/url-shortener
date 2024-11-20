package main

import (
	"log"
	"net/http"

	"github.com/jdaniecki/url-shortener/internal/api"
	"github.com/jdaniecki/url-shortener/internal/persistence"
	"github.com/jdaniecki/url-shortener/internal/server"
	_ "github.com/oapi-codegen/nethttp-middleware"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

var version string

func main() {
	log.Printf("Starting url-shortener version %s\n", version)

	storage := persistence.NewInMemoryStorage()
	s := server.New(storage)

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("Failed to load OpenAPI specification: %v\n", err)
	}

	nethttpvalidator := nethttpmiddleware.OapiRequestValidatorWithOptions(swagger, &nethttpmiddleware.Options{SilenceServersWarning: true})

	handler := nethttpvalidator(api.Handler(api.NewStrictHandler(s, nil)))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}

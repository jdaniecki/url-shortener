package server

import (
	"io"
	"log"
	"net/http"
)

type Server struct {
	// Add any necessary fields here
}

func (s *Server) PostShorten(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for shortening a URL
	log.Printf("Received request to shorten URL: %v", r.URL)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	log.Printf("Request body: %s", body)
}

func (s *Server) GetShortUrl(w http.ResponseWriter, r *http.Request, shortUrl string) {
	// Implement the logic for retrieving the original URL
	log.Printf("Received request for short URL: %s", shortUrl)
}

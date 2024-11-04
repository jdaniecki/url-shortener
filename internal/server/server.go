package server

import (
	"encoding/json"
	"net/http"

	"github.com/jdaniecki/url-shortener/internal/persistence"
)

type Server struct {
	storage persistence.Storage
}

func NewServer(storage persistence.Storage) *Server {
	return &Server{storage: storage}
}

func (s *Server) PostShorten(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.URL == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}
	shortUrl, err := s.storage.Save(req.URL)
	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}
	resp := map[string]string{"shortUrl": "http://short.url/" + shortUrl}
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) GetShortUrl(w http.ResponseWriter, r *http.Request, shortUrl string) {
	originalUrl, err := s.storage.Load(shortUrl)
	if err != nil {
		http.Error(w, "Failed to load URL", http.StatusNotFound)
		return
	}
	resp := struct {
		OriginalURL string `json:"originalUrl"`
	}{
		OriginalURL: originalUrl,
	}
	json.NewEncoder(w).Encode(resp)
}

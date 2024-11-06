package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jdaniecki/url-shortener/internal/api"
	"github.com/jdaniecki/url-shortener/internal/persistence"
)

type Server struct {
	storage persistence.Storage
	url     string
}

// Make sure we conform to ServerInterface
var _ api.ServerInterface = (*Server)(nil)

func New(storage persistence.Storage) *Server {
	return &Server{
		storage: storage,
		url:     "http://localhost:8080/"}
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
	resp := map[string]string{"shortUrl": s.url + shortUrl}
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) GetShortUrl(w http.ResponseWriter, r *http.Request, shortUrl string) {
	originalUrl, err := s.storage.Load(shortUrl)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load URL for %v", shortUrl), http.StatusNotFound)
		return
	}
	resp := map[string]string{"originalUrl": originalUrl}
	json.NewEncoder(w).Encode(resp)
}

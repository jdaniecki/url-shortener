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
	json.NewDecoder(r.Body).Decode(&req)
	shortUrl, _ := s.storage.Save(req.URL)
	resp := map[string]string{"shortUrl": "http://short.url/" + shortUrl}
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) GetShortUrl(w http.ResponseWriter, r *http.Request, shortUrl string) {
	originalUrl, _ := s.storage.Load(shortUrl)
	resp := map[string]string{"originalUrl": originalUrl}
	json.NewEncoder(w).Encode(resp)
}

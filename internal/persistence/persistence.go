package persistence

import (
	"fmt"

	"log/slog"

	"github.com/jdaniecki/url-shortener/internal/shortener"
)

type Storage interface {
	Save(url string) (string, error)
	Load(shortUrl string) (string, error)
}

type InMemoryStorage struct {
	data      map[string]string
	shortener *shortener.Shortener
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		data:      make(map[string]string),
		shortener: shortener.NewShortener(),
	}
}

func (s *InMemoryStorage) Save(longURL string) (string, error) {
	shortUrl := s.shortener.Shorten(longURL)
	s.data[shortUrl] = longURL
	slog.Debug("URL persisted in memory", "longURL", longURL, "shortURL", shortUrl)
	return shortUrl, nil
}

func (s *InMemoryStorage) Load(shortUrl string) (string, error) {
	longURL, exists := s.data[shortUrl]
	if !exists {
		slog.Debug("URL not found", "shortURL", shortUrl)
		return "", fmt.Errorf("URL not found")
	}
	slog.Debug("URL retrived from memory", "longURL", longURL, "shortURL", shortUrl)
	return longURL, nil
}

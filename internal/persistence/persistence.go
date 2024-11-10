package persistence

import (
	"fmt"
	"log"

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

func (s *InMemoryStorage) Save(url string) (string, error) {
	shortUrl := s.shortener.Shorten(url)
	s.data[shortUrl] = url
	log.Printf("Saved URL %v as %v\n", url, shortUrl)
	return shortUrl, nil
}

func (s *InMemoryStorage) Load(shortUrl string) (string, error) {
	url, exists := s.data[shortUrl]
	log.Printf("Loaded URL %v for %v\n", url, shortUrl)
	if !exists {
		return "", fmt.Errorf("URL not found")
	}
	return url, nil
}

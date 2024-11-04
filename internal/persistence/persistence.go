package persistence

import "fmt"

type Storage interface {
	Save(url string) (string, error)
	Load(shortUrl string) (string, error)
}

type InMemoryStorage struct {
	data map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{data: make(map[string]string)}
}

func (s *InMemoryStorage) Save(url string) (string, error) {
	shortUrl := "abc123" // Simplified for example purposes
	s.data[shortUrl] = url
	return shortUrl, nil
}

func (s *InMemoryStorage) Load(shortUrl string) (string, error) {
	url, exists := s.data[shortUrl]
	if !exists {
		return "", fmt.Errorf("URL not found")
	}
	return url, nil
}

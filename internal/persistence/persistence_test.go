package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryStorage_Save(t *testing.T) {
	storage := NewInMemoryStorage()

	tests := []struct {
		url      string
		expected string
	}{
		{"http://example.com", "0"},
		{"http://example.org", "1"},
		{"http://example.net", "2"},
	}

	for _, test := range tests {
		shortUrl, err := storage.Save(test.url)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, shortUrl)
	}
}

func TestInMemoryStorage_Load(t *testing.T) {
	storage := NewInMemoryStorage()

	urls := []string{
		"http://example.com",
		"http://example.org",
		"http://example.net",
	}

	for _, url := range urls {
		storage.Save(url)
	}

	tests := []struct {
		shortUrl string
		expected string
	}{
		{"0", "http://example.com"},
		{"1", "http://example.org"},
		{"2", "http://example.net"},
	}

	for _, test := range tests {
		url, err := storage.Load(test.shortUrl)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, url)
	}

	_, err := storage.Load("nonexistent")
	assert.Error(t, err)
}

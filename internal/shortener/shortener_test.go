package shortener

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{61, "Z"},
		{62, "10"},
		{123, "1Z"},
	}

	for _, test := range tests {
		s := NewShortener()
		for i := 0; i < test.input; i++ {
			s.Shorten("http://example.com")
		}
		shortUrl := s.Shorten("http://example.com")
		assert.Equal(t, test.expected, shortUrl, "Shorten(%d) = %s; want %s", test.input, shortUrl, test.expected)
	}
}

func TestConcurrentShorten(t *testing.T) {
	s := NewShortener()
	const numGoroutines = 100
	const numUrls = 1000

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numUrls; j++ {
				s.Shorten("http://example.com")
			}
		}()
	}

	wg.Wait()

	expectedID := numGoroutines * numUrls
	assert.Equal(t, expectedID, s.id, "Expected id to be %d, got %d", expectedID, s.id)
}

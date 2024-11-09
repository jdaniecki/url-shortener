package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jdaniecki/url-shortener/internal/persistence"
	"github.com/jdaniecki/url-shortener/internal/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostShorten(t *testing.T) {
	storage := persistence.NewInMemoryStorage()
	s := server.New(storage)
	reqBody := bytes.NewBufferString(`{"url": "http://example.com"}`)
	req, err := http.NewRequest("POST", "/shorten", reqBody)
	assert.NoError(t, err, "Could not create request")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.PostShorten)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	expected := `{"shortUrl": "http://localhost:8080/0"}`
	assert.JSONEq(t, expected, rr.Body.String(), "handler returned unexpected body")
}

func TestPostShortenInvalidJSON(t *testing.T) {
	storage := persistence.NewInMemoryStorage()
	s := server.New(storage)
	reqBody := bytes.NewBufferString(`{"url": "http://example.com"`)
	req, err := http.NewRequest("POST", "/shorten", reqBody)
	assert.NoError(t, err, "Could not create request")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.PostShorten)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "handler returned wrong status code")
}

func TestPostShortenEmptyURL(t *testing.T) {
	storage := persistence.NewInMemoryStorage()
	s := server.New(storage)
	reqBody := bytes.NewBufferString(`{"url": ""}`)
	req, err := http.NewRequest("POST", "/shorten", reqBody)
	assert.NoError(t, err, "Could not create request")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.PostShorten)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "handler returned wrong status code")
}

func TestGetShortUrl(t *testing.T) {
	storage := persistence.NewInMemoryStorage()
	_, err := storage.Save("http://example.com")
	assert.NoError(t, err, "Could not save URL")
	s := server.New(storage)
	req, err := http.NewRequest("GET", "/1", nil)
	assert.NoError(t, err, "Could not create request")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.GetShortUrl(w, r, "0")
	})
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code; %v", rr.Body.String())

	expected := `{"originalUrl": "http://example.com"}`
	assert.NotNil(t, rr.Body, "handler returned nil body")
	assert.JSONEq(t, expected, rr.Body.String(), "handler returned unexpected body")
}

func TestGetShortUrlNotFound(t *testing.T) {
	storage := persistence.NewInMemoryStorage()
	s := server.New(storage)
	req, err := http.NewRequest("GET", "/short/unknown", nil)
	assert.NoError(t, err, "Could not create request")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.GetShortUrl(w, r, "unknown")
	})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "handler returned wrong status code")
}

package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostShorten(t *testing.T) {
	s := &Server{}
	reqBody := bytes.NewBufferString(`{"url": "http://example.com"}`)
	req, err := http.NewRequest("POST", "/shorten", reqBody)
	assert.NoError(t, err, "Could not create request")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.PostShorten)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	expected := `{"shortUrl": "http://short.url/abc123"}`
	assert.JSONEq(t, expected, rr.Body.String(), "handler returned unexpected body")
}

func TestGetShortUrl(t *testing.T) {
	s := &Server{}
	req, err := http.NewRequest("GET", "/short/abc123", nil)
	assert.NoError(t, err, "Could not create request")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.GetShortUrl(w, r, "abc123")
	})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	expected := `{"originalUrl": "http://example.com"}`
	assert.NotNil(t, rr.Body, "handler returned nil body")
	assert.JSONEq(t, expected, rr.Body.String(), "handler returned unexpected body")
}

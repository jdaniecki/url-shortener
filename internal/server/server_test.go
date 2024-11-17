package server_test

import (
	"context"
	"testing"

	"github.com/jdaniecki/url-shortener/internal/api"
	"github.com/jdaniecki/url-shortener/internal/persistence"
	"github.com/jdaniecki/url-shortener/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestPostShorten(t *testing.T) {
	storage := persistence.NewInMemoryStorage()
	server := server.New(storage)

	t.Run("valid URL", func(t *testing.T) {
		req := api.PostShortenRequestObject{
			Body: &api.PostShortenJSONRequestBody{Url: stringPtr("http://example.com")},
		}
		resp, err := server.PostShorten(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "0", *resp.(api.PostShorten200JSONResponse).ShortUrl)
	})

	t.Run("invalid URL", func(t *testing.T) {
		req := api.PostShortenRequestObject{
			Body: &api.PostShortenJSONRequestBody{Url: stringPtr("")},
		}
		resp, err := server.PostShorten(context.Background(), req)
		assert.NoError(t, err)
		assert.IsType(t, api.PostShorten400Response{}, resp)
	})
}
func TestGetShortUrl(t *testing.T) {
	storage := persistence.NewInMemoryStorage()
	server := server.New(storage)
	storage.Save("http://example.com")

	t.Run("existing short URL", func(t *testing.T) {
		req := api.GetShortUrlRequestObject{ShortUrl: "0"}
		resp, err := server.GetShortUrl(context.Background(), req)
		assert.NoError(t, err)
		assert.IsType(t, api.GetShortUrl302Response{}, resp)
		assert.Equal(t, "http://example.com", resp.(api.GetShortUrl302Response).Headers.Location)
	})

	t.Run("non-existing short URL", func(t *testing.T) {
		req := api.GetShortUrlRequestObject{ShortUrl: "nonExisting"}
		resp, err := server.GetShortUrl(context.Background(), req)
		assert.NoError(t, err)
		assert.IsType(t, api.GetShortUrl404Response{}, resp)
	})
}

func stringPtr(s string) *string {
	return &s
}

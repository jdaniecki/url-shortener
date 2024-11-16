package server

import (
	"context"

	"github.com/jdaniecki/url-shortener/internal/api"
	"github.com/jdaniecki/url-shortener/internal/persistence"
)

type Server struct {
	storage persistence.Storage
	url     string
}

// Make sure we conform to StrictServerInterface
var _ api.StrictServerInterface = (*Server)(nil)

func New(storage persistence.Storage) *Server {
	return &Server{
		storage: storage,
		url:     "http://localhost:8080/"}
}

func (s *Server) PostShorten(ctx context.Context, request api.PostShortenRequestObject) (api.PostShortenResponseObject, error) {
	if request.Body == nil || request.Body.Url == nil || *request.Body.Url == "" {
		return api.PostShorten400Response{}, nil
	}
	shortUrl, err := s.storage.Save(*request.Body.Url)
	if err != nil {
		return nil, err
	}
	resp := api.PostShorten200JSONResponse{ShortUrl: &shortUrl}
	return resp, nil
}

func (s *Server) GetShortUrl(ctx context.Context, request api.GetShortUrlRequestObject) (api.GetShortUrlResponseObject, error) {
	_, err := s.storage.Load(request.ShortUrl)
	if err != nil {
		return api.GetShortUrl404Response{}, nil
	}
	//http.Redirect(ctx.Value(http.ResponseWriter).(http.ResponseWriter), ctx.Value(http.Request).(*http.Request), originalUrl, http.StatusFound)
	return api.GetShortUrl302Response{}, nil
}

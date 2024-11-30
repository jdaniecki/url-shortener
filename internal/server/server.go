package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jdaniecki/url-shortener/internal/api"
	"github.com/jdaniecki/url-shortener/internal/persistence"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

type Server struct {
	storage persistence.Storage
	host    string
}

// Make sure we conform to StrictServerInterface
var _ api.StrictServerInterface = (*Server)(nil)

func New(storage persistence.Storage, host string) *Server {
	slog.Debug("Creating new server", "host", host)

	return &Server{
		storage: storage,
		host:    host,
	}
}

func (s *Server) Serve() error {
	slog.Debug("Starting server", "host", s.host)

	swagger, err := api.GetSwagger()
	if err != nil {
		return err
	}
	swagger.Servers = nil

	nethttpvalidator := nethttpmiddleware.OapiRequestValidatorWithOptions(swagger,
		&nethttpmiddleware.Options{ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			slog.Debug("Request validation failed", "message", message, "statusCode", statusCode)
			http.Error(w, message, statusCode)
		}})

	serverHandler := api.Handler(api.NewStrictHandler(s, nil))

	handler := nethttpvalidator(serverHandler)

	if err := http.ListenAndServe(s.host, handler); err != nil {
		return err
	}
	return nil
}

func (s *Server) PostShorten(ctx context.Context, request api.PostShortenRequestObject) (api.PostShortenResponseObject, error) {
	slog.Debug("Received 'PostShorten' request", "url", *request.Body.Url)

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
	slog.Debug("Received 'GetShortUrl' request", "shortUrl", request.ShortUrl)

	longUrl, err := s.storage.Load(request.ShortUrl)
	if err != nil {
		return api.GetShortUrl404Response{}, nil
	}
	headers := api.GetShortUrl302ResponseHeaders{Location: longUrl}
	return api.GetShortUrl302Response{Headers: headers}, nil
}

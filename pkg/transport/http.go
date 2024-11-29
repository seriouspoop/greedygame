package transport

import (
	"context"
	"fmt"
	"net/http"
	"seriouspoop/greedygame/pkg/config"
	"seriouspoop/greedygame/pkg/svc"

	"github.com/rs/zerolog"
)

type Http struct {
	HttpServer *http.Server
	WebConfig  config.WebServer
	logger     *zerolog.Logger
	svc        *svc.Svc
}

func NewHTTPServer(webConfig config.WebServer, svc *svc.Svc, logger *zerolog.Logger) *Http {
	return &Http{&http.Server{}, webConfig, logger, svc}
}

func (h *Http) Initialize(ctx context.Context) error {
	router := NewRouter(h.svc, h.logger).Initialize(h.WebConfig.RoutePrefix)
	h.HttpServer.Addr = fmt.Sprintf(":%d", h.WebConfig.Port)
	h.HttpServer.Handler = router
	return nil
}

func (h *Http) Run(ctx context.Context) error {
	fmt.Println("Starting server on port ", h.WebConfig.Port)
	if err := h.HttpServer.ListenAndServe(); err != nil {
		h.logger.Error().Err(err).Msg("error while ListenAndServe")
		return err
	}
	return nil
}

func (h *Http) Shutdown(ctx context.Context) error {
	return h.HttpServer.Close()
}

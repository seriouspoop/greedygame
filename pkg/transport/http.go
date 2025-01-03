package transport

import (
	"context"
	"fmt"
	"net/http"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/go-common/observer"
	"seriouspoop/greedygame/pkg/config"
	"seriouspoop/greedygame/pkg/svc"

	"go.uber.org/zap"
)

type Http struct {
	HttpServer *http.Server
	WebConfig  config.WebServer
	obs        *observer.Observer
	logger     *logging.Logger
	svc        *svc.Svc
}

func NewHTTPServer(webConfig config.WebServer, obs *observer.Observer, svc *svc.Svc, logger *logging.Logger) *Http {
	return &Http{&http.Server{}, webConfig, obs, logger, svc}
}

func (h *Http) Initialize(ctx context.Context) error {
	// Initialize gateway
	// conn, err := grpc.NewClient(
	// 	fmt.Sprintf(":%d", h.WebConfig.GrpcPort),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	// )

	// if err != nil {
	// 	return err
	// }
	// Initialize http
	router := NewRouter(h.svc, h.logger, h.obs).Initialize(ctx)
	h.HttpServer.Addr = fmt.Sprintf(":%d", h.WebConfig.RestPort)
	h.HttpServer.Handler = router.mux
	return nil
}

func (h *Http) Run(ctx context.Context) error {
	log := h.logger.Ctx(ctx)
	fmt.Println("Starting server on port ", h.WebConfig.RestPort)
	if err := h.HttpServer.ListenAndServe(); err != nil {
		log.Error("error while ListenAndServe", zap.Error(err))
		return err
	}
	return nil
}

func (h *Http) Shutdown(ctx context.Context) error {
	return h.HttpServer.Close()
}

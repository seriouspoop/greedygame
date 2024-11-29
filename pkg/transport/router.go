package transport

import (
	"net/http"
	"seriouspoop/greedygame/pkg/svc"
	"seriouspoop/greedygame/pkg/transport/handler"
	"seriouspoop/greedygame/pkg/transport/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Router struct {
	*mux.Router
	svc    *svc.Svc
	logger *zerolog.Logger
}

func NewRouter(svc *svc.Svc, logger *zerolog.Logger) *Router {
	return &Router{mux.NewRouter(), svc, logger}
}

func (r *Router) Initialize(routePrefix string) *Router {
	logMw := middleware.NewLog(r.logger, zerolog.InfoLevel)

	r.Use(logMw.LogMiddleware)

	deliveryMux := r.PathPrefix(routePrefix).Subrouter()
	deliveryMux.HandleFunc("/healthcheck", handler.HealthCheck(r.svc)).Methods(http.MethodGet)

	// v1
	// v1 := deliveryMux.PathPrefix("/v1").Subrouter()

	return r
}

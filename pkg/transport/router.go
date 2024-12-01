package transport

import (
	"net/http"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/go-common/middleware"
	"seriouspoop/greedygame/pkg/svc"
	"seriouspoop/greedygame/pkg/transport/handler"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Router struct {
	*mux.Router
	svc    *svc.Svc
	logger *logging.Logger
}

func NewRouter(svc *svc.Svc, logger *logging.Logger) *Router {
	return &Router{mux.NewRouter(), svc, logger}
}

func (r *Router) Initialize() *Router {
	logMw := middleware.NewLog(r.logger, zerolog.InfoLevel)

	r.Use(middleware.TraceMiddleware)
	r.Use(logMw.LogMiddleware)

	r.HandleFunc("/healthcheck", handler.HealthCheck(r.svc)).Methods(http.MethodGet)

	// v1
	v1 := r.PathPrefix("/v1").Subrouter()

	//delivery
	v1.HandleFunc("/delivery", handler.Delivery(r.svc, r.logger)).Methods(http.MethodGet)

	return r
}

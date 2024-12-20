package transport

import (
	"net/http"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/go-common/middleware"
	"seriouspoop/greedygame/go-common/observer"
	"seriouspoop/greedygame/pkg/svc"
	"seriouspoop/greedygame/pkg/transport/handler"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Router struct {
	*mux.Router
	svc    *svc.Svc
	logger *logging.Logger
	obs    *observer.Observer
}

func NewRouter(svc *svc.Svc, logger *logging.Logger, obs *observer.Observer) *Router {
	return &Router{mux.NewRouter(), svc, logger, obs}
}

func (r *Router) Initialize() *Router {
	logMw := middleware.NewLog(r.logger, zap.InfoLevel)
	traceMw := r.obs.TraceSDK()

	r.Use(traceMw.TraceHTTPMiddleware)
	r.Use(logMw.LogMiddleware)

	r.HandleFunc("/healthcheck", handler.HealthCheck(r.svc)).Methods(http.MethodGet)

	// v1
	v1 := r.PathPrefix("/v1").Subrouter()

	//delivery
	v1.HandleFunc("/delivery", handler.Delivery(r.svc, r.logger)).Methods(http.MethodGet)

	return r
}

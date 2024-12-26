package transport

import (
	"context"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/go-common/observer"
	"seriouspoop/greedygame/pkg/svc"
	"seriouspoop/greedygame/pkg/transport/handler"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Router struct {
	mux    *runtime.ServeMux
	svc    *svc.Svc
	logger *logging.Logger
	obs    *observer.Observer
	// conn   *grpc.ClientConn
}

// func lund(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx := context.WithValue(r.Context(), "Lund", "Value")
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func NewRouter(svc *svc.Svc, logger *logging.Logger, obs *observer.Observer) *Router {
	// TODO add custom middlewares for the mux
	// logMw := middleware.NewLog(logger, zap.InfoLevel)
	// traceMw := obs.TraceSDK()

	mux := runtime.NewServeMux(runtime.WithMiddlewares())
	return &Router{mux, svc, logger, obs}
}

func (r *Router) Initialize(ctx context.Context) *Router {

	// Gateway handlers
	handler.NewDeliveryGatewayHandler(ctx, r.mux, r.svc, r.logger)
	handler.NewHealthCheckGatewayHandler(ctx, r.mux, r.svc, r.logger)

	// // TODO - remove these and implement self handlers
	// commonpb.RegisterHealthHandler(ctx, r.mux, r.conn)
	// deliverypb.RegisterDeliveryHandler(ctx, r.mux, r.conn)

	return r
}

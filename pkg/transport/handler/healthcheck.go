package handler

import (
	"context"
	"net/http"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/protos/go/commonpb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthCheckHandler struct {
	s      healthServicer
	logger *logging.Logger
	commonpb.UnimplementedHealthServer
}

// New Grpc server handler
func NewHealthCheckHandler(grpc *grpc.Server, s healthServicer, logger *logging.Logger) {
	handler := &HealthCheckHandler{s: s, logger: logger}
	commonpb.RegisterHealthServer(grpc, handler)
}

// New Grpc gateway handler - Currently Unused
func NewHealthCheckGatewayHandler(ctx context.Context, runtime *runtime.ServeMux, s healthServicer, logger *logging.Logger) {
	handler := &HealthCheckHandler{s: s, logger: logger}
	commonpb.RegisterHealthHandlerServer(ctx, runtime, handler)
}

func (h *HealthCheckHandler) CheckHealth(ctx context.Context, req *emptypb.Empty) (*commonpb.HealthResponse, error) {
	log := h.logger.Ctx(ctx)

	log.Info("hello inside check handler")

	if h.s.IsUnhealthy(ctx) {
		return &commonpb.HealthResponse{
			Code:    http.StatusServiceUnavailable,
			Message: "service not healthy",
		}, nil
	}
	return &commonpb.HealthResponse{
		Code:    http.StatusOK,
		Message: "service healthy",
	}, nil
}

package transport

import (
	"context"
	"fmt"
	"net"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/pkg/config"
	"seriouspoop/greedygame/pkg/svc"
	"seriouspoop/greedygame/pkg/transport/handler"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Grpc struct {
	GrpcServer *grpc.Server
	WebConfig  config.WebServer
	logger     *logging.Logger
	svc        *svc.Svc
}

func NewGRPCServer(webConfig config.WebServer, logger *logging.Logger, s *svc.Svc) *Grpc {
	return &Grpc{grpc.NewServer(), webConfig, logger, s}
}

func (g *Grpc) Initialize(ctx context.Context) error {
	// Register service handlers
	handler.NewDeliveryHandler(g.GrpcServer, g.svc, g.logger)
	handler.NewHealthCheckHandler(g.GrpcServer, g.svc, g.logger)
	return nil
}

func (g *Grpc) Run(ctx context.Context) error {
	log := g.logger.Ctx(ctx)
	port := fmt.Sprintf(":%d", g.WebConfig.GrpcPort)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Error("error while setting up listener", zap.Error(err))
		return err
	}

	return g.GrpcServer.Serve(lis)
}

func (g *Grpc) Shutdown(ctx context.Context) error {
	g.GrpcServer.GracefulStop()
	return nil
}

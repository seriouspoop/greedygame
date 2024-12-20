package transport

import (
	"context"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/go-common/observer"
	"seriouspoop/greedygame/pkg/config"
	"seriouspoop/greedygame/pkg/repo/postgres"
	"seriouspoop/greedygame/pkg/svc"
)

type server struct {
	http     Booter
	observer *observer.Observer
	logger   *logging.Logger
}

func NewServer(ctx context.Context, appCfg *config.App) (*server, error) {
	// setup exporter
	ex := observer.NewProductionExporter("otel-collector")

	// setup observer
	obs, err := observer.New(ctx, "seriouspoop/greedygame", ex)
	if err != nil {
		return nil, err
	}

	logger, err := logging.NewWithService(appCfg.WebServer.Service, appCfg.Log.Level, obs.LogSDK().NewLoggerCore())
	if err != nil {
		return nil, err
	}

	db, err := postgres.New(logger)
	if err != nil {
		return nil, err
	}

	s := svc.New(db, logger, obs.TraceSDK())

	return &server{
		http:     NewHTTPServer(appCfg.WebServer, obs, s, logger),
		observer: obs,
		logger:   logger,
	}, nil
}

func (s *server) Initialize(ctx context.Context) error {
	return s.http.Initialize(ctx)
}

func (s *server) Run(ctx context.Context) error {
	return s.http.Run(ctx)
}

func (s *server) Shutdown(ctx context.Context) error {
	err := s.observer.Shutdown(ctx)
	if err != nil {
		return err
	}
	err = s.logger.Sync()
	if err != nil {
		return err
	}
	return s.http.Shutdown(ctx)
}

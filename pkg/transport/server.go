package transport

import (
	"context"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/pkg/config"
	"seriouspoop/greedygame/pkg/repo/db"
	"seriouspoop/greedygame/pkg/svc"
)

type server struct {
	http Booter
}

func NewServer(ctx context.Context, appCfg *config.App) (*server, error) {
	logger, err := logging.NewWithService(appCfg.WebServer.Service, appCfg.Log.Level)
	if err != nil {
		return nil, err
	}

	db, err := db.New(logger)
	if err != nil {
		return nil, err
	}

	s := svc.New(db, logger)

	return &server{
		http: NewHTTPServer(appCfg.WebServer, s, logger),
	}, nil
}

func (s *server) Initialize(ctx context.Context) error {
	s.http.Initialize(ctx)
	return nil
}

func (s *server) Run(ctx context.Context) error {
	return s.http.Run(ctx)
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

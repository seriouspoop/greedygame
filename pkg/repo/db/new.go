package db

import (
	"context"
	"seriouspoop/greedygame/go-common/db/postgres"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/go-common/observer"
	"seriouspoop/greedygame/pkg/config"
	schema "seriouspoop/greedygame/pkg/repo/db/schema/gen"
	"seriouspoop/greedygame/pkg/svc"
)

type DB struct {
	query  *schema.Queries
	conn   *postgres.DB
	tracer *observer.Tracer
	logger *logging.Logger
}

func New(ctx context.Context, cfg config.Postgres, logger *logging.Logger, tracer *observer.Tracer) (*DB, error) {
	svcErrors := &postgres.Errors{
		TimeoutErr:    svc.ErrTimeout,
		NoDataErr:     svc.ErrNoData,
		DownErr:       svc.ErrUnexpected,
		UnexpectedErr: svc.ErrUnexpected,
	}

	conn, err := postgres.New(ctx, postgres.Config{
		ConnectionString: cfg.ConnectionString,
		Database:         cfg.Database,
	}, svcErrors, logger)

	if err != nil {
		return nil, err
	}
	query := schema.New(conn.Conn())

	return &DB{
		query:  query,
		conn:   conn,
		tracer: tracer,
		logger: logger,
	}, nil
}

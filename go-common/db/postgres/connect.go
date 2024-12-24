package postgres

import (
	"context"
	"fmt"
	"seriouspoop/greedygame/go-common/logging"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	connTimeout = 10 * time.Second
)

var (
	e   *Errors
	log *logging.LoggerWithCtx
)

type Config struct {
	ConnectionString string `required:"true" split_words:"true"`
	Database         string `required:"true" split_words:"true"`
}

type DB struct {
	conn *pgxpool.Pool
}

func New(ctx context.Context, cfg Config, errs *Errors, logger *logging.Logger) (*DB, error) {
	ctx, cancel := context.WithTimeout(ctx, connTimeout)
	defer cancel()

	connectionString := fmt.Sprintf("%s/%s", cfg.ConnectionString, cfg.Database)
	conn, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	// setting up global environment
	e = errs
	log = logger.Ctx(ctx)
	return &DB{conn: conn}, nil
}

func (d *DB) Conn() *pgxpool.Pool {
	return d.conn
}

func (d *DB) Ping(ctx context.Context) error {
	return d.conn.Ping(ctx)
}

// TODO - transaction implementation
func (d *DB) Tx(ctx context.Context) (pgx.Tx, error) {
	// var tx *pgxpool.Tx
	return d.conn.Begin(ctx)
}

func (d *DB) Close() {
	d.conn.Close()
}

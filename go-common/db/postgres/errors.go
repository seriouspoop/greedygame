package postgres

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx"
	"go.uber.org/zap"
)

type Errors struct {
	TimeoutErr    error //timeout occurred
	NoDataErr     error //no document found
	DownErr       error //when DB appears to be offline
	UnexpectedErr error //any unhandled errors
}

// converts pgx errors to service layer errors.
func SvcError(err error) error {
	if err != nil {
		if errors.Is(err, pgx.ErrAcquireTimeout) {
			log.Error("db timeout error", zap.Error(err))
			return e.TimeoutErr
		} else if errors.Is(err, pgx.ErrDeadConn) || errors.Is(err, pgx.ErrClosedPool) {
			log.Error("db seems offline", zap.Error(err))
			return e.DownErr
		} else if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, sql.ErrNoRows) {
			return e.NoDataErr
		} else {
			log.Error("unexpected error in db", zap.Error(err))
			return e.UnexpectedErr
		}
	}
	return err
}

package svc

import "errors"

var (
	ErrUnexpected            = errors.New("unexpected error")
	ErrNoData                = errors.New("no data found")
	ErrBadInput              = errors.New("invalid data found")
	ErrImportantFieldMissing = errors.New("important data field missing")
	ErrDuplicateData         = errors.New("duplicate data")
	ErrTimeout               = errors.New("timeout error")
)

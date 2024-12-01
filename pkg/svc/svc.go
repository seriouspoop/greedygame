package svc

import (
	"seriouspoop/greedygame/go-common/logging"
)

const (
	maxRangeNumber1 = 100
	minRangeNumber1 = 10
	maxRangeNumber2 = 10
	minRangeNumber2 = 1
)

type Svc struct {
	db     dbHelper
	logger *logging.Logger
	health error
}

func New(db dbHelper, logger *logging.Logger) *Svc {
	return &Svc{db: db, logger: logger}
}

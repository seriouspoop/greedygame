package svc

import (
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/go-common/observer"
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
	tracer *observer.Tracer
	health error
}

func New(db dbHelper, logger *logging.Logger, tracer *observer.Tracer) *Svc {
	return &Svc{db: db, logger: logger, tracer: tracer}
}

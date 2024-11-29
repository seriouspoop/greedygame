package svc

import (
	"github.com/rs/zerolog"
)

const (
	maxRangeNumber1 = 100
	minRangeNumber1 = 10
	maxRangeNumber2 = 10
	minRangeNumber2 = 1
)

type Svc struct {
	logger *zerolog.Logger
	health error
}

func New(logger *zerolog.Logger) *Svc {
	return &Svc{logger: logger}
}

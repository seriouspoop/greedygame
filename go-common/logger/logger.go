package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func NewWithService(service, logLevel string) (*zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}

	log := zerolog.New(os.Stderr).
		Level(level).
		With().
		Caller().
		Timestamp().
		Str("service", service).
		Logger()

	log = log.Hook(NewTraceHook())
	return &log, nil
}

func New(logLevel zerolog.Level) *zerolog.Logger {
	log := zerolog.New(os.Stderr).
		Level(logLevel).
		With().
		Caller().
		Timestamp().
		Logger()

	log = log.Hook(NewTraceHook())
	return &log
}

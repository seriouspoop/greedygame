package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func New(service, logLevel string) (*zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}

	log := zerolog.New(os.Stderr).
		Level(level).
		With().
		Timestamp().
		Str("service", service).
		Logger()

	l := &Logger{log}
	return &l.logger, nil
}

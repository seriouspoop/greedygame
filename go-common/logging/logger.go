package logging

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func (l *Logger) WithCtxLogger(ctx context.Context) *zerolog.Logger {
	lg := l.logger.With().Ctx(ctx).Logger()
	return &lg
}

func NewWithService(service, logLevel string) (*Logger, error) {
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
	l := &Logger{&log}
	return l, nil
}

func New(logLevel zerolog.Level) *Logger {
	log := zerolog.New(os.Stderr).
		Level(logLevel).
		With().
		Caller().
		Timestamp().
		Logger()

	log = log.Hook(NewTraceHook())
	l := &Logger{&log}
	return l
}

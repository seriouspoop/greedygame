package logging

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}

func (l *Logger) Logger() *zap.Logger {
	return l.logger
}

func NewWithService(service, logLevel string, core ...zapcore.Core) (*Logger, error) {
	level, err := zap.ParseAtomicLevel(logLevel)
	if err != nil {
		return nil, err
	}

	config := zap.Config{
		Level:             level,
		Encoding:          "json",
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields:     map[string]interface{}{"service": service},
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		DisableStacktrace: true,
	}

	if level.Level() == zap.DebugLevel {
		config.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	if len(core) > 0 {
		otelCore := zap.WrapCore(func(c zapcore.Core) zapcore.Core {
			core = append(core, c)
			return zapcore.NewTee(core...)
		})
		log = log.WithOptions(otelCore)
	}
	return &Logger{log}, nil
}

func New(level zapcore.Level) *zap.Logger {

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	if level == zap.DebugLevel {
		config.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	log := zap.Must(config.Build())
	return log
}

func NewTestLogger(ctx context.Context) *Logger {
	log := zap.NewNop()
	return &Logger{log}
}

package observer

import (
	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel/sdk/log"
)

type Logger struct {
	name string
}

func NewLogger(name string) *Logger {
	return &Logger{}
}

func (l *Logger) NewLoggerCore() *otelzap.Core {
	provider := log.NewLoggerProvider()
	core := otelzap.NewCore(l.name, otelzap.WithLoggerProvider(provider))
	return core
}

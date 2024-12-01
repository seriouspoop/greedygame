package logging

import (
	"seriouspoop/greedygame/go-common/globals"

	"github.com/rs/zerolog"
)

type tracingHook struct{}

func NewTraceHook() *tracingHook {
	return &tracingHook{}
}

func (t *tracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	traceId := ctx.Value(globals.TraceIDContextKey).(string)
	e.Str("trace-id", traceId)
}

package logging

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func (l *Logger) Ctx(ctx context.Context) *zap.Logger {
	span := trace.SpanFromContext(ctx)
	traceID := zap.String("trace_id", span.SpanContext().TraceID().String())
	spanID := zap.String("span_id", span.SpanContext().SpanID().String())
	fields := []zap.Field{traceID, spanID}
	return l.logger.With(fields...)
}

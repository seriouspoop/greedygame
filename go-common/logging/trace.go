package logging

import (
	"context"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerWithCtx struct {
	ctx context.Context
	l   *Logger
}

func (l *Logger) Ctx(ctx context.Context) *LoggerWithCtx {
	span := trace.SpanContextFromContext(ctx)
	traceID := zap.String("trace_id", span.TraceID().String())
	spanID := zap.String("span_id", span.SpanID().String())
	fields := []zap.Field{traceID, spanID}
	return &LoggerWithCtx{ctx, &Logger{l.logger.With(fields...)}}
}

func (l *LoggerWithCtx) Debug(msg string, fields ...zap.Field) {
	if span := trace.SpanFromContext(l.ctx); span.IsRecording() {
		span.AddEvent(msg)
	}
	l.l.logger.Debug(msg, fields...)
}

func (l *LoggerWithCtx) Info(msg string, fields ...zap.Field) {
	if span := trace.SpanFromContext(l.ctx); span.IsRecording() {
		span.AddEvent(msg)
	}
	l.l.logger.Info(msg, fields...)
}

func (l *LoggerWithCtx) Error(msg string, fields ...zap.Field) {
	if span := trace.SpanFromContext(l.ctx); span.IsRecording() {
		span.SetStatus(codes.Error, msg)
	}
	l.l.logger.Error(msg, fields...)
}

func (l *LoggerWithCtx) Log(level zapcore.Level, msg string, fields ...zap.Field) {
	l.l.logger.Log(level, msg, fields...)
}

func (l *LoggerWithCtx) With(fields ...zap.Field) *LoggerWithCtx {
	log := l.l.logger.With(fields...)
	return &LoggerWithCtx{l.ctx, &Logger{logger: log}}
}

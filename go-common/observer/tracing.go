package observer

import (
	"context"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

type Tracer struct {
	name     string
	tracer   trace.Tracer
	provider *sdktrace.TracerProvider
}

// returns a no-op tracer, to be used for testing.
// Provider is nil, Shutdown is not required.
func NewNoopTracer() *Tracer {
	provider := noop.NewTracerProvider()
	tracer := provider.Tracer("test/tracer")

	// provider set to nil for testing
	return &Tracer{name: "test/tracer", tracer: tracer, provider: nil}
}

func NewTracer(ctx context.Context, name string, ex Exporter) (*Tracer, error) {
	// setup exporter based on service choice
	e, err := setupTraceExporter(ctx, ex)
	if err != nil {
		return nil, err
	}
	provider, err := newTraceProvider(e, name)
	if err != nil {
		return nil, err
	}
	// set global provider for all libraries using otel - otelmux etc.
	otel.SetTracerProvider(provider)
	tracer := provider.Tracer(name)
	return &Tracer{name: name, tracer: tracer, provider: provider}, nil
}

// Wrapper ---------------------------------------

// return name of the tracer in use.
func (t *Tracer) Name() string {
	return t.name
}

func (t *Tracer) Shutdown(ctx context.Context) error {
	return t.provider.Shutdown(ctx)
}

func (t *Tracer) Start(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return t.tracer.Start(ctx, spanName)
}

// returns mock context, mimics contexts produced by tracers.
// To be used for testing.
func (t *Tracer) MockContext(ctx context.Context) context.Context {
	return trace.ContextWithSpan(ctx, noop.Span{})
}

// Middleware -----------------------------------------

func (t *Tracer) TraceHTTPMiddleware(next http.Handler) http.Handler {
	// use otelmux instrumentation
	return otelmux.Middleware(t.name).Middleware(next)
}

// SDK -------------------------------------------

func newTraceProvider(exp sdktrace.SpanExporter, name string) (*sdktrace.TracerProvider, error) {
	r, err := newResource(name)
	if err != nil {
		return nil, err
	}

	// TODO - add sampling option
	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	), nil
}

// Exporters -----------------------------------------

func setupTraceExporter(ctx context.Context, ex Exporter) (e sdktrace.SpanExporter, err error) {
	switch ex {
	case ConsoleExporter:
		e, err = newTraceConsoleExporter()
	case OTLPHttpExporter:
		e, err = newTraceOTLPHttpExporter(ctx)
	case OTLPGrpcExporter:
		e, err = newTraceOTLPGrpcExporter(ctx)
	default:
		e, err = newTraceConsoleExporter()
	}
	if err != nil {
		return nil, err
	}
	return
}

func newTraceOTLPHttpExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return otlptracehttp.New(ctx)
}

func newTraceOTLPGrpcExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return otlptracegrpc.New(ctx)
}

func newTraceConsoleExporter() (sdktrace.SpanExporter, error) {
	return stdouttrace.New()
}

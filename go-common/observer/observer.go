package observer

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type Exporter int

const (
	ConsoleExporter Exporter = iota
	OTLPExporter
)

type Observer struct {
	name         string
	traceSDK     *Tracer
	logSDK       *Logger
	exporterType Exporter
}

func New(ctx context.Context, name string, ex Exporter) (*Observer, error) {
	tracer, err := NewTracer(ctx, name, ex)
	if err != nil {
		return nil, err
	}

	return &Observer{
		name:         name,
		traceSDK:     tracer,
		logSDK:       NewLogger(name),
		exporterType: ex,
	}, nil
}

func (o *Observer) Shutdown(ctx context.Context) error {
	return o.traceSDK.Shutdown(ctx)
}

// returns logging SDK of the observer, use to extract log related instrumentation
func (o *Observer) LogSDK() *Logger {
	return o.logSDK
}

// returns tracing SDK of observer, use to extract trace related instrumentation
func (o *Observer) TraceSDK() *Tracer {
	return o.traceSDK
}

func newResource(name string) (*resource.Resource, error) {
	return resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(name),
		),
	)
}

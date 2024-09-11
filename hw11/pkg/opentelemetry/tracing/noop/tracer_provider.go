package noopsdktracerprovider

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
	"go.opentelemetry.io/otel/trace/noop"
)

type TracerProvider struct {
	embedded.TracerProvider
	tp noop.TracerProvider
}

func NewTracerProvider() *TracerProvider {
	return &TracerProvider{
		tp: noop.NewTracerProvider(),
	}
}

func (t *TracerProvider) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	return t.tp.Tracer(name, opts...)
}

func (t *TracerProvider) Shutdown(ctx context.Context) error {
	return nil
}

package tracing

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	noopsdktracerprovider "zoo/pkg/opentelemetry/tracing/noop"
	"zoo/pkg/opentelemetry/tracing/tracer_provider"
)

type TracerProvider interface {
	trace.TracerProvider
	Shutdown(ctx context.Context) error
}

func NewTracerProvider(ctx context.Context, cfg Config) (TracerProvider, error) {
	var (
		tracerProvider TracerProvider
		err            error
	)
	if cfg.Enabled {
		tracerProvider, err = sdktracerprovider.NewTracerProvider(ctx, cfg.SdkConfig)
		if err != nil {
			return nil, err
		}
	} else {
		tracerProvider = noopsdktracerprovider.NewTracerProvider()
	}

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))
	otel.SetTracerProvider(tracerProvider)

	return tracerProvider, nil
}

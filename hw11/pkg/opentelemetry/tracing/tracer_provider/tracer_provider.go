package sdktracerprovider

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
	"os"
)

type TracerProvider struct {
	embedded.TracerProvider
	tp       *sdkTrace.TracerProvider
	exporter sdkTrace.SpanExporter
}

func NewTracerProvider(ctx context.Context, cfg Config) (*TracerProvider, error) {
	var tracerProvider TracerProvider
	exporter, err := NewGrpcExporter(ctx, cfg.CollectorEndpoint)
	if err != nil {
		return nil, err
	}

	tracerProvider.exporter = exporter

	tracerProvider.tp = sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(exporter, sdkTrace.WithBatchTimeout(cfg.BatchTimeout)),
		sdkTrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(cfg.Service.ServiceName),
			semconv.DeploymentEnvironment(cfg.Service.DeploymentEnv),
			semconv.HostName(getHostname()),
		)),
	)

	return &tracerProvider, nil
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

func (t *TracerProvider) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	return t.tp.Tracer(name, opts...)
}

func (t *TracerProvider) Shutdown(ctx context.Context) error {
	var err error
	err = t.tp.Shutdown(ctx)
	if err != nil {
		err = fmt.Errorf("failed to shutdown tracer provider: %w", err)
	}

	if t.exporter != nil {
		errExporter := t.exporter.Shutdown(ctx)
		if errExporter != nil {
			errExporter = fmt.Errorf("failed to shutdown tracer exporter: %w", errExporter)
			if err != nil {
				err = fmt.Errorf("%w; %w", err, errExporter)
			} else {
				err = errExporter
			}
		}
	}

	return err
}

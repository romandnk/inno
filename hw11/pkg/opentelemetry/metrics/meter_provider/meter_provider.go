package sdkmeterprovider

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	sdkMeter "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"os"
	"time"
)

type MeterProvider struct {
	embedded.MeterProvider
	mp       *sdkMeter.MeterProvider
	exporter sdkMeter.Exporter
}

func NewMeterProvider(ctx context.Context, cfg Config) (*MeterProvider, error) {
	var meterProvider MeterProvider
	exporter, err := NewExporter(ctx, cfg.ExporterConfig)
	if err != nil {
		return nil, err
	}

	meterProvider.exporter = exporter

	reader := sdkMeter.NewPeriodicReader(
		exporter,
		sdkMeter.WithInterval(time.Duration(cfg.ExporterConfig.ExportInterval)*time.Millisecond),
		sdkMeter.WithTimeout(time.Duration(cfg.ExporterConfig.ExportTimeout)*time.Millisecond),
	)

	meterProvider.mp = sdkMeter.NewMeterProvider(
		sdkMeter.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(cfg.ServiceConfig.ServiceName),
			semconv.DeploymentEnvironment(cfg.ServiceConfig.DeploymentEnv),
			semconv.HostName(getHostname()),
		)),
		sdkMeter.WithReader(reader),
	)

	if cfg.RuntimeMetricsConfig.Enabled {
		err = runtime.Start(
			runtime.WithMinimumReadMemStatsInterval(cfg.RuntimeMetricsConfig.MinimumReadMemStatsInterval),
			runtime.WithMeterProvider(meterProvider.mp),
		)
		if err != nil {
			return nil, err
		}
	}

	return &meterProvider, nil
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

func (m *MeterProvider) Meter(name string, opts ...metric.MeterOption) metric.Meter {
	return m.mp.Meter(name, opts...)
}

func (m *MeterProvider) Shutdown(ctx context.Context) error {
	var err error
	err = m.mp.Shutdown(ctx)
	if err != nil {
		err = fmt.Errorf("failed to shutdown meter provider: %w", err)
	}

	if m.exporter != nil {
		errExporter := m.exporter.Shutdown(ctx)
		if errExporter != nil {
			errExporter = fmt.Errorf("failed to shutdown meter exporter: %w", errExporter)
			if err != nil {
				err = fmt.Errorf("%w; %w", err, errExporter)
			} else {
				err = errExporter
			}
		}
	}

	return err
}

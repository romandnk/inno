package sdkmeterprovider

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkMeter "go.opentelemetry.io/otel/sdk/metric"
)

type ExporterConfig struct {
	CollectorEndpoint string `env:"OTEL_EXPORTER_OTLP_METRICS_ENDPOINT"`
	Compression       string `env:"OTEL_EXPORTER_OTLP_METRICS_COMPRESSION"`
	ExportInterval    int    `env:"OTEL_METRIC_EXPORT_INTERVAL"`
	ExportTimeout     int    `env:"OTEL_METRIC_EXPORT_TIMEOUT"`
}

func NewExporter(ctx context.Context, config ExporterConfig) (sdkMeter.Exporter, error) {
	return otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(config.CollectorEndpoint),
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithCompressor(config.Compression),
	)
}

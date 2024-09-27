package metrics

import (
	sdkmeterprovider "zoo/pkg/opentelemetry/metrics/meter_provider"
)

type Config struct {
	Enabled   bool `env:"OTEL_METRICS_ENABLED"`
	SdkConfig sdkmeterprovider.Config
}

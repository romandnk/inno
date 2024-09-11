package tracing

import sdktracerprovider "zoo/pkg/opentelemetry/tracing/tracer_provider"

type Config struct {
	Enabled   bool `env:"OTEL_TRACES_ENABLED"`
	SdkConfig sdktracerprovider.Config
}

package sdkmeterprovider

import "time"

type Config struct {
	RuntimeMetricsConfig RuntimeMetricsConfig
	ExporterConfig       ExporterConfig
	ServiceConfig        ServiceConfig
}

type RuntimeMetricsConfig struct {
	Enabled                     bool          `env:"OTEL_RUNTIME_METRICS_ENABLED"`
	MinimumReadMemStatsInterval time.Duration `env:"OTEL_RUNTIME_METRICS_MIN_READ_MEM_STATS_INTERVAL"`
}

type ServiceConfig struct {
	DeploymentEnv string `env:"DEPLOYMENT_ENVIRONMENT"`
	ServiceName   string `env:"SERVICE_NAME"`
}

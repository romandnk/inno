package sdktracerprovider

import "time"

type Config struct {
	CollectorEndpoint string        `env:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	BatchTimeout      time.Duration `env:"OTEL_TRACES_BATCH_TIMEOUT"`
	Service           ServiceConfig
}

type ServiceConfig struct {
	DeploymentEnv string `env:"DEPLOYMENT_ENVIRONMENT"`
	ServiceName   string `env:"SERVICE_NAME"`
}

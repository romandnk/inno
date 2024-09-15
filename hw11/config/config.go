package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
	"zoo/pkg/opentelemetry/metrics"
	"zoo/pkg/opentelemetry/tracing"
	httpserver "zoo/pkg/server/http"
	"zoo/pkg/storage/postgres"
)

type Config struct {
	RequestPerUser  int           `env:"REQUEST_PER_USER" env-default:"100"`
	RateLimitWindow time.Duration `env:"RATE_WINDOW" env-default:"1s"`
	HTTPServer      httpserver.Config
	Postgres        postgres.Config
	Tracer          tracing.Config
	Metrics         metrics.Config
}

func New() (Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}

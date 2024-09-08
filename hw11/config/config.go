package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
	httpserver "zoo/pkg/server/http"
	"zoo/pkg/storage/postgres"
)

type Config struct {
	RequestPerUser  int           `env:"REQUEST_PER_USER" env-default:"3"`
	RateLimitWindow time.Duration `env:"RATE_WINDOW" env-default:"10s"`
	HTTPServer      httpserver.Config
	Postgres        postgres.Config
}

func New() (Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}

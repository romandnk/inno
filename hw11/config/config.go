package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	httpserver "inno/hw11/pkg/server/http"
	"time"
)

type Config struct {
	RequestPerUser  int           `env:"REQUEST_PER_USER" env-default:"3"`
	RateLimitWindow time.Duration `env:"RATE_WINDOW" env-default:"10s"`
	HTTPServer      httpserver.Config
}

func NewConfig() (Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}

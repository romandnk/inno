package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	WorkerNum       int    `env:"WORKER_NUM" env-default:"3"`
	ValidationToken string `env:"VALIDATION_TOKEN" env-default:"secret"`
}

func New() (Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("error reading envs: %w", err)
	}

	err = cleanenv.UpdateEnv(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("error updating envs: %w", err)
	}

	return cfg, nil
}

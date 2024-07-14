package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"inno/attestation1/internal/token"
	"inno/attestation1/internal/worker"
	"time"
)

type Config struct {
	Worker          worker.Config `yaml:"worker"`
	Token           token.Config  `yaml:"token"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout"`
}

func NewConfig() (Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

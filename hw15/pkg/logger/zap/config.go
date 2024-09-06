package zaplogger

import (
	"fmt"
	"strings"

	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level  string `env:"LOG_LEVEL" env-default:"info"`
	Format string `env:"LOG_FORMAT" env-default:"json"`
	Env    string `env:"LOG_ENV" env-default:"dev"`
}

func (c Config) Validate() error {
	err := validateConfigLevel(c.Level)
	if err != nil {
		return err
	}

	err = validateConfigFormat(c.Format)
	if err != nil {
		return err
	}

	err = validateConfigEnv(c.Env)
	if err != nil {
		return err
	}

	return nil
}

func validateConfigLevel(level string) error {
	_, err := zapcore.ParseLevel(level)
	return err
}

func validateConfigFormat(format string) error {
	switch strings.ToLower(format) {
	case "json", "console":
		return nil
	default:
		return fmt.Errorf("invalid logger format: %s", format)
	}
}

func validateConfigEnv(env string) error {
	switch strings.ToLower(env) {
	case "local", "dev", "prod":
		return nil
	default:
		return fmt.Errorf("invalid logger env: %s", env)
	}
}

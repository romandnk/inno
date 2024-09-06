package appfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"chat/pkg/config"
	zaplogger "chat/pkg/logger/zap"
)

func LoggerModule() fx.Option {
	return fx.Module("logger",
		fx.Provide(
			fx.Private,
			parseLoggerConfig,
		),
		fx.Provide(
			func(cfg zaplogger.Config) (*zap.Logger, error) {
				logger, err := zaplogger.New(cfg)
				if err != nil {
					return nil, err
				}
				return logger, nil
			},
		),
	)
}

func parseLoggerConfig() (zaplogger.Config, error) {
	var cfg zaplogger.Config
	err := config.ParseEnv(&cfg)
	if err != nil {
		return zaplogger.Config{}, err
	}

	return cfg, nil
}

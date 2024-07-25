package app

import (
	"go.uber.org/fx"
	"inno/hw11/config"
)

func ConfigModule() fx.Option {
	return fx.Options(
		fx.Provide(
			config.NewConfig,
		),
	)
}

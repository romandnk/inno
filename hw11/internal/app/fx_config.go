package app

import (
	"go.uber.org/fx"
	"zoo/config"
)

func ConfigModule() fx.Option {
	return fx.Options(
		fx.Provide(
			config.New,
		),
	)
}

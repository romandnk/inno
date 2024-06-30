package app

import (
	"go.uber.org/fx"
	"inno/attestation1/config"
)

func ConfigModule() fx.Option {
	return fx.Module("config",
		fx.Provide(
			config.New,
		),
	)
}

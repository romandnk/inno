package app

import "go.uber.org/fx"

func App() fx.Option {
	return fx.Options(
		ConfigModule(),
	)
}

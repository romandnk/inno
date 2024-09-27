package app

import (
	"context"
	"go.uber.org/fx"
)

func NewApp() fx.Option {
	return fx.Options(
		fx.Provide(
			func() context.Context {
				return context.Background()
			},
		),

		ConfigModule(),

		DBModule(),
		RepoModule(),
		CacheModule(),

		TracingModule(),
		MetricsModule(),

		HTTPHandlerModule(),
		HTTPServerModule(),
	)
}

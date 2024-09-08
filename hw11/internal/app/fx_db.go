package app

import (
	"context"
	"go.uber.org/fx"
	"zoo/config"
	"zoo/pkg/storage/postgres"
)

func DBModule() fx.Option {
	return fx.Options(
		fx.Provide(
			func(cfg config.Config) postgres.Config {
				return cfg.Postgres
			},
			postgres.New,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, pg *postgres.Postgres) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return pg.Connect(ctx)
					},
					OnStop: func(ctx context.Context) error {
						pg.Close()
						return nil
					},
				})
			},
		),
	)
}

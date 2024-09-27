package app

import (
	"go.uber.org/fx"
	"zoo/internal/cache"
	"zoo/pkg/storage/inmem"
)

func CacheModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				inmem.New,
				fx.As(new(cache.Cache)),
			),
		),
	)
}

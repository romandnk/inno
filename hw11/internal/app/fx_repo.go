package app

import (
	"go.uber.org/fx"
	"zoo/internal/repository"
)

func RepoModule() fx.Option {
	return fx.Options(
		fx.Provide(
			repository.New,
		),
	)
}

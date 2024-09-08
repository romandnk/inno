package app

import "go.uber.org/fx"

func NewApp() fx.Option {
	return fx.Options(
		ConfigModule(),
		DBModule(),
		RepoModule(),
		CacheModule(),
		HTTPHandlerModule(),
		HTTPServerModule(),
	)
}

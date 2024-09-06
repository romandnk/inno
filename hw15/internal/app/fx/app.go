package appfx

import (
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Options(
		LoggerModule(),
		ServiceModule(),
		HttpServerModule(),
	)
}

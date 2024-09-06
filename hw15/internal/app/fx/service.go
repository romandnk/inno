package appfx

import (
	"chat/internal/service"
	"go.uber.org/fx"
)

func ServiceModule() fx.Option {
	return fx.Module("service",
		fx.Provide(
			service.New,
		),
	)
}

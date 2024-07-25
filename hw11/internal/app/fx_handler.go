package app

import (
	"go.uber.org/fx"
	"inno/hw11/config"
	v1 "inno/hw11/internal/http/v1"
	"net/http"
	"time"
)

func HTTPHandlerModule() fx.Option {
	return fx.Options(
		fx.Provide(
			func(cfg config.Config) (int, time.Duration) {
				return cfg.RequestPerUser, cfg.RateLimitWindow
			},
			fx.Annotate(
				func(requestNumPerUser int, rateLimitWindow time.Duration) *http.ServeMux {
					handler := v1.NewHandler(requestNumPerUser, rateLimitWindow)
					return handler.InitRoutes()
				},
				fx.As(new(http.Handler)),
			),
		),
	)
}

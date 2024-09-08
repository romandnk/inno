package app

import (
	"go.uber.org/fx"
	"net/http"
	"time"
	"zoo/config"
	v1 "zoo/internal/http/v1"
	"zoo/internal/repository"
)

func HTTPHandlerModule() fx.Option {
	return fx.Options(
		fx.Provide(
			func(cfg config.Config) (int, time.Duration) {
				return cfg.RequestPerUser, cfg.RateLimitWindow
			},
			fx.Annotate(
				func(requestNumPerUser int, rateLimitWindow time.Duration, repo *repository.Repository) *http.ServeMux {
					handler := v1.NewHandler(
						requestNumPerUser,
						rateLimitWindow,
						repo,
					)
					return handler.InitRoutes()
				},
				fx.As(new(http.Handler)),
			),
		),
	)
}

package appfx

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"

	v1 "chat/internal/controller/websocket/v1"
	"chat/pkg/config"
	httpserver "chat/pkg/server/http"
)

func HttpServerModule() fx.Option {
	return fx.Module("httpServer",
		// provide config
		fx.Provide(
			fx.Private,
			parseHttpServerConfigConfig,
		),
		// provide http handler
		fx.Provide(
			fx.Private,
			v1.NewHandler,
			fx.Annotate(
				func(handler *v1.Handler) *gin.Engine {
					return handler.InitRoutes()
				},
				fx.As(new(http.Handler)),
			),
		),
		// provide http server
		fx.Provide(
			httpserver.NewServer,
		),
		// start and stop http server
		fx.Invoke(
			func(lc fx.Lifecycle, srv *httpserver.Server, logger *zap.Logger) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							if err := srv.Start(); err != nil {
								logger.Error("error starting HTTP server", zap.String("address", srv.GetAddr()))
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return srv.Stop(ctx)
					},
				})
			},
		),
	)
}

func parseHttpServerConfigConfig() (httpserver.Config, error) {
	var cfg httpserver.Config
	err := config.ParseEnv(&cfg)
	if err != nil {
		return httpserver.Config{}, err
	}

	return cfg, nil
}

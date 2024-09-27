package app

import (
	"context"
	"go.uber.org/fx"
	"log/slog"
	"net"
	"strconv"
	"zoo/config"
	httpserver "zoo/pkg/server/http"
)

func HTTPServerModule() fx.Option {
	return fx.Module("http server",
		fx.Provide(
			func(cfg config.Config) httpserver.Config {
				return cfg.HTTPServer
			},
			httpserver.NewServer,
		),
		fx.Invoke(func(lc fx.Lifecycle, srv *httpserver.Server, cfg httpserver.Config) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := srv.Start(); err != nil {
							slog.ErrorContext(ctx, "error starting HTTP server",
								slog.String("address", net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))),
							)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return srv.Stop(ctx)
				},
			})
		}),
	)
}

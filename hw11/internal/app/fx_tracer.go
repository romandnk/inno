package app

import (
	"context"
	"go.uber.org/fx"
	"zoo/config"
	"zoo/pkg/opentelemetry/tracing"
)

func TracingModule() fx.Option {
	return fx.Options(
		fx.Provide(
			func(cfg config.Config) tracing.Config {
				return cfg.Tracer
			},
		),
		fx.Provide(
			tracing.NewTracerProvider,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, tracerProvider tracing.TracerProvider) error {
				lc.Append(fx.Hook{
					OnStop: func(ctx context.Context) error {
						return tracerProvider.Shutdown(ctx)
					},
				})
				return nil
			},
		),
	)
}

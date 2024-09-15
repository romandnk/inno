package app

import (
	"context"
	"go.uber.org/fx"
	"zoo/config"
	"zoo/pkg/opentelemetry/metrics"
)

func MetricsModule() fx.Option {
	return fx.Options(
		fx.Provide(
			func(cfg config.Config) metrics.Config {
				return cfg.Metrics
			},
		),
		fx.Provide(
			metrics.NewMeterProvider,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, meterProvider metrics.MeterProvider) error {
				lc.Append(fx.Hook{
					OnStop: func(ctx context.Context) error {
						return meterProvider.Shutdown(ctx)
					},
				})
				return nil
			},
		),
	)
}

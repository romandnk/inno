package metrics

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmeterprovider "zoo/pkg/opentelemetry/metrics/meter_provider"
	noopsdkmeterprovider "zoo/pkg/opentelemetry/metrics/noop"
)

type MeterProvider interface {
	metric.MeterProvider
	Shutdown(ctx context.Context) error
}

func NewMeterProvider(ctx context.Context, cfg Config) (MeterProvider, error) {
	var (
		meterProvider MeterProvider
		err           error
	)
	if cfg.Enabled {
		meterProvider, err = sdkmeterprovider.NewMeterProvider(ctx, cfg.SdkConfig)
		if err != nil {
			return nil, err
		}
	} else {
		meterProvider = noopsdkmeterprovider.NewMeterProvider()
	}

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))
	otel.SetMeterProvider(meterProvider)

	return meterProvider, nil
}

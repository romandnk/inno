package noopsdkmeterprovider

import (
	"context"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/metric/noop"
)

type MeterProvider struct {
	embedded.MeterProvider
	mp noop.MeterProvider
}

func NewMeterProvider() *MeterProvider {
	return &MeterProvider{
		mp: noop.NewMeterProvider(),
	}
}

func (m *MeterProvider) Meter(name string, opts ...metric.MeterOption) metric.Meter {
	return m.mp.Meter(name, opts...)
}

func (m *MeterProvider) Shutdown(ctx context.Context) error {
	return nil
}

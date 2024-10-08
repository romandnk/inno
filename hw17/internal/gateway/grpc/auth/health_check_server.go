package auth

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthChecker struct {
	logger *slog.Logger
}

func NewHealthChecker(logger *slog.Logger) *HealthChecker {
	return &HealthChecker{logger: logger.With("module", "health-checker")}
}

func (s *HealthChecker) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	s.logger.Debug("Serving the Check request for health check")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthChecker) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	s.logger.Debug("Serving the Watch request for health check")
	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}

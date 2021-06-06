package handler

import (
	"context"

	pb "github.com/kzmake/distributed-calculator/common/health/api/health/v1"
)

var Status = pb.HealthCheckResponse_NOT_SERVING.Enum()

type health struct {
	pb.UnimplementedHealthServer
}

var _ pb.HealthServer = new(health)

func NewHaelth() pb.HealthServer { return &health{} }

func (h *health) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: *Status}, nil
}

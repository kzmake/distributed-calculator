package health

import (
	"google.golang.org/grpc"

	pb "github.com/kzmake/distributed-calculator/api/health/v1"
	"github.com/kzmake/distributed-calculator/common/health/handler"
)

func Serving() {
	handler.Status = pb.HealthCheckResponse_SERVING.Enum()
}

func NotServing() {
	handler.Status = pb.HealthCheckResponse_NOT_SERVING.Enum()
}

func NewHealthServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterHealthServer(s, handler.NewHaelth())

	return s
}

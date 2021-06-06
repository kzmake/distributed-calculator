package handler

import (
	"context"

	pb "github.com/kzmake/distributed-calculator/microservices/adder/api/adder/v1"
)

type adder struct {
	pb.UnimplementedAdderServer
}

var _ pb.AdderServer = new(adder)

func NewAdder() pb.AdderServer { return &adder{} }

func (h *adder) Add(ctx context.Context, req *pb.AddRequest) (*pb.Addresponse, error) {
	return &pb.Addresponse{Result: req.OperandOne + req.OperandTwo}, nil
}

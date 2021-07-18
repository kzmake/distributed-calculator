package handler

import (
	"context"
	"encoding/json"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"golang.org/x/xerrors"

	adder "github.com/kzmake/distributed-calculator/api/adder/v1"
	pb "github.com/kzmake/distributed-calculator/api/subtractor/v1"
)

const (
	appID  = "adder"
	method = "add"
)

type subtractor struct {
	pb.UnimplementedSubtractorServer
}

var _ pb.SubtractorServer = new(subtractor)

func NewSubtractor() pb.SubtractorServer {
	return &subtractor{}
}

func (h *subtractor) Sub(ctx context.Context, req *pb.SubRequest) (*pb.Subresponse, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, xerrors.Errorf("failed to genereate client: %w", err)
	}

	in, err := json.Marshal(&adder.AddRequest{
		OperandOne: req.OperandOne,
		OperandTwo: req.OperandTwo * -1.0,
	})
	if err != nil {
		return nil, xerrors.Errorf("failed to marshal request: %w", err)
	}

	out, err := client.InvokeMethodWithContent(ctx, appID, method, http.MethodPost,
		&dapr.DataContent{ContentType: "application/json", Data: in},
	)

	var res *adder.AddResponse
	if jerr := json.Unmarshal(out, &res); jerr != nil {
		return nil, xerrors.Errorf("failed to unmarshal response: %w", err)
	}

	return &pb.Subresponse{Result: res.Result}, nil
}

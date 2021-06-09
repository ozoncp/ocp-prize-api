package api

import (
	"context"
	"os"

	desc "github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOcpPrizeApiServer
	log zerolog.Logger
}

func (a *api) Init() {
	a.log = zerolog.New(os.Stdout)
}

func NewOcpPrizeApi() desc.OcpPrizeApiServer {
	return &api{}
}

func (a *api) CreatePrizeV1(
	ctx context.Context,
	req *desc.CreatePrizeV1Request,
) (*desc.CreatePrizeV1Response, error) {
	a.log.Printf("CreatePrizeV1 request: %s", req.String())
	return nil, nil
}

func (a *api) DescribePrizeV1(
	ctx context.Context,
	req *desc.DescribePrizeV1Request,
) (*desc.DescribePrizeV1Response, error) {
	a.log.Printf("DescribePrizeV1 request: %s", req.String())
	if err := req.Validate(); err != nil {
		a.log.Printf("invalid request: %s, error: %s", req.String(), err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return nil, nil
}

func (a *api) ListPrizeV1(
	ctx context.Context,
	req *desc.ListPrizeV1Request,
) (*desc.ListPrizeV1Response, error) {
	a.log.Printf("ListPrizeV1 request: %s", req.String())
	return nil, nil
}

func (a *api) RemovePrizeV1(
	ctx context.Context,
	req *desc.RemovePrizeV1Request,
) (*desc.RemovePrizeV1Response, error) {
	if err := req.Validate(); err != nil {
		a.log.Printf("invalid request: %s, error: %s", req.String(), err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	a.log.Printf("RemovePrizeV1 request: %s", req.String())
	return nil, nil
}

package api

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-prize-api/internal/prize"
	"github.com/ozoncp/ocp-prize-api/internal/repo"
	desc "github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type API struct {
	desc.UnimplementedOcpPrizeApiServer
	currentRepo repo.IRepo
}

func NewOcpPrizeApi(db *sqlx.DB) desc.OcpPrizeApiServer {
	return &API{
		currentRepo: repo.NewRepo(db),
	}
}

func (a *API) CreatePrizeV1(
	ctx context.Context,
	req *desc.CreatePrizeV1Request,
) (*desc.CreatePrizeV1Response, error) {

	log.Printf("CreatePrizeV1 request: %s", req.String())
	prizeToAdd := prize.Prize{
		Link:    req.Link,
		IssueID: req.IssueId,
	}
	id, err := a.currentRepo.AddPrizes(ctx, []prize.Prize{prizeToAdd})
	if err != nil {
		log.Printf("CreatePrizeV1 error: %s", err.Error())
		return nil, err
	}
	response := desc.CreatePrizeV1Response{
		PrizeId: id,
	}
	return &response, nil
}

func (a *API) DescribePrizeV1(
	ctx context.Context,
	req *desc.DescribePrizeV1Request,
) (*desc.DescribePrizeV1Response, error) {
	log.Printf("DescribePrizeV1 request: %s", req.String())
	if err := req.Validate(); err != nil {
		log.Printf("invalid request: %s, error: %s", req.String(), err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	prize, err := a.currentRepo.DescribePrize(ctx, req.PrizeId)
	if err != nil {
		log.Printf("DescribePrizeV1 error: %s", err.Error())
		return nil, err
	}
	response := desc.DescribePrizeV1Response{
		Prize: &desc.Prize{
			Id:      prize.ID,
			IssueId: prize.IssueID,
			Link:    prize.Link,
		},
	}
	return &response, nil
}

func (a *API) ListPrizeV1(
	ctx context.Context,
	req *desc.ListPrizeV1Request,
) (*desc.ListPrizeV1Response, error) {
	log.Printf("ListPrizeV1 request: %s", req.String())
	prizes, err := a.currentRepo.ListPrizes(ctx, req.Limit, req.Offset)
	if err != nil {
		log.Printf("ListPrizeV1 error: %s", err.Error())
		return nil, err
	}
	var response desc.ListPrizeV1Response
	for _, prize := range prizes {
		response.Prizes = append(response.Prizes, &desc.Prize{
			Id:      prize.ID,
			IssueId: prize.IssueID,
			Link:    prize.Link,
		})
	}
	return &response, nil
}

func (a *API) RemovePrizeV1(
	ctx context.Context,
	req *desc.RemovePrizeV1Request,
) (*desc.RemovePrizeV1Response, error) {
	log.Printf("RemovePrizeV1 request: %s", req.String())
	if err := req.Validate(); err != nil {
		log.Printf("invalid request: %s, error: %s", req.String(), err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	found, err := a.currentRepo.RemovePrize(ctx, req.PrizeId)
	if err != nil {
		log.Printf("RemovePrizeV1 error: %s", err.Error())
		return nil, err
	}
	response := desc.RemovePrizeV1Response{
		Found: found,
	}
	return &response, nil
}

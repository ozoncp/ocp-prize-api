package api

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-prize-api/internal/configuration"
	"github.com/ozoncp/ocp-prize-api/internal/flusher"
	"github.com/ozoncp/ocp-prize-api/internal/metrics"
	"github.com/ozoncp/ocp-prize-api/internal/prize"
	"github.com/ozoncp/ocp-prize-api/internal/producer"
	"github.com/ozoncp/ocp-prize-api/internal/repo"
	desc "github.com/ozoncp/ocp-prize-api/pkg/ocp-prize-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var flusherMaximumChankSize int = 100

// API struct for implementation ocp prize api
type API struct {
	desc.UnimplementedOcpPrizeApiServer
	currentRepo repo.IRepo
	producer    producer.IProducer
}

// NewOcpPrizeApi creates new api with setted database and producer
func NewOcpPrizeApi(ctx context.Context, db *sqlx.DB, producer producer.IProducer) desc.OcpPrizeApiServer {
	val := ctx.Value("configuration")
	var conf *configuration.Configuration
	if val != nil {
		conf = val.(*configuration.Configuration)
	}
	if conf != nil {
		flusherMaximumChankSize = conf.FlusherMaximumChankSize
	}
	return &API{
		currentRepo: repo.NewRepo(db),
		producer:    producer,
	}
}

// CreatePrizeV1 creates prize in storage by request
func (a *API) CreatePrizeV1(
	ctx context.Context,
	req *desc.CreatePrizeV1Request,
) (*desc.CreatePrizeV1Response, error) {

	log.Printf("CreatePrizeV1 request: %s", req.String())

	var span opentracing.Span
	if span = opentracing.SpanFromContext(ctx); span != nil {
		span, ctx = opentracing.StartSpanFromContext(ctx, "CreatePrizeV1")
	} else {
		span = opentracing.GlobalTracer().StartSpan("CreatePrizeV1")
	}
	defer span.Finish()

	prizeToAdd := prize.Prize{
		Link:    req.Link,
		IssueID: req.IssueId,
	}
	id, err := a.currentRepo.AddPrizes(ctx, []prize.Prize{prizeToAdd})
	if err != nil {
		log.Printf("CreatePrizeV1 error: %s", err.Error())
		return nil, err
	}

	a.producer.SendMessage("CreatePrizeV1 succesful")

	metrics.IncrementSuccessfulCreate(1)

	response := desc.CreatePrizeV1Response{
		PrizeId: id[0],
	}
	return &response, nil
}

// MultiCreatePrizeV1 creates prizes in storage by request
func (a *API) MultiCreatePrizeV1(
	ctx context.Context,
	req *desc.MultiCreatePrizeV1Request,
) (*desc.MultiCreatePrizeV1Response, error) {
	log.Printf("MultiCreatePrizeV1 request: %s", req.String())

	var span opentracing.Span
	if span = opentracing.SpanFromContext(ctx); span != nil {
		span, ctx = opentracing.StartSpanFromContext(ctx, "MultiCreatePrizeV1")
	} else {
		span = opentracing.GlobalTracer().StartSpan("MultiCreatePrizeV1")
	}
	defer span.Finish()

	prizesToAdd := make([]prize.Prize, 0, len(req.Prizes))

	for _, prizeToAdd := range req.Prizes {
		prizesToAdd = append(prizesToAdd, prize.NewPrize(prizeToAdd.Id, prizeToAdd.IssueId, prizeToAdd.Link))
	}

	flusher := flusher.NewFlusher(a.currentRepo, flusherMaximumChankSize)
	_, ids, err := flusher.Flush(opentracing.ContextWithSpan(ctx, span), prizesToAdd)

	if err != nil {
		log.Printf("MultiCreatePrizeV1 error: %s", err.Error())
		return nil, err
	}

	a.producer.SendMessage("MultiCreatePrizeV1 succesful")

	metrics.IncrementSuccessfulCreate(1)

	response := desc.MultiCreatePrizeV1Response{
		PrizeIds: ids,
	}
	return &response, nil
}

// UpdatePrizeV1 update prize in storage by ID
func (a *API) UpdatePrizeV1(
	ctx context.Context,
	req *desc.UpdatePrizeV1Request,
) (*desc.UpdatePrizeV1Response, error) {

	log.Printf("UpdatePrizeV1 request: %s", req.String())

	var span opentracing.Span
	if span = opentracing.SpanFromContext(ctx); span != nil {
		span, ctx = opentracing.StartSpanFromContext(ctx, "UpdatePrizeV1")
	} else {
		span = opentracing.GlobalTracer().StartSpan("UpdatePrizeV1")
	}
	defer span.Finish()

	prizeToUpdate := prize.Prize{
		ID:      req.Id,
		Link:    req.Link,
		IssueID: req.IssueId,
	}
	err := a.currentRepo.UpdatePrize(ctx, prizeToUpdate)
	if err != nil {
		log.Printf("UpdatePrizeV1 error: %s", err.Error())
		return nil, err
	}

	a.producer.SendMessage("UpdatePrizeV1 succesful")

	metrics.IncrementSuccessfulUpdate(1)

	response := desc.UpdatePrizeV1Response{
		Succeed: true,
	}
	return &response, nil
}

//DescribePrizeV1 gets prize by id from storage by ID
func (a *API) DescribePrizeV1(
	ctx context.Context,
	req *desc.DescribePrizeV1Request,
) (*desc.DescribePrizeV1Response, error) {

	log.Printf("DescribePrizeV1 request: %s", req.String())

	var span opentracing.Span
	if span = opentracing.SpanFromContext(ctx); span != nil {
		span, ctx = opentracing.StartSpanFromContext(ctx, "DescribePrizeV1")
	} else {
		span = opentracing.GlobalTracer().StartSpan("DescribePrizeV1")
	}
	defer span.Finish()

	if err := req.Validate(); err != nil {
		log.Printf("invalid request: %s, error: %s", req.String(), err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	prize, err := a.currentRepo.DescribePrize(ctx, req.PrizeId)
	if err != nil {
		log.Printf("DescribePrizeV1 error: %s", err.Error())
		return nil, err
	}

	a.producer.SendMessage("DescribePrizeV1 succesful")

	metrics.IncrementSuccessfulDescribe(1)

	response := desc.DescribePrizeV1Response{
		Prize: &desc.Prize{
			Id:      prize.ID,
			IssueId: prize.IssueID,
			Link:    prize.Link,
		},
	}
	return &response, nil
}

// ListPrizeV1 gets list of prizes from storage by limit and offset
func (a *API) ListPrizeV1(
	ctx context.Context,
	req *desc.ListPrizeV1Request,
) (*desc.ListPrizeV1Response, error) {

	log.Printf("ListPrizeV1 request: %s", req.String())

	var span opentracing.Span
	if span = opentracing.SpanFromContext(ctx); span != nil {
		span, ctx = opentracing.StartSpanFromContext(ctx, "ListPrizeV1")
	} else {
		span = opentracing.GlobalTracer().StartSpan("ListPrizeV1")
	}
	defer span.Finish()

	prizes, err := a.currentRepo.ListPrizes(ctx, req.Limit, req.Offset)
	if err != nil {
		log.Printf("ListPrizeV1 error: %s", err.Error())
		return nil, err
	}

	a.producer.SendMessage("ListPrizeV1 succesful")

	metrics.IncrementSuccessfulList(1)

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

// RemovePrizeV1 removes prize from storage by id
func (a *API) RemovePrizeV1(
	ctx context.Context,
	req *desc.RemovePrizeV1Request,
) (*desc.RemovePrizeV1Response, error) {

	log.Printf("RemovePrizeV1 request: %s", req.String())

	var span opentracing.Span
	if span = opentracing.SpanFromContext(ctx); span != nil {
		span, ctx = opentracing.StartSpanFromContext(ctx, "RemovePrizeV1")
	} else {
		span = opentracing.GlobalTracer().StartSpan("RemovePrizeV1")
	}
	defer span.Finish()

	if err := req.Validate(); err != nil {
		log.Printf("invalid request: %s, error: %s", req.String(), err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	found, err := a.currentRepo.RemovePrize(ctx, req.PrizeId)

	if err != nil {
		log.Printf("RemovePrizeV1 error: %s", err.Error())
		return nil, err
	}

	a.producer.SendMessage("RemovePrizeV1 succesful")

	metrics.IncrementSuccessfulRemove(1)

	response := desc.RemovePrizeV1Response{
		Found: found,
	}
	return &response, nil
}

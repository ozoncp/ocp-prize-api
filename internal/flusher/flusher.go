package flusher

import (
	"context"
	"errors"
	"fmt"

	"github.com/ozoncp/ocp-prize-api/internal/prize"
	"github.com/ozoncp/ocp-prize-api/internal/repo"

	opentracing "github.com/opentracing/opentracing-go"
)

// IFlusher for prize
type IFlusher interface {
	Flush(ctx context.Context, prize []prize.Prize) ([]prize.Prize, []uint64, error)
}

// Flusher struct
type Flusher struct {
	repo      repo.IRepo
	chunkSize int
}

// NewFlusher create new flusher
func NewFlusher(originRepo repo.IRepo, chSize int) IFlusher {
	return &Flusher{
		repo:      originRepo,
		chunkSize: chSize,
	}
}

// Flush prizes in repo
func (originFlusher *Flusher) Flush(ctx context.Context, prizes []prize.Prize) ([]prize.Prize, []uint64, error) {
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan == nil {
		parentSpan = opentracing.GlobalTracer().StartSpan("MultiCreatePrizeV1")
	}
	chunkSizeToSplit := originFlusher.chunkSize
	if chunkSizeToSplit > len(prizes) {
		chunkSizeToSplit = len(prizes)
	}
	ids := make([]uint64, 0)
	batchedPrizes, err := prize.SplitPrizeSliceToBunches(prizes, chunkSizeToSplit)
	if err != nil {
		return prizes, nil, errors.New(err.Error())
	}
	for i, batchToAdd := range batchedPrizes {
		var childSpan opentracing.Span
		if parentSpan != nil {
			childSpan = opentracing.GlobalTracer().StartSpan("AddPrise", opentracing.ChildOf(parentSpan.Context()))
		}
		partOfIds, err := originFlusher.repo.AddPrizes(ctx, batchToAdd)
		ids = append(ids, partOfIds...)
		if childSpan != nil {
			childSpan.Finish()
		}
		if err != nil {
			return prizes[i*chunkSizeToSplit:], ids, errors.New("error writing prizes from:" + fmt.Sprint(i*chunkSizeToSplit))
		}
	}
	return nil, ids, nil
}

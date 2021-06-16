package flusher

import (
	"context"
	"errors"
	"fmt"

	"github.com/ozoncp/ocp-prize-api/internal/prize"
	"github.com/ozoncp/ocp-prize-api/internal/repo"
)

// IFlusher for prize
type IFlusher interface {
	Flush(ctx context.Context, prize []prize.Prize) ([]prize.Prize, error)
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
func (originFlusher *Flusher) Flush(ctx context.Context, prizes []prize.Prize) ([]prize.Prize, error) {
	chunkSizeToSplit := originFlusher.chunkSize
	if chunkSizeToSplit > len(prizes) {
		chunkSizeToSplit = len(prizes)
	}
	butchedPrizes, err := prize.SplitPrizeSliceToBunches(prizes, chunkSizeToSplit)
	if err != nil {
		return prizes, errors.New(err.Error())
	}
	for i, batchToAdd := range butchedPrizes {
		_, err := originFlusher.repo.AddPrizes(ctx, batchToAdd)
		if err != nil {
			return prizes[i*chunkSizeToSplit:], errors.New("error writing prizes from:" + fmt.Sprint(i*chunkSizeToSplit))
		}
	}
	return nil, nil
}

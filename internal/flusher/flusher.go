package flusher

import (
	"github.com/ozoncp/ocp-prize-api/internal/prize"
	"github.com/ozoncp/ocp-prize-api/internal/repo"
)

type Flusher interface {
	Flush(prize []prize.Prize) []prize.Prize
}

type flusher struct {
	repo      repo.Repo
	chunkSize int
}

func NewFlusher(originRepo repo.Repo, chSize int) Flusher {
	return &flusher{
		repo:      originRepo,
		chunkSize: chSize,
	}
}

func (originFlusher *flusher) Flush(prizes []prize.Prize) []prize.Prize {
	if originFlusher.chunkSize > len(prizes) {
		originFlusher.chunkSize = len(prizes)
	}
	butchedPrizes, err := prize.SplitPrizeSliceToBunches(prizes, originFlusher.chunkSize)
	if err != nil {
		return prizes
	}
	for i, batchToAdd := range butchedPrizes {
		err := originFlusher.repo.AddPrizes(batchToAdd)
		if err != nil {
			return prizes[i*originFlusher.chunkSize:]
		}
	}
	return nil
}

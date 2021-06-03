package repo

import "github.com/ozoncp/ocp-prize-api/internal/prize"

// Repo for prize
type Repo interface {
	AddPrizes(prize []prize.Prize) error
	RemovePrize(prizeID uint64) error
	DescribePrize(prizeID uint64) (*prize.Prize, error)
	ListPrizes(limit, offset uint64) ([]prize.Prize, error)
}

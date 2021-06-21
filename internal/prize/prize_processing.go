package prize

import (
	"errors"
	"log"
)

// SplitPrizeSliceToBunches split slice with bunch size
func SplitPrizeSliceToBunches(prizes []Prize, bunchSize int) (splittedPrizes [][]Prize, err error) {
	log.Print("Split Prize slice to bunches")
	if bunchSize == 0 {
		return splittedPrizes, errors.New("size of required bunch is zero")
	}
	if bunchSize > len(prizes) {
		return splittedPrizes, errors.New("size of required bunch is more than origin slice")
	}
	splittedSize := len(prizes) / bunchSize
	splittedPrizes = make([][]Prize, splittedSize)
	for i := 0; i < splittedSize; i++ {
		splittedPrizes[i] = prizes[i*bunchSize : (i+1)*bunchSize]
	}
	if len(prizes)%bunchSize != 0 {
		splittedPrizes = append(splittedPrizes, prizes[splittedSize*bunchSize:])
	}

	log.Print("Slice splitted successfully")

	return
}

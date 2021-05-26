package prize

import "errors"

func SplitPrizeSliceToBunches(prizes []Prize, bunchSize int) (splittedPrizes [][]Prize, err error) {
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
	return
}

func PrizeSliceToMap(prizeSlice []Prize) (prizeMap map[uint64]Prize, err error) {
	if prizeSlice == nil {
		return prizeMap, errors.New("nil original slice")
	}
	if len(prizeSlice) == 0 {
		return prizeMap, errors.New("empty original slice")
	}
	prizeMap = map[uint64]Prize{}
	for _, prize := range prizeSlice {
		prizeMap[prize.ID] = prize
	}
	return
}

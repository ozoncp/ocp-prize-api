package utils

func SplitSlice(originSlice []uint, size int) (splittedSlice [][]uint) {
	if size == 0 {
		return
	}
	if size >= len(originSlice) {
		splittedSlice = make([][]uint, 1)
		splittedSlice[0] = originSlice
		return
	}
	splittedSize := len(originSlice) / size
	splittedSlice = make([][]uint, splittedSize)
	for i := 0; i < splittedSize; i++ {
		splittedSlice[i] = originSlice[i*size : (i+1)*size]
	}
	if len(originSlice)%size != 0 {
		splittedSlice = append(splittedSlice, originSlice[splittedSize*size:])
	}
	return
}

package utils

// ReverseMap reverse map[key]value to map[value]key
func ReverseMap(originMap map[uint]string) (reversedMap map[string]uint) {
	if originMap == nil {
		return
	}
	reversedMap = map[string]uint{}
	for key, value := range originMap {
		reversedMap[value] = key
	}
	return
}

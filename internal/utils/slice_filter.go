package utils

var originList = [...]int{2, 4, 6, 8, 10, 12}

func FilterSlice(inputSlice []uint) (filteredSlice []uint) {
	isInOriginList := func(valueToSearch uint) bool {
		for _, value := range originList {
			if value == int(valueToSearch) {
				return true
			}
		}
		return false
	}
	for _, inputValue := range inputSlice {
		if !isInOriginList(inputValue) {
			filteredSlice = append(filteredSlice, inputValue)
		}
	}
	return
}

package utils

import (
	"reflect"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	testSlice := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedSlice := []uint{0, 1, 3, 5, 7, 9}
	resultSlice := FilterSlice(testSlice)
	equal := reflect.DeepEqual(resultSlice, expectedSlice)
	if !equal {
		t.Error("Incorrect filter")
	}
}

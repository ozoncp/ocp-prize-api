package utils

import (
	"reflect"
	"testing"
)

func TestSplitMultipleSlice(t *testing.T) {
	testSlice := []uint{0, 1, 2, 3, 4, 5}
	var splittedSize int = 2
	resultSlice := SplitSlice(testSlice, splittedSize)
	if len(resultSlice) != len(testSlice)/splittedSize {
		t.Error("Incorrect count of result slices")
	}
	if len(resultSlice[0]) != splittedSize {
		t.Error("Incorrect splitted size")
	}
	if len(resultSlice[len(resultSlice)-1]) != splittedSize {
		t.Error("Incorrect last slice size")
	}
}

func TestSplitNonMultipleSlice(t *testing.T) {
	testSlice := []uint{0, 1, 2, 3, 4}
	var splittedSize int = 2
	resultSlice := SplitSlice(testSlice, splittedSize)
	if len(resultSlice) != len(testSlice)/splittedSize+1 {
		t.Error("Incorrect count of result slices")
	}
	if len(resultSlice[0]) != splittedSize {
		t.Error("Incorrect splitted size")
	}
	if len(resultSlice[len(resultSlice)-1]) != len(testSlice)%splittedSize {
		t.Error("Incorrect last slice size")
	}
}

func TestSliceWithZeroSize(t *testing.T) {
	testSlice := []uint{0, 1, 2, 3, 4}
	resultSlice := SplitSlice(testSlice, 0)
	if resultSlice != nil {
		t.Error("Incorrect working with zero size")
	}
}
func TestSliceWithSizeEqualOrigin(t *testing.T) {
	testSlice := []uint{0, 1, 2, 3, 4}
	resultSlice := SplitSlice(testSlice, len(testSlice))
	result := reflect.DeepEqual(resultSlice[0], testSlice)
	if !result {
		t.Error("Incorrect splitting with size equal origin size")
	}
}

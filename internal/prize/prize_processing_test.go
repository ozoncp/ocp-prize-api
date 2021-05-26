package prize

import (
	"reflect"
	"testing"
)

func TestSplitMultiplePrizeSlice(t *testing.T) {
	testSlice := []Prize{NewPrize(1, 2, "www"),
		NewPrize(2, 2, "www"), NewPrize(3, 2, "www"),
		NewPrize(4, 2, "www"), NewPrize(5, 2, "www"), NewPrize(6, 2, "www")}
	var splittedSize int = 2
	resultSlice, err := SplitPrizeSliceToBunches(testSlice, splittedSize)
	if err != nil {
		t.Error(err.Error())
	}
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
func TestSplitNonMultiplePrizeSlice(t *testing.T) {
	testSlice := []Prize{NewPrize(1, 2, "www"),
		NewPrize(2, 2, "www"), NewPrize(3, 2, "www"),
		NewPrize(4, 2, "www"), NewPrize(5, 2, "www")}
	var splittedSize int = 2
	resultSlice, err := SplitPrizeSliceToBunches(testSlice, splittedSize)
	if err != nil {
		t.Error(err.Error())
	}
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

func TestPrizeSliceWithZeroSize(t *testing.T) {
	testSlice := []Prize{NewPrize(1, 2, "www"),
		NewPrize(2, 2, "www"), NewPrize(3, 2, "www"),
		NewPrize(4, 2, "www"), NewPrize(5, 2, "www"), NewPrize(6, 2, "www")}
	resultSlice, err := SplitPrizeSliceToBunches(testSlice, 0)
	if resultSlice != nil || err == nil {
		t.Error("Incorrect working with zero size")
	}
}
func TestPrizeSliceWithSizeEqualOrigin(t *testing.T) {
	testSlice := []Prize{NewPrize(1, 2, "www"),
		NewPrize(2, 2, "www"), NewPrize(3, 2, "www"),
		NewPrize(4, 2, "www"), NewPrize(5, 2, "www"), NewPrize(6, 2, "www")}
	resultSlice, err := SplitPrizeSliceToBunches(testSlice, len(testSlice))
	result := reflect.DeepEqual(resultSlice[0], testSlice)
	if !result || err != nil {
		t.Error("Incorrect splitting with size equal origin size")
	}
}

func TestPrizeSliceWithSizeMoreThanOrigin(t *testing.T) {
	testSlice := []Prize{NewPrize(1, 2, "www"),
		NewPrize(2, 2, "www"), NewPrize(3, 2, "www"),
		NewPrize(4, 2, "www"), NewPrize(5, 2, "www"), NewPrize(6, 2, "www")}
	resultSlice, err := SplitPrizeSliceToBunches(testSlice, len(testSlice)+1)
	if resultSlice != nil || err == nil {
		t.Error("Incorrect splitting with size more than origin size")
	}
}

func TestPrizeSliceToMap(t *testing.T) {
	testSlice := []Prize{NewPrize(1, 2, "www"),
		NewPrize(2, 2, "www"), NewPrize(3, 2, "www"),
		NewPrize(4, 2, "www"), NewPrize(5, 2, "www"), NewPrize(6, 2, "www")}
	resultMap, err := PrizeSliceToMap(testSlice)
	if err != nil {
		t.Error(err.Error())
	}
	for _, prize := range testSlice {
		if !reflect.DeepEqual(resultMap[prize.ID], prize) {
			t.Error("Incorrect result map")
		}
	}
}

func TestNilPrizeSliceToMap(t *testing.T) {
	var testSlice []Prize
	resultMap, err := PrizeSliceToMap(testSlice)
	if resultMap != nil || err == nil {
		t.Error("Incorrect working nil slice")
	}
}

func TestEmptyPrizeSliceToMap(t *testing.T) {
	var testSlice = []Prize{}
	resultMap, err := PrizeSliceToMap(testSlice)
	if resultMap != nil || err == nil {
		t.Error("Incorrect working empty slice")
	}
}

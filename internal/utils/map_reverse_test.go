package utils

import "testing"

func TestReverseMap(t *testing.T) {
	testMap := map[uint]string{
		0: "abc",
		1: "def",
		2: "ghi",
		3: "jkl",
	}
	resultMap := ReverseMap(testMap)
	for key, value := range resultMap {
		if testMap[value] != key {
			t.Error("Incorrect reversing")
		}
	}
}

func TestReversenilMap(t *testing.T) {
	var testMap map[uint]string
	resultMap := ReverseMap(testMap)
	if resultMap != nil {
		t.Error("Incorrect reversing")
	}
}

package prize

import "testing"

func TestPrize(t *testing.T) {
	testPrize := NewPrize(1, "sale10")
	if testPrize.String() != "1 sale10 " {
		t.Error("Incorrect cast to string")
	}
	testPrize.SetDescription("123")
	if testPrize.String() != "1 sale10 123" {
		t.Error("Incorrect setting description")
	}
	err := testPrize.SetProbability(90)
	if testPrize.Probability != 90 || err != nil {
		t.Error("Incorrect setting probability from 0 to 100")
	}
	err = testPrize.SetProbability(120)
	if testPrize.Probability != 90 || err == nil {
		t.Error("Incorrect setting probability more than 100")
	}
}

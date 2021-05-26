package prize

import "testing"

func TestPrize(t *testing.T) {
	testPrize := NewPrize(1, 2, "www")
	if testPrize.String() != "ID: 1 IssueID: 2 Link: www" {
		t.Error("Incorrect cast to string")
	}
}

package prize

import (
	"fmt"
)

// Prize entity description
type Prize struct {
	ID      uint64
	IssueID uint64
	Link    string
}

// NewPrize Create new prize
func NewPrize(id uint64, issueID uint64, link string) (newPrize Prize) {
	newPrize.ID = id
	newPrize.IssueID = issueID
	newPrize.Link = link
	return
}

// String get string representation of prize
func (prize *Prize) String() (outString string) {
	outString += "ID:" + " " + fmt.Sprint(prize.ID)
	outString += " " + "IssueID:" + " " + fmt.Sprint(prize.IssueID)
	outString += " " + "Link:" + " " + prize.Link
	return
}

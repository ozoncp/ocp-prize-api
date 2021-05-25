package prize

import (
	"errors"
	"fmt"
)

type Prize struct {
	IDPrize     uint64
	Name        string
	Description string
	Probability uint // In percentage
}

func NewPrize(id uint64, name string) (newPrize Prize) {
	newPrize.IDPrize = id
	newPrize.Name = name
	return
}

func (prize *Prize) SetDescription(description string) {
	prize.Description = description
}

func (prize *Prize) GetDescription() string {
	return prize.Description
}

func (prize *Prize) SetProbability(probability uint) error {
	if probability > 100 {
		return errors.New("incorrect probability: should be from 0 to 100 percent")
	}
	prize.Probability = probability
	return nil
}

func (prize *Prize) String() string {
	return fmt.Sprint(prize.IDPrize) + " " + prize.Name + " " + prize.Description
}

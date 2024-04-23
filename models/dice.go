package models

import "math/rand"

type Dice interface {
	Roll() int
}

// TODO: change max to faces
type diceImpl struct {
	max int
}

func NewDice(max int) Dice {
	return &diceImpl{
		max: max,
	}
}

func (d *diceImpl) Roll() int {
	return rand.Intn(d.max) + 1
}

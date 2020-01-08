package domain

import (
	"math/rand"
	"time"
)

type RandomPlayer struct{}

func (j RandomPlayer) Name() Name {
	return "Random"
}

func (j RandomPlayer) Start() {}

func (j RandomPlayer) Move() Move {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(2)
	if r == 0 {
		return RefuseMove
	}
	return CooperateMove
}

func (j RandomPlayer) Result(Result) {}

func (p RandomPlayer) Clone() IPlayer {
	return p
}

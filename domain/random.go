package domain

import (
	"math/rand"
	"time"
)

type RandomPlayer struct{}

func (p RandomPlayer) Name() Name {
	return "Random"
}

func (p RandomPlayer) Start() {}

func (p RandomPlayer) Move() Move {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(2)
	if r == 0 {
		return RefuseMove
	}
	return CooperateMove
}

func (p RandomPlayer) Result(Result) {}

func (p RandomPlayer) Clone() IPlayer {
	return p
}

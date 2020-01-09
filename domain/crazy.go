package domain

import (
	"math/rand"
	"time"
)

type CrazyPlayer struct{}

func (p CrazyPlayer) Name() Name {
	return "Crazy"
}

func (p CrazyPlayer) Start() {}

func (p CrazyPlayer) Move() Move {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(2)
	if r == 0 {
		return RefuseMove
	}
	return CooperateMove
}

func (p CrazyPlayer) Result(Result) {}

func (p CrazyPlayer) Clone() IPlayer {
	return p
}

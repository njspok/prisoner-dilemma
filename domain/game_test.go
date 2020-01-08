package domain

import "testing"

func TestGame_Play(t *testing.T) {
	p1 := &JudahPlayer{}
	p2 := &TalionPlayer{}

	game := Game{
		P1: p1,
		P2: p2,
	}
	tbl := game.Play(100)
	tbl.Print()
}

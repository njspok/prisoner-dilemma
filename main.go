package main

import "github.com/njspok/prisoner-dilemma/domain"

func main() {
	t := domain.NewTournament([]domain.IPlayer{
		&domain.JesusPlayer{},
		&domain.JudahPlayer{},
		&domain.RandomPlayer{},
		&domain.TalionPlayer{},
	})
	tbl := t.Play()
	tbl.Print()
}

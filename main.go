package main

import "github.com/njspok/prisoner-dilemma/domain"

func main() {
	t := domain.NewTournament([]domain.IPlayer{
		domain.NewJesusPlayer(),
		domain.NewJudahPlayer(),
		domain.NewCrazyPlayer(),
		domain.NewTalionPlayer(),
	})
	tbl := t.Play()
	tbl.Print()
}

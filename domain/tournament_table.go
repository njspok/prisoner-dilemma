package domain

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type TournamentTable struct {
	games []Record
}

func (tt *TournamentTable) Add(r Record) {
	tt.games = append(tt.games, r)
}

func (tt *TournamentTable) Print() {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintln(w, "Player 1", "\t", "Score", "\t", "Player 2", "\t", "Score", " \t", "Total")
	for _, g := range tt.games {
		fmt.Fprintln(w, g.p1Name, "\t", g.p1Score, "\t", g.p2Name, "\t", g.p2Score, " \t", g.Total)
	}

	w.Flush()
}

type Record struct {
	p1Name  Name
	p2Name  Name
	p1Score Score
	p2Score Score
	Total   Score
}

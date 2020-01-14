package domain

import "fmt"

type GameTable struct {
	rounds []RoundRow
	score1 Score
	score2 Score
}

func (t *GameTable) Append(s1, s2 Stat) {
	t.rounds = append(t.rounds, RoundRow{
		Player1: s1,
		Player2: s2,
	})
	t.score1 += s1.Score
	t.score2 += s2.Score
}

func (t *GameTable) Print() {
	for _, row := range t.rounds {
		fmt.Println(row)
	}

	fmt.Println("player 1: ", t.Score1())
	fmt.Println("player 2: ", t.Score2())
	fmt.Println("total: ", t.Total())
}

func (t *GameTable) Total() Score {
	return t.score1 + t.score2
}

func (t *GameTable) Score1() Score {
	return t.score1
}

func (t *GameTable) Score2() Score {
	return t.score2
}

type RoundRow struct {
	Player1 Stat
	Player2 Stat
}

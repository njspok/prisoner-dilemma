package domain

func NewTournament(p []IPlayer) *Tournament {
	if len(p) == 0 {
		panic("minimum 1 players")
	}

	return &Tournament{
		players: p,
	}
}

type Tournament struct {
	players []IPlayer
}

type pair struct {
	p1 IPlayer
	p2 IPlayer
}

func (t *Tournament) Play() TournamentTable {
	// create players pairs
	pairs := makePlayerPairs(t.players)

	tt := TournamentTable{}

	// each pairs make game
	for _, pair := range pairs {
		game := NewGame(pair.p1, pair.p2)
		tbl := game.Play(100)

		// save result game in table
		r := Record{
			Total:   tbl.Total(),
			p1Name:  pair.p1.Name(),
			p1Score: tbl.Score1(),
			p2Name:  pair.p2.Name(),
			p2Score: tbl.Score2(),
		}
		tt.Add(r)
	}

	return tt
}

func makePlayerPairs(players []IPlayer) []pair {
	var pairs []pair
	for _, p1 := range players {
		for _, p2 := range players {
			pairs = append(pairs, pair{
				p1: p1.Clone(),
				p2: p2.Clone(),
			})
		}
	}
	return pairs
}

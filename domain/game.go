package domain

import "fmt"

func NewGame(p1 IPlayer, p2 IPlayer) *Game {
	return &Game{
		P1: p1,
		P2: p2,
	}
}

type Game struct {
	P1 IPlayer
	P2 IPlayer
}

func (g *Game) Play(rounds int) GameTable {
	g.P1.Start()
	g.P2.Start()

	table := GameTable{}

	for i := 1; i <= rounds; i++ {
		move1 := g.P1.Move()
		move2 := g.P2.Move()

		var stat1, stat2 Stat

		switch move1 + move2 {
		case CooperateMove + CooperateMove:
			stat1 = NewStat(move1, 3, ForCooperateReward)
			stat2 = NewStat(move2, 3, ForCooperateReward)
		case RefuseMove + RefuseMove:
			stat1 = NewStat(move1, 1, ForMutualRefuseReward)
			stat2 = NewStat(move2, 1, ForMutualRefuseReward)
		case CooperateMove + RefuseMove:
			stat1 = NewStat(move1, 0, ForNaivetyRewards)
			stat2 = NewStat(move2, 5, ForRiskReward)
		case RefuseMove + CooperateMove:
			stat1 = NewStat(move1, 5, ForRiskReward)
			stat2 = NewStat(move2, 0, ForNaivetyRewards)
		default:
			panic(fmt.Sprintf("unknown move combination %s %s", move1, move2))
		}

		table.Append(stat1, stat2)

		g.P1.Result(Result{
			You: stat1,
			His: stat2,
		})
		g.P2.Result(Result{
			You: stat2,
			His: stat1,
		})
	}

	return table
}

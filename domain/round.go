package domain

import "fmt"

func NewRounder(p1, p2 IPlayer) *Rounder {
	return &Rounder{
		p1: p1,
		p2: p2,
	}
}

type Rounder struct {
	p1 IPlayer
	p2 IPlayer
}

func (r *Rounder) Play() (Stat, Stat) {
	move1 := r.p1.Move()
	move2 := r.p2.Move()

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

	r.p1.Result(Result{
		You: stat1,
		His: stat2,
	})
	r.p2.Result(Result{
		You: stat2,
		His: stat1,
	})

	return stat1, stat2
}

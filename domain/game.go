package domain

func NewGame(p1 IPlayer, p2 IPlayer, r IRound) *Game {
	return &Game{
		p1:      p1,
		p2:      p2,
		rounder: r,
	}
}

type IRound interface {
	Play() (Stat, Stat)
}

type Game struct {
	p1      IPlayer
	p2      IPlayer
	rounder IRound
}

func (g *Game) Play(rounds int) GameTable {
	g.p1.Start()
	g.p2.Start()

	table := GameTable{}

	for i := 1; i <= rounds; i++ {
		stat1, stat2 := g.rounder.Play()
		table.Append(stat1, stat2)
	}

	return table
}

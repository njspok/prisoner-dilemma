package domain

func NewVindictivePlayer() *VindictivePlayer {
	return &VindictivePlayer{
		wasShamed: false,
	}
}

type VindictivePlayer struct {
	wasShamed bool
}

func (p *VindictivePlayer) Name() Name {
	return "Vindictive"
}

func (p *VindictivePlayer) Start() {}

func (p *VindictivePlayer) Move() Move {
	if p.wasShamed {
		return RefuseMove
	}
	return CooperateMove
}

func (p *VindictivePlayer) Result(r Result) {
	if r.His.Move == RefuseMove {
		p.wasShamed = true
	}
}

func (p *VindictivePlayer) Clone() IPlayer {
	return NewVindictivePlayer()
}

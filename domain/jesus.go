package domain

type JesusPlayer struct{}

func (p JesusPlayer) Name() Name {
	return "Jesus"
}

func (p JesusPlayer) Start() {}

func (p JesusPlayer) Move() Move {
	return CooperateMove
}

func (p JesusPlayer) Result(r Result) {}

func (p JesusPlayer) Clone() IPlayer {
	return p
}

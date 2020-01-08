package domain

type JudahPlayer struct{}

func (j JudahPlayer) Name() Name {
	return "Judah"
}

func (j JudahPlayer) Start() {}

func (j JudahPlayer) Move() Move {
	return RefuseMove
}

func (j JudahPlayer) Result(Result) {}

func (p JudahPlayer) Clone() IPlayer {
	return p
}

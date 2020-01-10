package domain

func NewJudahPlayer() *JudahPlayer {
	return &JudahPlayer{}
}

type JudahPlayer struct{}

func (p JudahPlayer) Name() Name {
	return "Judah"
}

func (p JudahPlayer) Start() {}

func (p JudahPlayer) Move() Move {
	return RefuseMove
}

func (p JudahPlayer) Result(Result) {}

func (p JudahPlayer) Clone() IPlayer {
	return NewJudahPlayer()
}

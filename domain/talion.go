package domain

const notMoved Move = ""

func NewTalionPlayer() *TalionPlayer {
	t := &TalionPlayer{}
	t.init()
	return t
}

type TalionPlayer struct {
	lastPartnerMove Move
}

func (p *TalionPlayer) Name() Name {
	return "Talion"
}

func (p *TalionPlayer) Start() {
	p.init()
}

func (p *TalionPlayer) Move() Move {
	if p.lastPartnerMove == notMoved {
		return CooperateMove
	}

	if p.lastPartnerMove == CooperateMove {
		return CooperateMove
	}

	return RefuseMove
}

func (p *TalionPlayer) Result(r Result) {
	p.lastPartnerMove = r.His.Move
}

func (p *TalionPlayer) Clone() IPlayer {
	return NewTalionPlayer()
}

func (p *TalionPlayer) init() {
	p.lastPartnerMove = notMoved
}

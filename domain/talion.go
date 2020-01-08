package domain

type TalionPlayer struct {
	lastPartnerMove Move
}

func (j *TalionPlayer) Name() Name {
	return "Talion"
}

func (j *TalionPlayer) Start() {
	j.lastPartnerMove = ""
}

func (j *TalionPlayer) Move() Move {
	if j.lastPartnerMove == "" {
		return CooperateMove
	}

	if j.lastPartnerMove == CooperateMove {
		return CooperateMove
	}

	return RefuseMove
}

func (j *TalionPlayer) Result(r Result) {
	j.lastPartnerMove = r.His.Move
}

func (p *TalionPlayer) Clone() IPlayer {
	return &TalionPlayer{}
}

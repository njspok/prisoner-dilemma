package domain

type Reward string

const (
	ForRiskReward         Reward = "risk"
	ForCooperateReward    Reward = "cooperate"
	ForMutualRefuseReward Reward = "mutual refuse"
	ForNaivetyRewards     Reward = "naivety"
)

type Score int

type Move string

type Name string

const (
	CooperateMove Move = "cooperate"
	RefuseMove    Move = "refuse"
)

func NewStat(m Move, s Score, r Reward) Stat {
	return Stat{
		Move:   m,
		Score:  s,
		Reward: r,
	}
}

type Stat struct {
	Move   Move
	Score  Score
	Reward Reward
}

type Result struct {
	You Stat
	His Stat
}

type IPlayer interface {
	Name() Name
	Start()
	Move() Move
	Result(Result)
	Clone() IPlayer
}

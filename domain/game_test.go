package domain

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/mock"
)

func TestGame_Play(t *testing.T) {
	// prepare data

	p1 := &testPlayer{}
	p1.On("Start").Return()

	p2 := &testPlayer{}
	p2.On("Start").Return()

	r := &testRounder{
		rounds: []struct {
			s1 Stat
			s2 Stat
		}{
			{NewStat(CooperateMove, 100, ForCooperateReward), NewStat(CooperateMove, 100, ForCooperateReward)},
			{NewStat(RefuseMove, 50, ForMutualRefuseReward), NewStat(RefuseMove, 50, ForMutualRefuseReward)},
			{NewStat(CooperateMove, 0, ForNaivetyRewards), NewStat(RefuseMove, 200, ForRiskReward)},
		},
	}

	// run game

	g := NewGame(p1, p2, r)
	tbl := g.Play(3)

	// check result

	require.Equal(t, GameTable{
		rounds: []RoundRow{
			{
				Player1: NewStat(CooperateMove, 100, ForCooperateReward),
				Player2: NewStat(CooperateMove, 100, ForCooperateReward),
			},
			{
				Player1: NewStat(RefuseMove, 50, ForMutualRefuseReward),
				Player2: NewStat(RefuseMove, 50, ForMutualRefuseReward),
			},
			{
				Player1: NewStat(CooperateMove, 0, ForNaivetyRewards),
				Player2: NewStat(RefuseMove, 200, ForRiskReward),
			},
		},
		score1: 150,
		score2: 350,
	}, tbl)
	require.Equal(t, Score(500), tbl.Total())
	p1.AssertExpectations(t)
	p2.AssertExpectations(t)
}

type testPlayer struct {
	mock.Mock
}

func (p *testPlayer) Name() Name {
	return p.Called().Get(0).(Name)
}

func (p *testPlayer) Start() {
	p.Called()
}

func (p *testPlayer) Move() Move {
	return p.Called().Get(0).(Move)
}

func (p *testPlayer) Result(r Result) {
	p.Called(r)
}

func (p *testPlayer) Clone() IPlayer {
	panic("implement me")
}

type testRounder struct {
	rounds []struct {
		s1 Stat
		s2 Stat
	}
	n int
}

func (t *testRounder) Play() (Stat, Stat) {
	s1 := t.rounds[t.n].s1
	s2 := t.rounds[t.n].s2
	t.n++
	return s1, s2
}

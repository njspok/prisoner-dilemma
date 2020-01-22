package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRounder_Play(t *testing.T) {
	t.Run("Cooperate AND Cooperate", func(t *testing.T) {
		// prepare data
		p1 := &testPlayer{}
		p2 := &testPlayer{}

		p1.On("Move").Return(CooperateMove)
		p1.On("Result", Result{
			You: NewStat(CooperateMove, 3, ForCooperateReward),
			His: NewStat(CooperateMove, 3, ForCooperateReward),
		}).Return()

		p2.On("Move").Return(CooperateMove)
		p2.On("Result", Result{
			You: NewStat(CooperateMove, 3, ForCooperateReward),
			His: NewStat(CooperateMove, 3, ForCooperateReward),
		}).Return()

		// run
		r := NewRounder(p1, p2)
		s1, s2 := r.Play()

		// check result
		require.Equal(t, NewStat(CooperateMove, 3, ForCooperateReward), s1)
		require.Equal(t, NewStat(CooperateMove, 3, ForCooperateReward), s2)
		p1.AssertExpectations(t)
		p2.AssertExpectations(t)
	})
	t.Run("unknown move", func(t *testing.T) {
		p1 := &testPlayer{}
		p2 := &testPlayer{}

		p1.On("Move").Return(Move("shit1"))
		p2.On("Move").Return(Move("shit2"))

		r := NewRounder(p1, p2)

		require.PanicsWithValue(t, "unknown move combination shit1 shit2", func() {
			r.Play()
		})
		p1.AssertExpectations(t)
		p2.AssertExpectations(t)
	})
}

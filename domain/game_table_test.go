package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGameTable(t *testing.T) {
	gt := &GameTable{}
	require.Equal(t, Score(0), gt.Total())
	require.Equal(t, Score(0), gt.Score1())
	require.Equal(t, Score(0), gt.Score2())

	gt.Append(Stat{Score: 99}, Stat{Score: 11})
	require.Equal(t, Score(110), gt.Total())
	require.Equal(t, Score(99), gt.Score1())
	require.Equal(t, Score(11), gt.Score2())

	gt.Append(Stat{Score: 50}, Stat{Score: 19})
	require.Equal(t, Score(179), gt.Total())
	require.Equal(t, Score(149), gt.Score1())
	require.Equal(t, Score(30), gt.Score2())
}

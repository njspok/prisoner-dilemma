package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTalionPlayer_Move(t *testing.T) {
	opponentCooperated := Result{His: Stat{Move: CooperateMove}}
	opponentRefused := Result{His: Stat{Move: RefuseMove}}

	p := NewTalionPlayer()
	p.Start()

	// first step
	require.Equal(t, CooperateMove, p.Move())
	p.Result(opponentCooperated)

	// next
	require.Equal(t, CooperateMove, p.Move())
	p.Result(opponentCooperated)

	// next
	require.Equal(t, CooperateMove, p.Move())
	p.Result(opponentRefused)

	// next
	require.Equal(t, RefuseMove, p.Move())
	p.Result(opponentCooperated)

	// next
	require.Equal(t, CooperateMove, p.Move())
}

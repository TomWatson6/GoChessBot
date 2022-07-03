package rules

import (
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"math"
	"sort"
)

// LastTurnPositionChangeRule checks to see the position change of the last move matches the structs
type LastTurnPositionChangeRule struct {
	Diff move.Position
}

func (r LastTurnPositionChangeRule) Check(p piece.Piece) bool {
	var turns []int

	for k := range p.History {
		turns = append(turns, k)
	}

	sort.Ints(turns)

	last := turns[len(turns)-1]
	pen := turns[len(turns)-2]

	before := p.History[pen]
	after := p.History[last]

	file := math.Abs(float64(after.File - before.File))
	rank := math.Abs(float64(after.Rank - before.Rank))

	return file == float64(r.Diff.File) && rank == float64(r.Diff.Rank)
}

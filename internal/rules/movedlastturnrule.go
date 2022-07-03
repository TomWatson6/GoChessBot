package rules

import (
	"github.com/tomwatson6/chessbot/internal/piece"
	"sort"
)

// MovedLastTurnRule checks to see that the given pawn moved last turn
type MovedLastTurnRule struct{}

func (r MovedLastTurnRule) Check(p piece.Piece) bool {
	var turns []int

	for k := range p.History {
		turns = append(turns, k)
	}

	sort.Ints(turns)

	last := turns[len(turns)-1]
	pen := turns[len(turns)-2]

	return p.History[last] != p.History[pen]
}

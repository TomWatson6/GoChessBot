package rules

import (
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/pkg/utility/set"
)

// FirstMoveRule checks to see that the last move made by a pawn was it's first move
type FirstMoveRule struct{}

func (r FirstMoveRule) Check(p piece.Piece) bool {
	moveSet := set.NewFromMapValues(p.History)

	return len(moveSet) == 2
}

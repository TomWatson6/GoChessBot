package power

import "github.com/tomwatson6/chessbot/internal/move"

// Power is a struct that contains the power of a piece.
// The power of a piece is defined as the number of squares that the piece can move to.
type Power []move.Position

func (p Power) ControlsPosition(pos move.Position) bool {
	for _, p := range p {
		if p == pos {
			return true
		}
	}
	return false
}

func (p Power) Get() int {
	return len(p)
}

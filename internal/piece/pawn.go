package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Pawn struct {
	Colour   colour.Colour
	Position move.Position
	HasMoved bool
}

func (p Pawn) GetLetter() PieceLetter {
	return PieceLetterPawn
}

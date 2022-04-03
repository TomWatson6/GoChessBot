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

func (p Pawn) GetColour() colour.Colour {
	return p.Colour
}

func (p Pawn) GetPiecePoints() PiecePoints {
	return PiecePointsPawn
}

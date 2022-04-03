package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Rook struct {
	Colour   colour.Colour
	Position move.Position
	HasMoved bool
}

func (r Rook) GetLetter() PieceLetter {
	return PieceLetterRook
}

func (r Rook) GetColour() colour.Colour {
	return r.Colour
}

func (r Rook) GetPiecePoints() PiecePoints {
	return PiecePointsRook
}

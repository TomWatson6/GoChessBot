package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type King struct {
	Colour   colour.Colour
	Position move.Position
	HasMoved bool
}

func (k King) GetLetter() PieceLetter {
	return PieceLetterKing
}

func (k King) GetColour() colour.Colour {
	return k.Colour
}

func (k King) GetPiecePoints() PiecePoints {
	return PiecePointsKing
}

func (k King) GetPieceType() PieceType {
	return PieceTypeKing
}

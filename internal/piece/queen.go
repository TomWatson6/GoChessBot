package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Queen struct {
	Colour   colour.Colour
	Position move.Position
	HasMoved bool
}

func (q Queen) GetLetter() PieceLetter {
	return PieceLetterQueen
}

func (q Queen) GetColour() colour.Colour {
	return q.Colour
}

func (q Queen) GetPiecePoints() PiecePoints {
	return PiecePointsQueen
}

func (q Queen) GetPieceType() PieceType {
	return PieceTypeQueen
}

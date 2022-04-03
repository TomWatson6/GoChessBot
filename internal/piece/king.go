package piece

import (
	"math"

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

func (k King) GetPosition() move.Position {
	return k.Position
}

func (k King) GetPiecePoints() PiecePoints {
	return PiecePointsKing
}

func (k King) GetPieceType() PieceType {
	return PieceTypeKing
}

func (k King) IsValidMove(dest move.Position) bool {
	return math.Abs(float64(dest.File-k.Position.File)) <= 1 &&
		math.Abs(float64(dest.Rank-k.Position.Rank)) <= 1
}

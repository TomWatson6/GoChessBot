package piece

import (
	"math"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Knight struct {
	Colour   colour.Colour
	Position move.Position
	HasMoved bool
}

func (k Knight) GetLetter() PieceLetter {
	return PieceLetterKnight
}

func (k Knight) GetColour() colour.Colour {
	return k.Colour
}

func (k Knight) GetPiecePoints() PiecePoints {
	return PiecePointsKnight
}

func (k Knight) GetPieceType() PieceType {
	return PieceTypeKnight
}

func (k Knight) IsValidMove(dest move.Position) bool {
	if math.Abs(float64(dest.File-k.Position.File)) == 2 {
		return math.Abs(float64(dest.Rank-k.Position.Rank)) == 1
	} else if math.Abs(float64(dest.Rank-k.Position.Rank)) == 2 {
		return math.Abs(float64(dest.File-k.Position.File)) == 1
	}

	return false
}

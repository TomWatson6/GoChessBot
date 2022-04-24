package piece

import (
	"math"

	"github.com/tomwatson6/chessbot/internal/move"
)

type King struct{}

func (k King) GetPieceLetter() PieceLetter {
	return PieceLetterKing
}

func (k King) GetPiecePoints() PiecePoints {
	return PiecePointsKing
}

func (k King) GetPieceType() PieceType {
	return PieceTypeKing
}

func (k King) IsValidMove(m move.Move) bool {
	return math.Abs(float64(m.To.File-m.From.File)) <= 1 &&
		math.Abs(float64(m.To.Rank-m.From.Rank)) <= 1
}

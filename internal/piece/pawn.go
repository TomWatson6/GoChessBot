package piece

import (
	"math"

	"github.com/tomwatson6/chessbot/internal/move"
)

type Pawn struct {
	HasMoved bool
}

func (p Pawn) GetPieceLetter() PieceLetter {
	return PieceLetterPawn
}

func (p Pawn) GetPiecePoints() PiecePoints {
	return PiecePointsPawn
}

func (p Pawn) GetPieceType() PieceType {
	return PieceTypePawn
}

func (p Pawn) IsValidMove(m move.Move) bool {
	y := math.Abs(float64(m.To.Rank - m.From.Rank))
	x := math.Abs(float64(m.To.File - m.From.File))

	if y == 1 && x <= 1 {
		return true
	} else if y == 2 && !p.HasMoved {
		return x == 0
	}

	return false
}

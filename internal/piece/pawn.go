package piece

import (
	"math"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Pawn struct {
	Colour   colour.Colour
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
	y := m.To.Rank - m.From.Rank
	x := math.Abs(float64(m.To.File - m.From.File))

	mult := 1

	// Pawns can only move in 1 direction
	if p.Colour == colour.Black {
		mult = -1
	}

	if y == mult && x <= 1 {
		return true
	} else if y == mult*2 && !p.HasMoved {
		return x == 0
	}

	return false
}

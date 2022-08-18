package piece

import (
	"github.com/tomwatson6/chessbot/internal/piece/rules"
	"math"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Pawn struct {
	maxRange int
	Colour   colour.Colour
	HasMoved bool
}

func NewPawn(c colour.Colour) *Pawn {
	return &Pawn{
		maxRange: 1,
		Colour:   c,
	}
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

func (p *Pawn) Move(m move.Move) error {
	r := p.maxRange

	if !p.HasMoved {
		r += 1
	}

	rs := rules.Assert(
		rules.IsValidLine(m),
		rules.DoesNotExceedMaxRange(r, m),
		rules.IsCorrectDirection(p.Colour, m),
		rules.DoesNotExceedMaxRangeIfDiagonal(p.maxRange, m),
	)

	if err := rs(); err != nil {
		return err
	}

	return nil
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

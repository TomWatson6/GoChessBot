package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece/rules"
)

type Pawn struct {
	minRange, maxRange int
	colour             colour.Colour
	hasMoved           bool
}

type PawnOption func(p *Pawn)

func NewPawn(opts ...PawnOption) Pawn {
	p := Pawn{
		minRange: 1,
		maxRange: 1,
		colour:   colour.White,
	}

	for _, opt := range opts {
		opt(&p)
	}

	return p
}

func PawnWithColour(c colour.Colour) PawnOption {
	return func(p *Pawn) {
		p.colour = c
	}
}

func PawnWithHasMoved(moved bool) PawnOption {
	return func(p *Pawn) {
		p.hasMoved = moved
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

func (p Pawn) IsValidMove(m move.Move) error {
	r := p.maxRange

	if !p.hasMoved {
		r += 1
	}

	rs := rules.Assert(
		rules.IsValidLine(m),
		rules.IsLargerThanOrEqualToMinRange(p.minRange, m),
		rules.DoesNotExceedMaxRange(r, m),
		rules.IsCorrectDirection(p.colour, m),
		rules.DoesNotExceedMaxRangeIfDiagonal(p.maxRange, m),
	)

	if err := rs(); err != nil {
		return err
	}

	return nil
}

func (p Pawn) HasMoved() bool {
	return p.hasMoved
}

//func (p Pawn) IsValidMove(m move.Move) bool {
//	y := m.To.Rank - m.From.Rank
//	x := math.Abs(float64(m.To.File - m.From.File))
//
//	mult := 1
//
//	// Pawns can only move in 1 direction
//	if p.Colour == colour.Black {
//		mult = -1
//	}
//
//	if y == mult && x <= 1 {
//		return true
//	} else if y == mult*2 && !p.HasMoved {
//		return x == 0
//	}
//
//	return false
//}

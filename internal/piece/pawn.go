package piece

import (
	"math"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Pawn struct {
	Colour     colour.Colour
	Position   move.Position
	HasMoved   bool
	ValidMoves []move.Position
}

func (p Pawn) GetLetter() PieceLetter {
	return PieceLetterPawn
}

func (p Pawn) GetColour() colour.Colour {
	return p.Colour
}

func (p Pawn) GetPosition() move.Position {
	return p.Position
}

func (p Pawn) SetPosition(pos move.Position) Piece {
	p.Position = pos
	return p
}

func (p Pawn) GetPiecePoints() PiecePoints {
	return PiecePointsPawn
}

func (p Pawn) GetPieceType() PieceType {
	return PieceTypePawn
}

func (p Pawn) IsValidMove(dest move.Position) bool {
	y := dest.Rank - p.Position.Rank
	x := math.Abs(float64(dest.File - p.Position.File))

	if y == 1 && x <= 1 {
		return true
	} else if y == 2 && !p.HasMoved {
		return x == 0
	}

	return false
}

func (p Pawn) AppendValidMove(dest move.Position) Piece {
	p.ValidMoves = append(p.ValidMoves, dest)
	return p
}

func (p Pawn) ResetValidMoves() Piece {
	p.ValidMoves = []move.Position{}
	return p
}

func (p Pawn) GetValidMoves() []move.Position {
	return p.ValidMoves
}

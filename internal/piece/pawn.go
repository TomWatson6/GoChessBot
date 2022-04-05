package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Pawn struct {
	Colour   colour.Colour
	Position move.Position
	HasMoved bool
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

	return y == 1 || (!p.HasMoved && y == 2)
}

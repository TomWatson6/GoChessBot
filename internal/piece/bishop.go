package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Bishop struct {
	Colour   colour.Colour
	Position move.Position
	HasMoved bool
}

func (b Bishop) GetLetter() PieceLetter {
	return PieceLetterBishop
}

func (b Bishop) GetColour() colour.Colour {
	return b.Colour
}

func (b Bishop) GetPosition() move.Position {
	return b.Position
}

func (b Bishop) SetPosition(pos move.Position) Piece {
	b.Position = pos
	return b
}

func (b Bishop) GetPiecePoints() PiecePoints {
	return PiecePointsBishop
}

func (b Bishop) GetPieceType() PieceType {
	return PieceTypeBishop
}

func (b Bishop) IsValidMove(dest move.Position) bool {
	x := dest.File - b.Position.File
	y := dest.Rank - b.Position.Rank
	return (x == y || x == -y) && (x != 0 && y != 0)
}

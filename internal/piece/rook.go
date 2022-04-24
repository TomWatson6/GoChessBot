package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Rook struct {
	Colour     colour.Colour
	Position   move.Position
	HasMoved   bool
	ValidMoves []move.Position
}

func (r Rook) GetLetter() PieceLetter {
	return PieceLetterRook
}

func (r Rook) GetColour() colour.Colour {
	return r.Colour
}

func (r Rook) GetPosition() move.Position {
	return r.Position
}

func (r Rook) SetPosition(pos move.Position) Piece {
	r.Position = pos
	return r
}

func (r Rook) GetPiecePoints() PiecePoints {
	return PiecePointsRook
}

func (r Rook) GetPieceType() PieceType {
	return PieceTypeRook
}

func (r Rook) IsValidMove(dest move.Position) bool {
	return r.Position.File == dest.File || r.Position.Rank == dest.Rank
}

func (r Rook) AppendValidMove(dest move.Position) Piece {
	r.ValidMoves = append(r.ValidMoves, dest)
	return r
}

func (r Rook) ResetValidMoves() Piece {
	r.ValidMoves = []move.Position{}
	return r
}

func (r Rook) GetValidMoves() []move.Position {
	return r.ValidMoves
}

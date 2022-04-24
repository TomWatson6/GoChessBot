package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Queen struct {
	Colour     colour.Colour
	Position   move.Position
	HasMoved   bool
	ValidMoves []move.Position
}

func (q Queen) GetLetter() PieceLetter {
	return PieceLetterQueen
}

func (q Queen) GetColour() colour.Colour {
	return q.Colour
}

func (q Queen) GetPosition() move.Position {
	return q.Position
}

func (q Queen) SetPosition(pos move.Position) Piece {
	q.Position = pos
	return q
}

func (q Queen) GetPiecePoints() PiecePoints {
	return PiecePointsQueen
}

func (q Queen) GetPieceType() PieceType {
	return PieceTypeQueen
}

func (q Queen) IsValidMove(dest move.Position) bool {
	x := dest.File - q.Position.File
	y := dest.Rank - q.Position.Rank

	// Horizontal and Vertical moves
	if (x == 0 && y != 0) || (y == 0 && x != 0) {
		return true
	}

	// Diagonal moves
	if (x == y || x == -y) && (x != 0 && y != 0) {
		return true
	}

	return false
}

func (q Queen) AppendValidMove(dest move.Position) Piece {
	q.ValidMoves = append(q.ValidMoves, dest)
	return q
}

func (q Queen) ResetValidMoves() Piece {
	q.ValidMoves = []move.Position{}
	return q
}

func (q Queen) GetValidMoves() []move.Position {
	return q.ValidMoves
}

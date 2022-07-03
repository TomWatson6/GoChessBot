package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type PieceDetails interface {
	GetPieceType() PieceType
	GetPiecePoints() PiecePoints
	GetPieceLetter() PieceLetter
	IsValidMove(m move.Move) bool
}

type Piece struct {
	Position   move.Position
	Colour     colour.Colour
	ValidMoves map[move.Position]bool
	History    map[int]move.Position
	PieceDetails
}

func (p Piece) Equals(p2 Piece) bool {
	if p.GetPieceType() != p2.GetPieceType() {
		return false
	}
	if p.Position != p2.Position {
		return false
	}
	if p.Colour != p2.Colour {
		return false
	}
	return true
}

type PieceType byte

const (
	PieceTypePawn PieceType = iota
	PieceTypeKnight
	PieceTypeBishop
	PieceTypeRook
	PieceTypeQueen
	PieceTypeKing
)

type PiecePoints int

const (
	PiecePointsPawn   PiecePoints = 1
	PiecePointsKnight PiecePoints = 3
	PiecePointsBishop PiecePoints = 3
	PiecePointsRook   PiecePoints = 5
	PiecePointsQueen  PiecePoints = 9
	PiecePointsKing   PiecePoints = 100
)

type PieceLetter rune

const (
	PieceLetterPawn   PieceLetter = 'P'
	PieceLetterKnight PieceLetter = 'N'
	PieceLetterBishop PieceLetter = 'B'
	PieceLetterRook   PieceLetter = 'R'
	PieceLetterQueen  PieceLetter = 'Q'
	PieceLetterKing   PieceLetter = 'K'
)

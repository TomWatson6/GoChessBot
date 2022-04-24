package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

// Needs to be simplified at some point, getting way too big!
type Piece interface {
	GetLetter() PieceLetter
	GetColour() colour.Colour
	GetPosition() move.Position
	SetPosition(pos move.Position) Piece
	GetPiecePoints() PiecePoints
	GetPieceType() PieceType
	IsValidMove(dest move.Position) bool
	AppendValidMove(dest move.Position) Piece
	ResetValidMoves() Piece
	GetValidMoves() []move.Position
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

package piece

import (
	"github.com/tomwatson6/chessbot/internal/ai/power"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Piece interface {
	GetLetter() PieceLetter
	GetColour() colour.Colour
	GetPower() power.Power
	GetThreatLevel() int
	GetPiecePoints() PiecePoints
	GetPieceType() PieceType
	IsValidMove(m move.Move) bool
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

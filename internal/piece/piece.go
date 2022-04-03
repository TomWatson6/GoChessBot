package piece

import (
	"github.com/tomwatson6/chessbot/internal/ai/power"
	"github.com/tomwatson6/chessbot/internal/colour"
)

type Piece interface {
	GetLetter() rune
	GetColour() colour.Colour
	GetPower() power.Power
	GetThreatLevel() int
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

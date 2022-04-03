package piece

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Knight struct {
	Colour   colour.Colour
	Position move.Position
	HasMoved bool
}

func (k Knight) GetLetter() PieceLetter {
	return PieceLetterKnight
}

func (k Knight) GetColour() colour.Colour {
	return k.Colour
}

func (k Knight) GetPiecePoints() PiecePoints {
	return PiecePointsKnight
}

func (k Knight) GetPieceType() PieceType {
	return PieceTypeKnight
}

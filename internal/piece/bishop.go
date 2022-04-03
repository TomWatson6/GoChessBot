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

func (b Bishop) GetPiecePoints() PiecePoints {
	return PiecePointsBishop
}

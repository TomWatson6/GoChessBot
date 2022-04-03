package piece

import "github.com/tomwatson6/chessbot/internal/colour"

type Piece interface {
	GetLetter() rune
	GetColour() colour.Colour
}

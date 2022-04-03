package piece

import "github.com/tomwatson6/chessbot/colour"

type Piece interface {
	GetLetter() rune
	GetColour() colour.Colour
}

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

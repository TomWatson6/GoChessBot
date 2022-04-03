package chess

import (
	"github.com/tomwatson6/chessbot/colour"
	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Chess struct {
	Board board.Board
	Turn  colour.Colour
}

func (c *Chess) MakeMove(m move.Move) error {
	return nil
}

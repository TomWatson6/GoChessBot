package chess

import (
	"github.com/tomwatson6/chessbot/cmd/board"
	"github.com/tomwatson6/chessbot/cmd/move"
	"github.com/tomwatson6/chessbot/colour"
)

type Chess struct {
	Board board.Board
	Turn  colour.Colour
}

func (c *Chess) MakeMove(m move.Move) error {
	return nil
}

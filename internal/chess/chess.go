package chess

import (
	"fmt"

	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

type Chess struct {
	Board board.Board
	Turn  colour.Colour
}

func (c *Chess) MakeMove(m move.Move) error {
	if err := c.Board.MovePiece(m); err != nil {
		return fmt.Errorf("invalid move: %v", m)
	} else {
		return nil
	}
}

func (c *Chess) NextTurn() {
	if c.Turn == colour.White {
		c.Turn = colour.Black
	} else {
		c.Turn = colour.White
	}
}

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
	if c.Board.Pieces[m.From].GetColour() != c.Turn {
		return fmt.Errorf("invalid move for current turn: %v", m)
	}
	if err := c.Board.MovePiece(m); err != nil {
		return fmt.Errorf("invalid move: %v", m)
	}

	c.nextTurn()
	return nil
}

func (c *Chess) nextTurn() {
	if c.Turn == colour.White {
		c.Turn = colour.Black
	} else {
		c.Turn = colour.White
	}
}

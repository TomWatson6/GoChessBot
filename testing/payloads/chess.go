package payloads

import (
	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/chess"
	"github.com/tomwatson6/chessbot/internal/colour"
)

type ChessOption func(c *chess.Chess) error

func NewStandardChessGame(opts ...ChessOption) chess.Chess {
	c := chess.New(colour.White)

	for _, opt := range opts {
		if err := opt(&c); err != nil {
			panic(err)
		}
	}

	return c
}

func ChessGameWithTurn(turn colour.Colour) ChessOption {
	return func(c *chess.Chess) error {
		c.Turn = turn
		return nil
	}
}

func ChessGameWithBoard(b board.Board) ChessOption {
	return func(c *chess.Chess) error {
		c.Board = b
		return nil
	}
}

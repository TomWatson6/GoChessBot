package chess

import (
	"bufio"
	"fmt"
	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"os"
	"strings"
)

type Chess struct {
	Board board.Board
	Turn  colour.Colour
}

func New(col colour.Colour) Chess {
	var c Chess

	c.Board = board.New(8, 8)
	c.Turn = col

	return c
}

// Play starts a game for the colour specified, and cycles turns until a winner is determined - returns winning colour
//func (c *Chess) Play(col colour.Colour) colour.Colour {
//Outer:
//	for {
//		if c.Board.IsCheckMate(c.Turn) {
//			return c.Turn.Opposite()
//		}
//
//		input, err := getInput(c.Turn)
//		if err != nil {
//			continue
//		}
//
//		ms, err := c.TranslateNotation(input)
//
//		for _, m := range ms {
//			if err := c.MakeMove(m); err != nil {
//				continue Outer
//			}
//		}
//
//		c.NextTurn()
//	}
//}

func (c *Chess) MakeMove(m move.Move) error {
	if c.Board.Pieces[m.From].Colour != c.Turn {
		return fmt.Errorf("invalid move for current turn: %v", m)
	}
	if err := c.Board.Move(m); err != nil {
		return fmt.Errorf("failed validation of move: %v, err: %w", m, err)
	}

	return nil
}

func (c *Chess) NextTurn() {
	if c.Turn == colour.White {
		c.Turn = colour.Black
	} else {
		c.Turn = colour.White
	}
}

func getInput(c colour.Colour) (string, error) {
	fmt.Printf("%s's move: ", c)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	// remove CRLF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text, nil
}

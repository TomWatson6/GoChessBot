package chess

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/tomwatson6/chessbot/cmd/config"
	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

var (
	// ErrorPieceNotInStartPosition is thrown when there is no piece in the start position provided
	ErrorPieceNotInStartPosition = errors.New("there is no piece in the start position provided")
)

type Chess struct {
	Board board.Board   `json:"board"`
	Turn  colour.Colour `json:"turn"`
}

func New(col colour.Colour) Chess {
	var c Chess

	width, height := config.GetBoardDimensions()
	c.Board = board.New(width, height)
	c.Turn = col

	return c
}

func (c Chess) MarshalJSON() ([]byte, error) {
	b, err := c.Board.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bMap := map[string]any{}
	err = json.Unmarshal(b, &bMap)
	if err != nil {
		return nil, err
	}

	return json.Marshal(
		struct {
			Board map[string]any `json:"board"`
			Turn  string         `json:"turn"`
		}{
			bMap,
			c.Turn.String(),
		},
	)
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

func (c *Chess) MakeMove(m move.Move) ([]move.Move, error) {
	if _, ok := c.Board.Pieces[m.From]; !ok {
		return []move.Move{}, ErrorPieceNotInStartPosition
	}

	if c.Board.Pieces[m.From].Colour != c.Turn {
		return []move.Move{}, fmt.Errorf("invalid move for current turn: %v", m)
	}

	moves, err := c.Board.Move(m)
	if err != nil {
		return []move.Move{}, fmt.Errorf("failed validation of move: %v, err: %w", m, err)
	}

	c.NextTurn()

	return moves, nil
}

func (c *Chess) NextTurn() {
	c.Turn = c.Turn.Opposite()
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

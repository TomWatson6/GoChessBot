package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tomwatson6/chessbot/internal/chess"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/pkg/output"
)

func getUserInput(c colour.Colour) (string, error) {
	fmt.Printf("%s's move: ", c)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text, nil
}

// Known bugs:
// - pawn promotion to queen
// - en passant
// -

func main() {
	c := chess.New(colour.White)

	for {
		output.PrintBoard(c.Board, c.Turn)
		input, err := getUserInput(c.Turn)
		if err != nil {
			fmt.Println(err)
			continue
		}

		ms, err := c.TranslateNotation(input)

		fmt.Printf("%v\n", ms)

		if err != nil {
			fmt.Println(err)
			continue
		}

		hasError := false

		for _, m := range ms {
			if err := c.MakeMove(m); err != nil {
				fmt.Println(err)
				hasError = true
			}
		}

		if hasError {
			continue
		}

		c.NextTurn()
	}
}

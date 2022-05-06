package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/pkg/input"
	"github.com/tomwatson6/chessbot/pkg/output"
)

func main() {
	c := input.Get(colour.White,
		make(map[move.Position]piece.Piece),
		[]move.Position{},
	)
	reader := bufio.NewReader(os.Stdin)

	for {
		output.PrintBoard(c.Board, c.Turn)
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		ms, err := c.TranslateNotation(text)
		fmt.Printf("%v\n", ms)
		if err == nil {
			for _, m := range ms {
				if err := c.MakeMove(m); err != nil {
					fmt.Println(err)
				}
			}
			c.NextTurn()
		} else {
			fmt.Println(err)
			break
		}
	}
}

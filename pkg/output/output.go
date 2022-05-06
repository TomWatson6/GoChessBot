package output

import (
	"fmt"

	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func getPieceDisplay(p piece.Piece) string {
	output := ""
	if p.Colour == colour.White {
		output += "w"
	} else {
		output += "b"
	}

	output += string(p.GetPieceLetter())

	return output
}

// Need to fix board printing!!
func PrintBoard(b board.Board, c colour.Colour) {
	lower := 0
	upper := 8
	step := 1

	if c == colour.White {
		lower = 7
		upper = -1
		step = -1
	}

	for r := lower; r != upper; r += step {
		fmt.Printf("%d ", r+1)

		for f := upper - 1; f != lower-1; f -= step {
			if p, ok := b.Pieces[move.Position{File: f, Rank: r}]; ok {
				fmt.Printf("%s ", getPieceDisplay(p))
			} else {
				fmt.Printf("## ")
			}
		}
		fmt.Println()
	}

	if c == colour.White {
		fmt.Printf("   A  B  C  D  E  F  G  H\n")
	} else {
		fmt.Printf("   H  G  F  E  D  C  B  A\n")
	}
}

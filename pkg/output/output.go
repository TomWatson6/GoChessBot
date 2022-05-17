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
	if c == colour.White {
		for r := 7; r >= 0; r-- {
			fmt.Printf("%d ", r+1)

			for f := 0; f < 8; f++ {
				if p, ok := b.Pieces[move.Position{File: f, Rank: r}]; ok {
					fmt.Printf("%s ", getPieceDisplay(p))
				} else {
					fmt.Printf("## ")
				}
			}

			fmt.Println()
		}

		fmt.Printf("   A  B  C  D  E  F  G  H\n")
	} else {
		for r := 0; r < 8; r++ {
			fmt.Printf("%d ", r+1)

			for f := 7; f >= 0; f-- {
				if p, ok := b.Pieces[move.Position{File: f, Rank: r}]; ok {
					fmt.Printf("%s ", getPieceDisplay(p))
				} else {
					fmt.Printf("## ")
				}
			}

			fmt.Println()
		}

		fmt.Printf("   H  G  F  E  D  C  B  A\n")
	}
}

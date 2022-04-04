package output

import (
	"fmt"

	"github.com/tomwatson6/chessbot/internal/chess"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func GetPieceDisplay(p piece.Piece) string {
	output := ""
	if p.GetColour() == colour.White {
		output += "w"
	} else {
		output += "b"
	}

	output += string(p.GetLetter())

	return output
}

func PrintBoard(c chess.Chess) {
	lower := 0
	upper := 8
	step := 1

	if c.Turn == colour.White {
		lower = 7
		upper = -1
		step = -1
	}

	for rank := lower; rank != upper; rank += step {
		fmt.Printf("%d ", rank+1)

		for file := 0; file < 8; file++ {
			if piece := c.Board.GetPiece(move.Position{File: file, Rank: rank}); piece != nil {
				fmt.Printf("%s ", GetPieceDisplay(piece))
			} else {
				fmt.Printf("## ")
			}
		}
		fmt.Println()
	}

	if c.Turn == colour.White {
		fmt.Printf("   A  B  C  D  E  F  G  H\n")
	} else {
		fmt.Printf("   H  G  F  E  D  C  B  A\n")
	}
}

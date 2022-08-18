package generation

import (
	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

func GetValidMoves(b board.Board, c colour.Colour) []move.Move {
	var moves []move.Move

	for _, piece := range b.Pieces {
		if piece.Colour != c {
			continue
		}

		for p := range piece.ValidMoves {
			moves = append(moves, move.Move{From: piece.Position, To: p})
		}
	}

	return moves
}

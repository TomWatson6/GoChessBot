package input

import (
	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/chess"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func Get(c colour.Colour, alterations map[move.Position]piece.Piece) chess.Chess {
	// Make board for standard game of chess
	pieces := make(map[move.Position]piece.Piece)

	piecesList := []piece.Piece{
		&piece.Rook{Colour: colour.White, Position: move.Position{File: 0, Rank: 0}},
		&piece.Knight{Colour: colour.White, Position: move.Position{File: 1, Rank: 0}},
		&piece.Bishop{Colour: colour.White, Position: move.Position{File: 2, Rank: 0}},
		&piece.Queen{Colour: colour.White, Position: move.Position{File: 3, Rank: 0}},
		&piece.King{Colour: colour.White, Position: move.Position{File: 4, Rank: 0}},
		&piece.Bishop{Colour: colour.White, Position: move.Position{File: 5, Rank: 0}},
		&piece.Knight{Colour: colour.White, Position: move.Position{File: 6, Rank: 0}},
		&piece.Rook{Colour: colour.White, Position: move.Position{File: 7, Rank: 0}},
		&piece.Pawn{Colour: colour.White, Position: move.Position{File: 0, Rank: 1}},
		&piece.Pawn{Colour: colour.White, Position: move.Position{File: 1, Rank: 1}},
		&piece.Pawn{Colour: colour.White, Position: move.Position{File: 2, Rank: 1}},
		&piece.Pawn{Colour: colour.White, Position: move.Position{File: 3, Rank: 1}},
		&piece.Pawn{Colour: colour.White, Position: move.Position{File: 4, Rank: 1}},
		&piece.Pawn{Colour: colour.White, Position: move.Position{File: 5, Rank: 1}},
		&piece.Pawn{Colour: colour.White, Position: move.Position{File: 6, Rank: 1}},
		&piece.Pawn{Colour: colour.White, Position: move.Position{File: 7, Rank: 1}},
		&piece.Rook{Colour: colour.Black, Position: move.Position{File: 0, Rank: 7}},
		&piece.Knight{Colour: colour.Black, Position: move.Position{File: 1, Rank: 7}},
		&piece.Bishop{Colour: colour.Black, Position: move.Position{File: 2, Rank: 7}},
		&piece.Queen{Colour: colour.Black, Position: move.Position{File: 3, Rank: 7}},
		&piece.King{Colour: colour.Black, Position: move.Position{File: 4, Rank: 7}},
		&piece.Bishop{Colour: colour.Black, Position: move.Position{File: 5, Rank: 7}},
		&piece.Knight{Colour: colour.Black, Position: move.Position{File: 6, Rank: 7}},
		&piece.Rook{Colour: colour.Black, Position: move.Position{File: 7, Rank: 7}},
		&piece.Pawn{Colour: colour.Black, Position: move.Position{File: 0, Rank: 6}},
		&piece.Pawn{Colour: colour.Black, Position: move.Position{File: 1, Rank: 6}},
		&piece.Pawn{Colour: colour.Black, Position: move.Position{File: 2, Rank: 6}},
		&piece.Pawn{Colour: colour.Black, Position: move.Position{File: 3, Rank: 6}},
		&piece.Pawn{Colour: colour.Black, Position: move.Position{File: 4, Rank: 6}},
		&piece.Pawn{Colour: colour.Black, Position: move.Position{File: 5, Rank: 6}},
		&piece.Pawn{Colour: colour.Black, Position: move.Position{File: 6, Rank: 6}},
		&piece.Pawn{Colour: colour.Black, Position: move.Position{File: 7, Rank: 6}},
	}

	for _, p := range piecesList {
		pieces[p.GetPosition()] = p
	}

	for k, v := range alterations {
		pieces[k] = v
	}

	chess := chess.Chess{
		Board: board.Board{
			Pieces: pieces,
		},
		Turn: c,
	}

	return chess
}

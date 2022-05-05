package input

import (
	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/chess"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func Get(c colour.Colour, alterations map[move.Position]piece.Piece, deletions []move.Position) chess.Chess {
	// Make board for standard game of chess

	var squares []move.Position

	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			squares = append(squares, move.Position{File: f, Rank: r})
		}
	}

	pieces := make(map[move.Position]piece.Piece)

	piecesList := []piece.Piece{
		{Colour: colour.White, Position: move.Position{File: 0, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Rook{}},
		{Colour: colour.White, Position: move.Position{File: 1, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Knight{}},
		{Colour: colour.White, Position: move.Position{File: 2, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Bishop{}},
		{Colour: colour.White, Position: move.Position{File: 3, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Queen{}},
		{Colour: colour.White, Position: move.Position{File: 4, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.King{}},
		{Colour: colour.White, Position: move.Position{File: 5, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Bishop{}},
		{Colour: colour.White, Position: move.Position{File: 6, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Knight{}},
		{Colour: colour.White, Position: move.Position{File: 7, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Rook{}},
		{Colour: colour.White, Position: move.Position{File: 0, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.White, Position: move.Position{File: 1, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.White, Position: move.Position{File: 2, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.White, Position: move.Position{File: 3, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.White, Position: move.Position{File: 4, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.White, Position: move.Position{File: 5, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.White, Position: move.Position{File: 6, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.White, Position: move.Position{File: 7, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.Black, Position: move.Position{File: 0, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Rook{}},
		{Colour: colour.Black, Position: move.Position{File: 1, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Knight{}},
		{Colour: colour.Black, Position: move.Position{File: 2, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Bishop{}},
		{Colour: colour.Black, Position: move.Position{File: 3, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Queen{}},
		{Colour: colour.Black, Position: move.Position{File: 4, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.King{}},
		{Colour: colour.Black, Position: move.Position{File: 5, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Bishop{}},
		{Colour: colour.Black, Position: move.Position{File: 6, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Knight{}},
		{Colour: colour.Black, Position: move.Position{File: 7, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Rook{}},
		{Colour: colour.Black, Position: move.Position{File: 0, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.Black, Position: move.Position{File: 1, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.Black, Position: move.Position{File: 2, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.Black, Position: move.Position{File: 3, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.Black, Position: move.Position{File: 4, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.Black, Position: move.Position{File: 5, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.Black, Position: move.Position{File: 6, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
		{Colour: colour.Black, Position: move.Position{File: 7, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{}},
	}

	for _, p := range piecesList {
		pieces[p.Position] = p
	}

	for k, v := range alterations {
		pieces[k] = v
	}

	for _, d := range deletions {
		delete(pieces, d)
	}

	chess := chess.Chess{
		Board: board.Board{
			Squares: squares,
			Pieces:  pieces,
		},
		Turn: c,
	}

	chess.Board.GenerateMoveMap()
	chess.Board.GenerateThreatMap()

	return chess
}

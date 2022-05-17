package config

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func GetStandardPieces() []piece.Piece {
	standardPieces := []piece.Piece{
		{Colour: colour.White, Position: move.Position{File: 0, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Rook{}},
		{Colour: colour.White, Position: move.Position{File: 1, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Knight{}},
		{Colour: colour.White, Position: move.Position{File: 2, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Bishop{}},
		{Colour: colour.White, Position: move.Position{File: 3, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Queen{}},
		{Colour: colour.White, Position: move.Position{File: 4, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.King{}},
		{Colour: colour.White, Position: move.Position{File: 5, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Bishop{}},
		{Colour: colour.White, Position: move.Position{File: 6, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Knight{}},
		{Colour: colour.White, Position: move.Position{File: 7, Rank: 0}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Rook{}},
		{Colour: colour.White, Position: move.Position{File: 0, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.White}},
		{Colour: colour.White, Position: move.Position{File: 1, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.White}},
		{Colour: colour.White, Position: move.Position{File: 2, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.White}},
		{Colour: colour.White, Position: move.Position{File: 3, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.White}},
		{Colour: colour.White, Position: move.Position{File: 4, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.White}},
		{Colour: colour.White, Position: move.Position{File: 5, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.White}},
		{Colour: colour.White, Position: move.Position{File: 6, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.White}},
		{Colour: colour.White, Position: move.Position{File: 7, Rank: 1}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.White}},
		{Colour: colour.Black, Position: move.Position{File: 0, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Rook{}},
		{Colour: colour.Black, Position: move.Position{File: 1, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Knight{}},
		{Colour: colour.Black, Position: move.Position{File: 2, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Bishop{}},
		{Colour: colour.Black, Position: move.Position{File: 3, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Queen{}},
		{Colour: colour.Black, Position: move.Position{File: 4, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.King{}},
		{Colour: colour.Black, Position: move.Position{File: 5, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Bishop{}},
		{Colour: colour.Black, Position: move.Position{File: 6, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Knight{}},
		{Colour: colour.Black, Position: move.Position{File: 7, Rank: 7}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Rook{}},
		{Colour: colour.Black, Position: move.Position{File: 0, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.Black}},
		{Colour: colour.Black, Position: move.Position{File: 1, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.Black}},
		{Colour: colour.Black, Position: move.Position{File: 2, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.Black}},
		{Colour: colour.Black, Position: move.Position{File: 3, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.Black}},
		{Colour: colour.Black, Position: move.Position{File: 4, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.Black}},
		{Colour: colour.Black, Position: move.Position{File: 5, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.Black}},
		{Colour: colour.Black, Position: move.Position{File: 6, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.Black}},
		{Colour: colour.Black, Position: move.Position{File: 7, Rank: 6}, ValidMoves: make(map[move.Position]bool), PieceDetails: piece.Pawn{Colour: colour.Black}},
	}

	return standardPieces
}

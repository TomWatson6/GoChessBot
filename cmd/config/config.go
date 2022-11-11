package config

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func GetStandardPieces() []*piece.Piece {
	standardPieces := []*piece.Piece{
		{Colour: colour.White, Position: move.Position{File: 0, Rank: 0}, PieceDetails: piece.NewRook()},
		{Colour: colour.White, Position: move.Position{File: 1, Rank: 0}, PieceDetails: piece.NewKnight()},
		{Colour: colour.White, Position: move.Position{File: 2, Rank: 0}, PieceDetails: piece.NewBishop()},
		{Colour: colour.White, Position: move.Position{File: 3, Rank: 0}, PieceDetails: piece.NewQueen()},
		{Colour: colour.White, Position: move.Position{File: 4, Rank: 0}, PieceDetails: piece.NewKing()},
		{Colour: colour.White, Position: move.Position{File: 5, Rank: 0}, PieceDetails: piece.NewBishop()},
		{Colour: colour.White, Position: move.Position{File: 6, Rank: 0}, PieceDetails: piece.NewKnight()},
		{Colour: colour.White, Position: move.Position{File: 7, Rank: 0}, PieceDetails: piece.NewRook()},
		{Colour: colour.White, Position: move.Position{File: 0, Rank: 1}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.White))},
		{Colour: colour.White, Position: move.Position{File: 1, Rank: 1}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.White))},
		{Colour: colour.White, Position: move.Position{File: 2, Rank: 1}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.White))},
		{Colour: colour.White, Position: move.Position{File: 3, Rank: 1}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.White))},
		{Colour: colour.White, Position: move.Position{File: 4, Rank: 1}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.White))},
		{Colour: colour.White, Position: move.Position{File: 5, Rank: 1}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.White))},
		{Colour: colour.White, Position: move.Position{File: 6, Rank: 1}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.White))},
		{Colour: colour.White, Position: move.Position{File: 7, Rank: 1}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.White))},
		{Colour: colour.Black, Position: move.Position{File: 0, Rank: 7}, PieceDetails: piece.NewRook()},
		{Colour: colour.Black, Position: move.Position{File: 1, Rank: 7}, PieceDetails: piece.NewKnight()},
		{Colour: colour.Black, Position: move.Position{File: 2, Rank: 7}, PieceDetails: piece.NewBishop()},
		{Colour: colour.Black, Position: move.Position{File: 3, Rank: 7}, PieceDetails: piece.NewQueen()},
		{Colour: colour.Black, Position: move.Position{File: 4, Rank: 7}, PieceDetails: piece.NewKing()},
		{Colour: colour.Black, Position: move.Position{File: 5, Rank: 7}, PieceDetails: piece.NewBishop()},
		{Colour: colour.Black, Position: move.Position{File: 6, Rank: 7}, PieceDetails: piece.NewKnight()},
		{Colour: colour.Black, Position: move.Position{File: 7, Rank: 7}, PieceDetails: piece.NewRook()},
		{Colour: colour.Black, Position: move.Position{File: 0, Rank: 6}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.Black))},
		{Colour: colour.Black, Position: move.Position{File: 1, Rank: 6}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.Black))},
		{Colour: colour.Black, Position: move.Position{File: 2, Rank: 6}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.Black))},
		{Colour: colour.Black, Position: move.Position{File: 3, Rank: 6}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.Black))},
		{Colour: colour.Black, Position: move.Position{File: 4, Rank: 6}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.Black))},
		{Colour: colour.Black, Position: move.Position{File: 5, Rank: 6}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.Black))},
		{Colour: colour.Black, Position: move.Position{File: 6, Rank: 6}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.Black))},
		{Colour: colour.Black, Position: move.Position{File: 7, Rank: 6}, PieceDetails: piece.NewPawn(piece.PawnWithColour(colour.Black))},
	}

	return standardPieces
}

func GetBoardDimensions() (int, int) {
	return 8, 8
}

func GetNumPieceTypes() int {
	return 6 // Number of piece types in standard chess
}

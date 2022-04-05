package board_test

import (
	"testing"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/pkg/input"
)

func TestMovePiece(t *testing.T) {
	alterations := make(map[move.Position]piece.Piece)
	alterations[move.Position{File: 0, Rank: 2}] = piece.Queen{
		Colour:   colour.White,
		Position: move.Position{File: 0, Rank: 2},
	}

	game := input.Get(colour.White, alterations)
	b := game.Board

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## ##
	// 4 ## ## ## ## ## ## ## ##
	// 3 wQ ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP wP wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	cases := []struct {
		move  move.Move
		check move.Position
		want  piece.Piece
	}{
		{
			move.Move{
				From: move.Position{File: 0, Rank: 2},
				To:   move.Position{File: 1, Rank: 3},
			},
			move.Position{File: 1, Rank: 3},
			piece.Queen{
				Colour:   colour.White,
				Position: move.Position{File: 1, Rank: 3},
			},
		},
		{
			move.Move{
				From: move.Position{File: 1, Rank: 3},
				To:   move.Position{File: 1, Rank: 6},
			},
			move.Position{File: 1, Rank: 6},
			piece.Queen{
				Colour:   colour.White,
				Position: move.Position{File: 1, Rank: 6},
			},
		},
		{
			move.Move{
				From: move.Position{File: 1, Rank: 6},
				To:   move.Position{File: 1, Rank: 5},
			},
			move.Position{File: 1, Rank: 6},
			nil,
		},
	}

	for _, c := range cases {
		b.MovePiece(c.move)
		got := b.GetPiece(c.check)
		if got != c.want {
			t.Errorf("MovePiece(%v) => Piece at position %v == %v, want %v", c.move, c.check, got, c.want)
		}
	}
}

func TestIsValidMove(t *testing.T) {
	alterations := make(map[move.Position]piece.Piece)
	alterations[move.Position{File: 0, Rank: 2}] = piece.Queen{
		Colour:   colour.White,
		Position: move.Position{File: 0, Rank: 2},
	}

	game := input.Get(colour.White, alterations)
	b := game.Board

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## ##
	// 4 ## ## ## ## ## ## ## ##
	// 3 wQ ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP wP wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	cases := []struct {
		move move.Move
		want bool
	}{
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 1, Rank: 3}}, true},
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 1, Rank: 4}}, false},
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 0, Rank: 3}}, true},
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 4, Rank: 6}}, true},
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 5, Rank: 7}}, false},
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 0, Rank: 1}}, false},
		{move.Move{From: move.Position{File: 1, Rank: 0}, To: move.Position{File: 2, Rank: 0}}, false},
		{move.Move{From: move.Position{File: 1, Rank: 0}, To: move.Position{File: 2, Rank: 2}}, true},
	}

	for _, c := range cases {
		got := b.IsValidMove(c.move)
		if got != c.want {
			t.Errorf("IsValidMove(%v) == %v, want %v", c.move, got, c.want)
		}
	}
}

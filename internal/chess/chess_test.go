package chess_test

import (
	"testing"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/pkg/input"
)

func TestTranslateNotation(t *testing.T) {
	alterations := make(map[move.Position]piece.Piece)
	alterations[move.Position{File: 2, Rank: 2}] = piece.Queen{
		Colour:   colour.Black,
		Position: move.Position{File: 2, Rank: 2},
	}
	alterations[move.Position{File: 6, Rank: 4}] = piece.Knight{
		Colour:   colour.White,
		Position: move.Position{File: 6, Rank: 4},
	}
	alterations[move.Position{File: 2, Rank: 4}] = piece.Knight{
		Colour:   colour.White,
		Position: move.Position{File: 2, Rank: 4},
	}
	alterations[move.Position{File: 5, Rank: 2}] = piece.Bishop{
		Colour:   colour.Black,
		Position: move.Position{File: 5, Rank: 2},
	}

	ch := input.Get(colour.White, alterations)

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## wN ## ## ## wN ##
	// 4 ## ## ## ## ## ## ## ##
	// 3 ## ## bQ ## ## bB ## ##
	// 2 wP wP wP wP wP wP wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	cases := []struct {
		notation string
		want     []move.Move
	}{
		{
			notation: "e4",
			want: []move.Move{
				{
					From: move.Position{File: 4, Rank: 1},
					To:   move.Position{File: 4, Rank: 3},
				},
			},
		},
		{
			notation: "d2xc3",
			want: []move.Move{
				{
					From: move.Position{File: 3, Rank: 1},
					To:   move.Position{File: 2, Rank: 2},
				},
			},
		},
		{
			notation: "Nc3",
			want: []move.Move{
				{
					From: move.Position{File: 1, Rank: 0},
					To:   move.Position{File: 2, Rank: 2},
				},
			},
		},
		{
			notation: "g5Nxf3",
			want: []move.Move{
				{
					From: move.Position{File: 6, Rank: 4},
					To:   move.Position{File: 5, Rank: 2},
				},
			},
		},
		{
			notation: "c5Ne4",
			want: []move.Move{
				{
					From: move.Position{File: 2, Rank: 4},
					To:   move.Position{File: 4, Rank: 3},
				},
			},
		},
		{
			notation: "O-O",
			want: []move.Move{
				{
					From: move.Position{File: 4, Rank: 0},
					To:   move.Position{File: 6, Rank: 0},
				},
				{
					From: move.Position{File: 7, Rank: 0},
					To:   move.Position{File: 5, Rank: 0},
				},
			},
		},
		{
			notation: "O-O-O",
			want: []move.Move{
				{
					From: move.Position{File: 4, Rank: 0},
					To:   move.Position{File: 2, Rank: 0},
				},
				{
					From: move.Position{File: 0, Rank: 0},
					To:   move.Position{File: 3, Rank: 0},
				},
			},
		},
	}

	for _, c := range cases {
		got, err := ch.TranslateNotation(c.notation)
		if err != nil {
			t.Errorf("TranslateNotation(%v) returned error: %v", c.notation, err)
		}

		if len(got) != len(c.want) {
			t.Errorf("TranslateNotation(%v) returned %v moves, want %v", c.notation, len(got), len(c.want))
		}

		for i, g := range got {
			if g != c.want[i] {
				t.Errorf("TranslateNotation(%v) returned %v, want %v", c.notation, g, c.want[i])
			}
		}
	}
}

package chess_test

import (
	"testing"

	"github.com/tomwatson6/chessbot/internal/chess"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/testing/payloads"
)

func TestTranslateNotation(t *testing.T) {
	b := payloads.NewStandardBoard(
		payloads.BoardWithPiece(piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 2, Rank: 2},
			ValidMoves:   make(map[move.Position]bool),
			PieceDetails: piece.Queen{},
		}),
		payloads.BoardWithPiece(piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 6, Rank: 4},
			ValidMoves:   make(map[move.Position]bool),
			PieceDetails: piece.Knight{},
		}),
		payloads.BoardWithPiece(piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 2, Rank: 4},
			ValidMoves:   make(map[move.Position]bool),
			PieceDetails: piece.Knight{},
		}),
		payloads.BoardWithPiece(piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 5, Rank: 2},
			ValidMoves:   make(map[move.Position]bool),
			PieceDetails: piece.Bishop{},
		}),
	)

	ch := payloads.NewStandardChessGame(
		payloads.ChessGameWithTurn(colour.White),
		payloads.ChessGameWithBoard(b),
	)

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

func TestNextTurn(t *testing.T) {
	tcs := []struct {
		game chess.Chess
		want colour.Colour
	}{
		{
			game: payloads.NewStandardChessGame(
				payloads.ChessGameWithTurn(colour.White),
			),
			want: colour.Black,
		},
		{
			game: payloads.NewStandardChessGame(
				payloads.ChessGameWithTurn(colour.Black),
			),
			want: colour.White,
		},
	}

	for _, tc := range tcs {
		tc := tc // Rebind tc to this lexical scope
		t.Run(tc.game.Turn.String(), func(t *testing.T) {
			t.Parallel()
			tc.game.NextTurn()

			if tc.game.Turn != tc.want {
				t.Errorf("NextTurn() returned %s, want %s", tc.game.Turn.String(), tc.want.String())
			}
		})
	}
}

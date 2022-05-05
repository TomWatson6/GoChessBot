package board_test

import (
	"testing"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/pkg/input"
)

func TestIsCheck(t *testing.T) {
	// Alter board to setup check scenario
	alterations := make(map[move.Position]piece.Piece)
	alterations[move.Position{File: 7, Rank: 3}] = piece.Piece{
		Colour:       colour.Black,
		Position:     move.Position{File: 7, Rank: 3},
		ValidMoves:   make(map[move.Position]bool),
		PieceDetails: piece.Bishop{},
	}

	deletions := []move.Position{
		{File: 5, Rank: 1},
		{File: 4, Rank: 1},
	}

	game := input.Get(colour.White, alterations, deletions)
	b := game.Board
	b.Update()

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## ##
	// 4 ## ## ## ## ## ## ## bB
	// 3 ## ## ## ## ## ## ## ##
	// 2 wP wP wP wP ## ## wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	cases := []struct {
		colour colour.Colour
		want   bool
		// Add error case to test when king doesn't exist
	}{
		{colour.White, true},
		{colour.Black, false},
	}

	for _, c := range cases {
		c := c // rebind c into this lexical scope
		t.Run(c.colour.String(), func(t *testing.T) {
			t.Parallel()
			got := b.IsCheck(c.colour)
			if got != c.want {
				t.Errorf("IsCheck(%v) => %v, want %v", c.colour, got, c.want)
			}
		})
	}
}

func TestIsCheckMate(t *testing.T) {
	// Alter board to setup checkmate scenario
	alterations := make(map[move.Position]piece.Piece)
	alterations[move.Position{File: 7, Rank: 3}] = piece.Piece{
		Colour:       colour.Black,
		Position:     move.Position{File: 7, Rank: 3},
		ValidMoves:   make(map[move.Position]bool),
		PieceDetails: piece.Bishop{},
	}

	deletions := []move.Position{
		{File: 5, Rank: 1},
	}

	game := input.Get(colour.White, alterations, deletions)
	b := game.Board
	b.Update()

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## ##
	// 4 ## ## ## ## ## ## ## bB
	// 3 ## ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP ## wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	cases := []struct {
		colour colour.Colour
		want   bool
	}{
		{colour.White, true},
		{colour.Black, false},
	}

	for _, c := range cases {
		c := c // rebind c into this lexical scope
		t.Run(c.colour.String(), func(t *testing.T) {
			t.Parallel()
			got := b.IsCheckMate(c.colour)
			if got != c.want {
				t.Errorf("IsCheckMate(%v) => %v, want %v", c.colour, got, c.want)
			}
		})
	}
}

func TestIsNotCheckMate(t *testing.T) {
	// Alter board to setup non-checkmate scenario
	alterations := make(map[move.Position]piece.Piece)
	alterations[move.Position{File: 7, Rank: 3}] = piece.Piece{
		Colour:       colour.Black,
		Position:     move.Position{File: 7, Rank: 3},
		ValidMoves:   make(map[move.Position]bool),
		PieceDetails: piece.Bishop{},
	}

	deletions := []move.Position{
		{File: 5, Rank: 1},
		{File: 4, Rank: 1},
	}

	game := input.Get(colour.White, alterations, deletions)
	b := game.Board
	b.Update()

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## ##
	// 4 ## ## ## ## ## ## ## bB
	// 3 ## ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP ## wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	cases := []struct {
		colour colour.Colour
		want   bool
	}{
		{colour.White, false},
		{colour.Black, false},
	}

	for _, c := range cases {
		c := c // rebind c into this lexical scope
		t.Run(c.colour.String(), func(t *testing.T) {
			t.Parallel()
			got := b.IsCheckMate(c.colour)
			if got != c.want {
				t.Errorf("IsCheckMate(%v) => %v, want %v", c.colour, got, c.want)
			}
		})
	}
}

func TestMovePiece(t *testing.T) {
	alterations := make(map[move.Position]piece.Piece)
	alterations[move.Position{File: 0, Rank: 2}] = piece.Piece{
		Colour:       colour.White,
		Position:     move.Position{File: 0, Rank: 2},
		ValidMoves:   make(map[move.Position]bool),
		PieceDetails: piece.Queen{},
	}

	game := input.Get(colour.White, alterations, []move.Position{})
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
			piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 1, Rank: 3},
				PieceDetails: piece.Queen{},
			},
		},
		{
			move.Move{
				From: move.Position{File: 1, Rank: 3},
				To:   move.Position{File: 1, Rank: 6},
			},
			move.Position{File: 1, Rank: 6},
			piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 1, Rank: 6},
				PieceDetails: piece.Queen{},
			},
		},
		{
			move.Move{
				From: move.Position{File: 1, Rank: 6},
				To:   move.Position{File: 1, Rank: 5},
			},
			move.Position{File: 1, Rank: 6},
			piece.Piece{},
		},
		{
			move.Move{
				From: move.Position{File: 0, Rank: 1},
				To:   move.Position{File: 0, Rank: 3},
			},
			move.Position{File: 0, Rank: 3},
			piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 0, Rank: 3},
				PieceDetails: piece.Pawn{},
			},
		},
		{
			move.Move{
				From: move.Position{File: 1, Rank: 0},
				To:   move.Position{File: 3, Rank: 1},
			},
			move.Position{File: 3, Rank: 1},
			piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 3, Rank: 1},
				PieceDetails: piece.Pawn{},
			},
		},
	}

	for _, c := range cases {
		b.MovePiece(c.move)

		if got, ok := b.Pieces[c.check]; ok {
			if !got.Equals(c.want) {
				t.Errorf("MovePiece(%v) => Piece at position %v == %#v, want %#v",
					c.move, c.check, got, c.want)
			}
		} else {
			if c.want.PieceDetails != nil {
				t.Errorf("MovePiece(%v) => Piece at position %v == %#v, want %#v",
					c.move, c.check, got, c.want)
			}
		}
	}
}

func TestIsValidMove(t *testing.T) {
	alterations := make(map[move.Position]piece.Piece)
	alterations[move.Position{File: 0, Rank: 2}] = piece.Piece{
		Colour:       colour.White,
		Position:     move.Position{File: 0, Rank: 2},
		ValidMoves:   make(map[move.Position]bool),
		PieceDetails: piece.Queen{},
	}

	game := input.Get(colour.White, alterations, []move.Position{})
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
		c := c // rebind c into this lexical scope
		t.Run(c.move.String(), func(t *testing.T) {
			t.Parallel()
			got := b.IsValidMove(c.move)
			if got != c.want {
				t.Errorf("IsValidMove(%v) => %v, want %v", c.move, got, c.want)
			}
		})
	}
}

package piece_test

import (
	"testing"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func TestBishopValidMoves(t *testing.T) {
	b := piece.Piece{
		Colour:       colour.White,
		Position:     move.Position{File: 3, Rank: 3},
		PieceDetails: piece.Bishop{},
	}

	cases := []struct {
		m    move.Move
		want bool
	}{
		{move.Move{From: b.Position, To: move.Position{File: 4, Rank: 4}}, true},
		{move.Move{From: b.Position, To: move.Position{File: 4, Rank: 2}}, true},
		{move.Move{From: b.Position, To: move.Position{File: 5, Rank: 1}}, true},
		{move.Move{From: b.Position, To: move.Position{File: 3, Rank: 2}}, false},
		{move.Move{From: b.Position, To: move.Position{File: 7, Rank: 0}}, false},
	}

	for _, c := range cases {
		got := b.IsValidMove(c.m)
		if got != c.want {
			t.Errorf("IsValidMove(%v) from %v == %v, want %v", c.m, b.Position, got, c.want)
		}
	}
}

func TestKingValidMoves(t *testing.T) {
	k := piece.Piece{
		Colour:       colour.White,
		Position:     move.Position{File: 3, Rank: 3},
		PieceDetails: piece.King{},
	}

	cases := []struct {
		m    move.Move
		want bool
	}{
		{move.Move{From: k.Position, To: move.Position{File: 3, Rank: 4}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 3, Rank: 2}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 4, Rank: 3}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 2, Rank: 3}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 4, Rank: 4}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 2, Rank: 2}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 4, Rank: 2}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 2, Rank: 4}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 4, Rank: 1}}, false},
		{move.Move{From: k.Position, To: move.Position{File: 5, Rank: 0}}, false},
	}

	for _, c := range cases {
		got := k.IsValidMove(c.m)
		if got != c.want {
			t.Errorf("IsValidMove(%v) from %v == %v, want %v", c.m, k.Position, got, c.want)
		}
	}
}

func TestKnightValidMoves(t *testing.T) {
	k := piece.Piece{
		Colour:       colour.White,
		Position:     move.Position{File: 3, Rank: 3},
		PieceDetails: piece.Knight{},
	}

	cases := []struct {
		m    move.Move
		want bool
	}{
		{move.Move{From: k.Position, To: move.Position{File: 5, Rank: 4}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 2, Rank: 5}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 1, Rank: 4}}, true},
		{move.Move{From: k.Position, To: move.Position{File: 4, Rank: 4}}, false},
		{move.Move{From: k.Position, To: move.Position{File: 5, Rank: 5}}, false},
	}

	for _, c := range cases {
		got := k.IsValidMove(c.m)
		if got != c.want {
			t.Errorf("IsValidMove(%v) from %v == %v, want %v", c.m, k.Position, got, c.want)
		}
	}
}

func TestValidPawnMoves(t *testing.T) {
	movedPawn := piece.Piece{
		Colour:       colour.White,
		Position:     move.Position{File: 3, Rank: 3},
		PieceDetails: piece.Pawn{HasMoved: true},
	}

	unmovedPawn := piece.Piece{
		Colour:       colour.White,
		Position:     move.Position{File: 3, Rank: 2},
		PieceDetails: piece.Pawn{},
	}

	cases := []struct {
		pawn piece.Piece
		m    move.Move
		want bool
	}{
		{movedPawn, move.Move{From: movedPawn.Position, To: move.Position{File: 3, Rank: 4}}, true},
		{movedPawn, move.Move{From: movedPawn.Position, To: move.Position{File: 3, Rank: 5}}, false},
		{unmovedPawn, move.Move{From: unmovedPawn.Position, To: move.Position{File: 3, Rank: 3}}, true},
		{unmovedPawn, move.Move{From: unmovedPawn.Position, To: move.Position{File: 3, Rank: 4}}, true},
		{unmovedPawn, move.Move{From: unmovedPawn.Position, To: move.Position{File: 3, Rank: 5}}, false},
	}

	for _, c := range cases {
		got := c.pawn.IsValidMove(c.m)
		if got != c.want {
			t.Errorf("IsValidMove(%v) from %v == %v, want %v", c.m, c.pawn.Position, got, c.want)
		}
	}
}

func TestQueenValidMoves(t *testing.T) {
	q := piece.Piece{
		Colour:       colour.White,
		Position:     move.Position{File: 3, Rank: 3},
		PieceDetails: piece.Queen{},
	}

	cases := []struct {
		m    move.Move
		want bool
	}{
		{move.Move{From: q.Position, To: move.Position{File: 4, Rank: 4}}, true},
		{move.Move{From: q.Position, To: move.Position{File: 4, Rank: 2}}, true},
		{move.Move{From: q.Position, To: move.Position{File: 5, Rank: 1}}, true},
		{move.Move{From: q.Position, To: move.Position{File: 3, Rank: 2}}, true},
		{move.Move{From: q.Position, To: move.Position{File: 6, Rank: 0}}, true},
		{move.Move{From: q.Position, To: move.Position{File: 4, Rank: 3}}, true},
		{move.Move{From: q.Position, To: move.Position{File: 2, Rank: 3}}, true},
		{move.Move{From: q.Position, To: move.Position{File: 2, Rank: 4}}, true},
		{move.Move{From: q.Position, To: move.Position{File: 4, Rank: 5}}, false},
		{move.Move{From: q.Position, To: move.Position{File: 5, Rank: 0}}, false},
		{move.Move{From: q.Position, To: move.Position{File: 7, Rank: 0}}, false},
	}

	for _, c := range cases {
		got := q.IsValidMove(c.m)
		if got != c.want {
			t.Errorf("IsValidMove(%v) from %v == %v, want %v", c.m, q.Position, got, c.want)
		}
	}
}

func TestRookValidMoves(t *testing.T) {
	r := piece.Piece{
		Colour:       colour.White,
		Position:     move.Position{File: 3, Rank: 3},
		PieceDetails: piece.Rook{},
	}

	cases := []struct {
		m    move.Move
		want bool
	}{
		{move.Move{From: r.Position, To: move.Position{File: 3, Rank: 4}}, true},
		{move.Move{From: r.Position, To: move.Position{File: 3, Rank: 2}}, true},
		{move.Move{From: r.Position, To: move.Position{File: 4, Rank: 3}}, true},
		{move.Move{From: r.Position, To: move.Position{File: 2, Rank: 3}}, true},
		{move.Move{From: r.Position, To: move.Position{File: 4, Rank: 4}}, false},
		{move.Move{From: r.Position, To: move.Position{File: 2, Rank: 2}}, false},
		{move.Move{From: r.Position, To: move.Position{File: 4, Rank: 2}}, false},
		{move.Move{From: r.Position, To: move.Position{File: 2, Rank: 4}}, false},
	}

	for _, c := range cases {
		got := r.IsValidMove(c.m)
		if got != c.want {
			t.Errorf("IsValidMove(%v) from %v == %v, want %v", c.m, r.Position, got, c.want)
		}
	}
}

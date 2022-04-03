package piece_test

import (
	"testing"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func TestBishopValidMoves(t *testing.T) {
	b := piece.Bishop{
		Colour:   colour.White,
		Position: move.Position{File: 3, Rank: 3},
	}

	cases := []struct {
		dest move.Position
		want bool
	}{
		{move.Position{File: 4, Rank: 4}, true},
		{move.Position{File: 4, Rank: 2}, true},
		{move.Position{File: 5, Rank: 1}, true},
		{move.Position{File: 3, Rank: 2}, false},
		{move.Position{File: 7, Rank: 0}, false},
	}

	for _, c := range cases {
		got := b.IsValidMove(c.dest)
		if got != c.want {
			t.Errorf("IsValidMove(%v) == %v, want %v", c.dest, got, c.want)
		}
	}
}

func TestKingValidMoves(t *testing.T) {
	k := piece.King{
		Colour:   colour.White,
		Position: move.Position{File: 3, Rank: 3},
	}

	cases := []struct {
		dest move.Position
		want bool
	}{
		{move.Position{File: 3, Rank: 4}, true},
		{move.Position{File: 3, Rank: 2}, true},
		{move.Position{File: 4, Rank: 3}, true},
		{move.Position{File: 2, Rank: 3}, true},
		{move.Position{File: 4, Rank: 4}, true},
		{move.Position{File: 2, Rank: 2}, true},
		{move.Position{File: 4, Rank: 2}, true},
		{move.Position{File: 2, Rank: 4}, true},
		{move.Position{File: 4, Rank: 1}, false},
		{move.Position{File: 5, Rank: 0}, false},
	}

	for _, c := range cases {
		got := k.IsValidMove(c.dest)
		if got != c.want {
			t.Errorf("IsValidMove(%v) == %v, want %v", c.dest, got, c.want)
		}
	}
}

func TestKnightValidMoves(t *testing.T) {
	k := piece.Knight{
		Colour:   colour.White,
		Position: move.Position{File: 3, Rank: 3},
	}

	cases := []struct {
		dest move.Position
		want bool
	}{
		{move.Position{File: 5, Rank: 4}, true},
		{move.Position{File: 2, Rank: 5}, true},
		{move.Position{File: 1, Rank: 4}, true},
		{move.Position{File: 4, Rank: 4}, false},
		{move.Position{File: 5, Rank: 5}, false},
	}

	for _, c := range cases {
		got := k.IsValidMove(c.dest)
		if got != c.want {
			t.Errorf("IsValidMove(%v) == %v, want %v", c.dest, got, c.want)
		}
	}
}

func TestValidPawnMoves(t *testing.T) {
	movedPawn := piece.Pawn{
		Colour:   colour.White,
		Position: move.Position{File: 3, Rank: 3},
		HasMoved: true,
	}

	unmovedPawn := piece.Pawn{
		Colour:   colour.White,
		Position: move.Position{File: 3, Rank: 2},
	}

	cases := []struct {
		pawn piece.Pawn
		dest move.Position
		want bool
	}{
		{movedPawn, move.Position{File: 3, Rank: 4}, true},
		{movedPawn, move.Position{File: 3, Rank: 5}, false},
		{unmovedPawn, move.Position{File: 3, Rank: 3}, true},
		{unmovedPawn, move.Position{File: 3, Rank: 4}, true},
		{unmovedPawn, move.Position{File: 3, Rank: 5}, false},
	}

	for _, c := range cases {
		got := c.pawn.IsValidMove(c.dest)
		if got != c.want {
			t.Errorf("IsValidMove(%v) == %v, want %v", c.dest, got, c.want)
		}
	}
}

func TestQueenValidMoves(t *testing.T) {
	q := piece.Queen{
		Colour:   colour.White,
		Position: move.Position{File: 3, Rank: 3},
	}

	cases := []struct {
		dest move.Position
		want bool
	}{
		{move.Position{File: 4, Rank: 4}, true},
		{move.Position{File: 4, Rank: 2}, true},
		{move.Position{File: 5, Rank: 1}, true},
		{move.Position{File: 3, Rank: 2}, true},
		{move.Position{File: 6, Rank: 0}, true},
		{move.Position{File: 4, Rank: 3}, true},
		{move.Position{File: 2, Rank: 3}, true},
		{move.Position{File: 2, Rank: 4}, true},
		{move.Position{File: 4, Rank: 5}, false},
		{move.Position{File: 5, Rank: 0}, false},
		{move.Position{File: 7, Rank: 0}, false},
	}

	for _, c := range cases {
		got := q.IsValidMove(c.dest)
		if got != c.want {
			t.Errorf("IsValidMove(%v) == %v, want %v", c.dest, got, c.want)
		}
	}
}

func TestRookValidMoves(t *testing.T) {
	r := piece.Rook{
		Colour:   colour.White,
		Position: move.Position{File: 3, Rank: 3},
	}

	cases := []struct {
		dest move.Position
		want bool
	}{
		{move.Position{File: 3, Rank: 4}, true},
		{move.Position{File: 3, Rank: 2}, true},
		{move.Position{File: 4, Rank: 3}, true},
		{move.Position{File: 2, Rank: 3}, true},
		{move.Position{File: 4, Rank: 4}, false},
		{move.Position{File: 2, Rank: 2}, false},
		{move.Position{File: 4, Rank: 2}, false},
		{move.Position{File: 2, Rank: 4}, false},
	}

	for _, c := range cases {
		got := r.IsValidMove(c.dest)
		if got != c.want {
			t.Errorf("IsValidMove(%v) == %v, want %v", c.dest, got, c.want)
		}
	}
}

package piece_test

import (
	"github.com/tomwatson6/chessbot/internal/piece/rules"
	"testing"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func TestBishopValidMoves(t *testing.T) {
	b := piece.Piece{
		Colour:       colour.White,
		Position:     move.Position{File: 3, Rank: 3},
		PieceDetails: piece.NewBishop(),
	}

	cases := []struct {
		m    move.Move
		want error
	}{
		{move.Move{From: b.Position, To: move.Position{File: 4, Rank: 4}}, nil},
		{move.Move{From: b.Position, To: move.Position{File: 4, Rank: 2}}, nil},
		{move.Move{From: b.Position, To: move.Position{File: 5, Rank: 1}}, nil},
		{move.Move{From: b.Position, To: move.Position{File: 3, Rank: 2}}, rules.ErrorIsNotDiagonalLine},
		{move.Move{From: b.Position, To: move.Position{File: 7, Rank: 0}}, rules.ErrorIsNotValidLine},
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
		PieceDetails: piece.NewKing(),
	}

	cases := []struct {
		m    move.Move
		want error
	}{
		{move.Move{From: k.Position, To: move.Position{File: 3, Rank: 4}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 3, Rank: 2}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 4, Rank: 3}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 2, Rank: 3}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 4, Rank: 4}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 2, Rank: 2}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 4, Rank: 2}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 2, Rank: 4}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 4, Rank: 1}}, rules.ErrorIsNotValidLine},
		{move.Move{From: k.Position, To: move.Position{File: 5, Rank: 0}}, rules.ErrorIsNotValidLine},
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
		PieceDetails: piece.NewKnight(),
	}

	cases := []struct {
		m    move.Move
		want error
	}{
		{move.Move{From: k.Position, To: move.Position{File: 5, Rank: 4}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 2, Rank: 5}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 1, Rank: 4}}, nil},
		{move.Move{From: k.Position, To: move.Position{File: 4, Rank: 4}}, rules.ErrorIsNotValidKnightsMove},
		{move.Move{From: k.Position, To: move.Position{File: 5, Rank: 5}}, rules.ErrorIsNotValidKnightsMove},
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
		Colour:   colour.White,
		Position: move.Position{File: 3, Rank: 3},
		PieceDetails: piece.NewPawn(
			piece.PawnWithColour(colour.White),
			piece.PawnWithHasMoved(true),
		),
	}

	unmovedPawn := piece.Piece{
		Colour:   colour.White,
		Position: move.Position{File: 3, Rank: 2},
		PieceDetails: piece.NewPawn(
			piece.PawnWithColour(colour.White),
		),
	}

	cases := []struct {
		pawn piece.Piece
		m    move.Move
		want error
	}{
		{movedPawn, move.Move{From: movedPawn.Position, To: move.Position{File: 3, Rank: 4}}, nil},
		{movedPawn, move.Move{From: movedPawn.Position, To: move.Position{File: 3, Rank: 5}}, rules.ErrorExceedsMaxRange},
		{unmovedPawn, move.Move{From: unmovedPawn.Position, To: move.Position{File: 3, Rank: 3}}, nil},
		{unmovedPawn, move.Move{From: unmovedPawn.Position, To: move.Position{File: 3, Rank: 4}}, nil},
		{unmovedPawn, move.Move{From: unmovedPawn.Position, To: move.Position{File: 3, Rank: 5}}, rules.ErrorExceedsMaxRange},
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
		PieceDetails: piece.NewQueen(),
	}

	cases := []struct {
		m    move.Move
		want error
	}{
		{move.Move{From: q.Position, To: move.Position{File: 4, Rank: 4}}, nil},
		{move.Move{From: q.Position, To: move.Position{File: 4, Rank: 2}}, nil},
		{move.Move{From: q.Position, To: move.Position{File: 5, Rank: 1}}, nil},
		{move.Move{From: q.Position, To: move.Position{File: 3, Rank: 2}}, nil},
		{move.Move{From: q.Position, To: move.Position{File: 6, Rank: 0}}, nil},
		{move.Move{From: q.Position, To: move.Position{File: 4, Rank: 3}}, nil},
		{move.Move{From: q.Position, To: move.Position{File: 2, Rank: 3}}, nil},
		{move.Move{From: q.Position, To: move.Position{File: 2, Rank: 4}}, nil},
		{move.Move{From: q.Position, To: move.Position{File: 4, Rank: 5}}, rules.ErrorIsNotValidLine},
		{move.Move{From: q.Position, To: move.Position{File: 5, Rank: 0}}, rules.ErrorIsNotValidLine},
		{move.Move{From: q.Position, To: move.Position{File: 7, Rank: 0}}, rules.ErrorIsNotValidLine},
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
		PieceDetails: piece.NewRook(),
	}

	cases := []struct {
		m    move.Move
		want error
	}{
		{move.Move{From: r.Position, To: move.Position{File: 3, Rank: 4}}, nil},
		{move.Move{From: r.Position, To: move.Position{File: 3, Rank: 2}}, nil},
		{move.Move{From: r.Position, To: move.Position{File: 4, Rank: 3}}, nil},
		{move.Move{From: r.Position, To: move.Position{File: 2, Rank: 3}}, nil},
		{move.Move{From: r.Position, To: move.Position{File: 4, Rank: 4}}, rules.ErrorIsDiagonalLine},
		{move.Move{From: r.Position, To: move.Position{File: 2, Rank: 2}}, rules.ErrorIsDiagonalLine},
		{move.Move{From: r.Position, To: move.Position{File: 4, Rank: 2}}, rules.ErrorIsDiagonalLine},
		{move.Move{From: r.Position, To: move.Position{File: 2, Rank: 4}}, rules.ErrorIsDiagonalLine},
	}

	for _, c := range cases {
		got := r.IsValidMove(c.m)
		if got != c.want {
			t.Errorf("IsValidMove(%v) from %v == %v, want %v", c.m, r.Position, got, c.want)
		}
	}
}

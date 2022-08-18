package board_test

import (
	"github.com/tomwatson6/chessbot/internal/board"
	"testing"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/testing/payloads"
)

func TestIsCheck(t *testing.T) {
	b := payloads.NewStandardBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 7, Rank: 3},
			PieceDetails: piece.Bishop{},
		}),
		payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 1}),
		payloads.BoardWithDeletedPiece(move.Position{File: 4, Rank: 1}),
	)

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
		// TODO: Add error case to test when king doesn't exist
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
	b := payloads.NewStandardBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 7, Rank: 3},
			PieceDetails: piece.Bishop{},
		}),
		payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 1}),
	)

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
	b := payloads.NewStandardBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 7, Rank: 3},
			PieceDetails: piece.Bishop{},
		}),
		payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 1}),
		payloads.BoardWithDeletedPiece(move.Position{File: 4, Rank: 1}),
	)

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
	b := payloads.NewStandardBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 0, Rank: 2},
			PieceDetails: piece.Queen{},
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 4, Rank: 4},
			PieceDetails: piece.Queen{},
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 4, Rank: 6},
			PieceDetails: piece.Bishop{},
		}),
	)

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bB bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## wQ ## ## ##
	// 4 ## ## ## ## ## ## ## ##
	// 3 wQ ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP wP wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	cases := []struct {
		name  string
		move  move.Move
		check move.Position
		want  piece.Piece
	}{
		{
			name: "WhiteQueenDiagonalMove",
			move: move.Move{
				From: move.Position{File: 0, Rank: 2},
				To:   move.Position{File: 1, Rank: 3},
			},
			check: move.Position{File: 1, Rank: 3},
			want: piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 1, Rank: 3},
				PieceDetails: piece.Queen{},
			},
		},
		{
			name: "WhiteQueenCapturesBlackPawn",
			move: move.Move{
				From: move.Position{File: 1, Rank: 3},
				To:   move.Position{File: 1, Rank: 6},
			},
			check: move.Position{File: 1, Rank: 6},
			want: piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 1, Rank: 6},
				PieceDetails: piece.Queen{},
			},
		},
		{
			name: "WhiteQueenStandardMoveBackwards",
			move: move.Move{
				From: move.Position{File: 1, Rank: 6},
				To:   move.Position{File: 1, Rank: 5},
			},
			check: move.Position{File: 1, Rank: 6},
			want:  piece.Piece{},
		},
		{
			name: "WhitePawnDoubleMove",
			move: move.Move{
				From: move.Position{File: 0, Rank: 1},
				To:   move.Position{File: 0, Rank: 3},
			},
			check: move.Position{File: 0, Rank: 3},
			want: piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 0, Rank: 3},
				PieceDetails: piece.Pawn{},
			},
		},
		{
			name: "InvalidWhiteKnightMove",
			move: move.Move{
				From: move.Position{File: 1, Rank: 0},
				To:   move.Position{File: 3, Rank: 1},
			},
			check: move.Position{File: 3, Rank: 1},
			want: piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 3, Rank: 1},
				PieceDetails: piece.Pawn{},
			},
		},
		{
			name: "InvalidBishopMoveWhenPinned",
			move: move.Move{
				From: move.Position{File: 4, Rank: 6},
				To:   move.Position{File: 3, Rank: 5},
			},
			check: move.Position{File: 4, Rank: 6},
			want: piece.Piece{
				Colour:       colour.Black,
				Position:     move.Position{File: 4, Rank: 6},
				PieceDetails: piece.Bishop{},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b.Move(c.move)

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
		})
	}
}

func TestIsValidMove(t *testing.T) {
	history := []board.Turn{
		{
			colour.White: move.Move{
				From: move.Position{File: 7, Rank: 6},
				To:   move.Position{File: 7, Rank: 4},
			},
		},
	}

	b := payloads.NewStandardBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 0, Rank: 2},
			PieceDetails: piece.Queen{},
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:   colour.White,
			Position: move.Position{File: 6, Rank: 4},
			PieceDetails: piece.Pawn{
				HasMoved: true,
			},
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:   colour.White,
			Position: move.Position{File: 7, Rank: 4},
			PieceDetails: piece.Pawn{
				HasMoved: true,
			},
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:   colour.White,
			Position: move.Position{File: 5, Rank: 4},
			PieceDetails: piece.Pawn{
				HasMoved: true,
			},
		}),
		payloads.BoardWithDeletedPiece(move.Position{
			File: 6,
			Rank: 1,
		}),
		payloads.BoardWithDeletedPiece(move.Position{
			File: 7,
			Rank: 6,
		}),
		payloads.BoardWithDeletedPiece(move.Position{
			File: 5,
			Rank: 6,
		}),
		payloads.BoardWithHistory(history),
	)

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP ## bP ##
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## bP wP bP
	// 4 ## ## ## ## ## ## ## ##
	// 3 wQ ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP wP ## wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	cases := []struct {
		move  move.Move
		valid bool
	}{
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 1, Rank: 3}}, true},
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 1, Rank: 4}}, false},
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 0, Rank: 3}}, true},
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 4, Rank: 6}}, true},
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 5, Rank: 7}}, false},
		{move.Move{From: move.Position{File: 0, Rank: 2}, To: move.Position{File: 0, Rank: 1}}, false},
		{move.Move{From: move.Position{File: 1, Rank: 0}, To: move.Position{File: 2, Rank: 0}}, false},
		{move.Move{From: move.Position{File: 1, Rank: 0}, To: move.Position{File: 2, Rank: 2}}, true},
		{move.Move{From: move.Position{File: 6, Rank: 4}, To: move.Position{File: 7, Rank: 5}}, true},
		{move.Move{From: move.Position{File: 6, Rank: 4}, To: move.Position{File: 5, Rank: 5}}, false},
	}

	for _, c := range cases {
		c := c // rebind c into this lexical scope
		t.Run(c.move.String(), func(t *testing.T) {
			t.Parallel()
			err := b.IsValidMove(c.move)

			if (err == nil && !c.valid) || (err != nil && c.valid) {
				t.Errorf("IsValidMove(%v) => %v, want %v", c.move, err == nil, c.valid)
			}
		})
	}
}

func TestPawnHasMoved(t *testing.T) {
	t.Parallel()

	b := payloads.NewStandardBoard()

	hasMoved := b.Pieces[move.Position{File: 0, Rank: 1}].PieceDetails.(piece.Pawn).HasMoved
	if hasMoved {
		t.Errorf("HasMoved is %t, expected %t", hasMoved, false)
	}

	from := move.Position{File: 0, Rank: 1}
	to := move.Position{File: 0, Rank: 3}

	m := move.Move{From: from, To: to}

	err := b.Move(m)
	if err != nil {
		t.Errorf("MovePiece with move %v failed", m)
	}

	hasMoved = b.Pieces[to].PieceDetails.(piece.Pawn).HasMoved
	if !hasMoved {
		t.Errorf("HasMoved is %t, expected %t", hasMoved, true)
	}
}

// TODO: In random chess game, pieces can capture other pieces of their own colour
// TODO: When in check, it doesn't realise it can take the threatening piece

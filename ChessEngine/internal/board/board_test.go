package board_test

import (
	"reflect"
	"testing"

	"github.com/tomwatson6/chessbot/internal/board"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/testing/payloads"
)

func BenchmarkGetValidMoves(b *testing.B) {
	bo := payloads.NewStandardBoard()

	for i := 0; i < b.N; i++ {
		moves := bo.GetValidMoves()

		if len(moves) != 40 {
			b.Fatalf("Got incorrect number of moves, got: %d, expected: %d\n", len(moves), 40)
		}
	}
}

func TestGetValidMoves(t *testing.T) {
	bo := payloads.NewStandardBoard()

	moves := bo.GetValidMoves()
	expected := 40

	if len(moves) != expected {
		t.Fatalf("Number of valid moves incorrect -> expected: %d, got: %d", expected, len(moves))
	}
}

func TestIsCheck(t *testing.T) {
	b := payloads.NewStandardBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 7, Rank: 3},
			PieceDetails: piece.NewBishop(),
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
			_, got, err := b.IsCheck(c.colour)
			if err != nil {
				t.Fatal(err)
			}

			if got != c.want {
				t.Errorf("IsCheck(%v) => %v, want %v", c.colour, got, c.want)
			}
		})
	}
}

func TestIsCheckMate(t *testing.T) {
	tcs := []struct {
		name  string
		b     board.Board
		white bool
		black bool
	}{
		{
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

			name: "BishopPlacingWhiteInCheckWithBlockingMove",
			b: payloads.NewStandardBoard(
				payloads.BoardWithPiece(&piece.Piece{
					Colour:       colour.Black,
					Position:     move.Position{File: 7, Rank: 3},
					PieceDetails: piece.NewBishop(),
				}),
				payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 1}),
			),
			white: false,
			black: false,
		},
		{
			// Visualisation of the board
			// 8 bR bN bB bQ bK bB bN bR
			// 7 bP bP bP bP bP bP bP bP
			// 6 ## ## ## ## ## ## ## ##
			// 5 ## ## ## ## ## ## ## ##
			// 4 ## ## ## ## ## ## ## bB
			// 3 ## ## ## ## ## ## ## ##
			// 2 wP wP wP wP ## ## ## wP
			// 1 wR wN wB wQ wK wB wN wR
			//    A  B  C  D  E  F  G  H

			name: "BishopPlacingWhiteInCheckWithKingEscape",
			b: payloads.NewStandardBoard(
				payloads.BoardWithPiece(&piece.Piece{
					Colour:       colour.Black,
					Position:     move.Position{File: 7, Rank: 3},
					PieceDetails: piece.NewBishop(),
				}),
				payloads.BoardWithDeletedPiece(move.Position{File: 4, Rank: 1}),
				payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 1}),
				payloads.BoardWithDeletedPiece(move.Position{File: 6, Rank: 1}),
			),
			white: false,
			black: false,
		},
		{
			// Visualisation of the board
			// 8 bR bN bB bQ bK bB bN bR
			// 7 bP bP bP bP bP bP bP bP
			// 6 ## ## ## ## ## ## ## ##
			// 5 ## ## ## ## ## ## ## ##
			// 4 wR ## ## ## ## ## ## bB
			// 3 ## ## ## ## ## ## ## ##
			// 2 wP wP wP wP wP ## ## wP
			// 1 wR wN wB wQ wK wB wN wR
			//    A  B  C  D  E  F  G  H

			name: "AttackingBlackVulnerableBishop",
			b: payloads.NewStandardBoard(
				payloads.BoardWithPiece(&piece.Piece{
					Colour:       colour.Black,
					Position:     move.Position{File: 7, Rank: 3},
					PieceDetails: piece.NewBishop(),
				}),
				payloads.BoardWithPiece(&piece.Piece{
					Colour:       colour.White,
					Position:     move.Position{File: 0, Rank: 3},
					PieceDetails: piece.NewRook(),
				}),
				payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 1}),
				payloads.BoardWithDeletedPiece(move.Position{File: 6, Rank: 1}),
			),
			white: false,
			black: false,
		},
		{
			// Visualisation of the board
			// 8 bR bN bB bQ bK bB bN bR
			// 7 bP bP bP bP bP ## bP bP
			// 6 ## ## ## ## ## ## ## ##
			// 5 ## ## ## ## ## ## ## wB
			// 4 ## ## ## ## ## ## ## ##
			// 3 ## ## ## ## ## ## ## ##
			// 2 wP wP wP wP wP wP wP wP
			// 1 wR wN wB wQ wK wB wN wR
			//    A  B  C  D  E  F  G  H

			name: "BishopPlacingBlackInCheckWithBlockingMove",
			b: payloads.NewStandardBoard(
				payloads.BoardWithPiece(&piece.Piece{
					Colour:       colour.White,
					Position:     move.Position{File: 7, Rank: 4},
					PieceDetails: piece.NewBishop(),
				}),
				payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 6}),
			),
			white: false,
			black: false,
		},
		{
			// Visualisation of the board
			// 8 bR bN bB bQ bK bB bN bR
			// 7 bP bP bP bP ## ## ## bP
			// 6 ## ## ## ## ## ## ## ##
			// 5 ## ## ## ## ## ## ## wB
			// 4 ## ## ## ## ## ## ## ##
			// 3 ## ## ## ## ## ## ## ##
			// 2 wP wP wP wP wP wP wP wP
			// 1 wR wN wB wQ wK wB wN wR
			//    A  B  C  D  E  F  G  H

			name: "BishopPlacingBlackInCheckWithKingEscape",
			b: payloads.NewStandardBoard(
				payloads.BoardWithPiece(&piece.Piece{
					Colour:       colour.White,
					Position:     move.Position{File: 7, Rank: 4},
					PieceDetails: piece.NewBishop(),
				}),
				payloads.BoardWithDeletedPiece(move.Position{File: 4, Rank: 6}),
				payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 6}),
				payloads.BoardWithDeletedPiece(move.Position{File: 6, Rank: 6}),
			),
			white: false,
			black: false,
		},
		{
			// Visualisation of the board
			// 8 bR bN bB bQ bK bB bN bR
			// 7 bP bP bP bP bP ## ## bP
			// 6 ## ## ## ## ## ## ## ##
			// 5 bR ## ## ## ## ## ## wB
			// 4 ## ## ## ## ## ## ## ##
			// 3 ## ## ## ## ## ## ## ##
			// 2 wP wP wP wP wP wP wP wP
			// 1 wR wN wB wQ wK wB wN wR
			//    A  B  C  D  E  F  G  H

			name: "AttackingWhiteVulnerableBishop",
			b: payloads.NewStandardBoard(
				payloads.BoardWithPiece(&piece.Piece{
					Colour:       colour.White,
					Position:     move.Position{File: 7, Rank: 4},
					PieceDetails: piece.NewBishop(),
				}),
				payloads.BoardWithPiece(&piece.Piece{
					Colour:       colour.Black,
					Position:     move.Position{File: 0, Rank: 4},
					PieceDetails: piece.NewRook(),
				}),
				payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 6}),
				payloads.BoardWithDeletedPiece(move.Position{File: 6, Rank: 6}),
			),
			white: false,
			black: false,
		},
	}

	for _, c := range tcs {
		c := c // rebind c into this lexical scope
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			w, err := c.b.IsCheckMate(colour.White)
			if err != nil {
				t.Fatal(err)
			}

			b, err := c.b.IsCheckMate(colour.Black)
			if err != nil {
				t.Fatal(err)
			}

			if w != c.white {
				t.Errorf("IsCheckMate(%v) => %v, want %v", colour.White, w, c.white)
			}

			if b != c.black {
				t.Errorf("IsCheckMate(%v) => %v, want %v", colour.Black, b, c.black)
			}
		})
	}
}

func TestMovePiece(t *testing.T) {
	b := payloads.NewStandardBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 0, Rank: 2},
			PieceDetails: piece.NewQueen(),
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 4, Rank: 4},
			PieceDetails: piece.NewQueen(),
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 4, Rank: 6},
			PieceDetails: piece.NewBishop(),
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

	// Visualisation of the board after
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP ## bP bP bB bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## wQ ## ## ##
	// 4 ## ## ## ## ## ## ## ##
	// 3 ## ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP wP wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	cases := []struct {
		name  string
		move  move.Move
		check move.Position
		want  *piece.Piece
	}{
		{
			name: "WhiteQueenDiagonalMove",
			move: move.Move{
				From: move.Position{File: 0, Rank: 2},
				To:   move.Position{File: 1, Rank: 3},
			},
			check: move.Position{File: 1, Rank: 3},
			want: &piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 1, Rank: 3},
				PieceDetails: piece.NewQueen(),
			},
		},
		{
			name: "WhiteQueenCapturesBlackPawn",
			move: move.Move{
				From: move.Position{File: 1, Rank: 3},
				To:   move.Position{File: 1, Rank: 6},
			},
			check: move.Position{File: 1, Rank: 6},
			want: &piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 1, Rank: 6},
				PieceDetails: piece.NewQueen(),
			},
		},
		{
			name: "WhiteQueenStandardMoveBackwards",
			move: move.Move{
				From: move.Position{File: 1, Rank: 6},
				To:   move.Position{File: 1, Rank: 5},
			},
			check: move.Position{File: 1, Rank: 6},
			want:  nil,
		},
		{
			name: "WhitePawnDoubleMove",
			move: move.Move{
				From: move.Position{File: 0, Rank: 1},
				To:   move.Position{File: 0, Rank: 3},
			},
			check: move.Position{File: 0, Rank: 3},
			want: &piece.Piece{
				Colour:   colour.White,
				Position: move.Position{File: 0, Rank: 3},
				PieceDetails: piece.NewPawn(
					piece.PawnWithHasMoved(true),
				),
			},
		},
		{
			name: "InvalidWhiteKnightMove",
			move: move.Move{
				From: move.Position{File: 1, Rank: 0},
				To:   move.Position{File: 3, Rank: 1},
			},
			check: move.Position{File: 3, Rank: 1},
			want: &piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 3, Rank: 1},
				PieceDetails: piece.NewPawn(),
			},
		},
		{
			name: "InvalidBishopMoveWhenPinned",
			move: move.Move{
				From: move.Position{File: 4, Rank: 6},
				To:   move.Position{File: 3, Rank: 5},
			},
			check: move.Position{File: 4, Rank: 6},
			want: &piece.Piece{
				Colour:       colour.Black,
				Position:     move.Position{File: 4, Rank: 6},
				PieceDetails: piece.NewBishop(),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b.Move(c.move)

			if got, ok := b.Pieces[c.check]; ok {
				if !reflect.DeepEqual(got, c.want) {
					t.Errorf("MovePiece(%v) => Piece at position %v == %#v, want %#v",
						c.move, c.check, got, c.want)
				}
			} else {
				if c.want != nil {
					t.Errorf("MovePiece(%v) => Piece at position %v == %#v, want %#v",
						c.move, c.check, got, c.want)
				}
			}
		})
	}
}

func TestValidWhiteKingSideCastlingMove(t *testing.T) {
	b := payloads.NewStandardBoard(
		payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 0}),
		payloads.BoardWithDeletedPiece(move.Position{File: 6, Rank: 0}),
	)

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## ##
	// 4 ## ## ## ## ## ## ## ##
	// 3 ## ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP wP wP wP
	// 1 wR wN wB wQ wK ## ## wR
	//    A  B  C  D  E  F  G  H

	m := move.Move{
		From: move.Position{File: 4, Rank: 0},
		To:   move.Position{File: 6, Rank: 0},
	}

	if _, err := b.Move(m); err != nil {
		t.Fatalf("King side castling move invalid, failed with error: %s\n", err)
	}

	k, ok := b.Pieces[m.To]
	if !ok {
		t.Fatalf("There is no piece in destination square for the king")
	}

	if k.GetPieceType() != piece.PieceTypeKing || !k.HasMoved() || k.Position != m.To {
		t.Fatalf("King details at destination square are incorrect, got: %#v\n", k)
	}

	rDest := move.Position{File: 5, Rank: 0}

	r, ok := b.Pieces[rDest]
	if !ok {
		t.Fatalf("There is no piece in destination square for the rook")
	}

	if r.GetPieceType() != piece.PieceTypeRook || !r.HasMoved() || r.Position != rDest {
		t.Fatalf("Rook details at destination square are incorrect, got: %#v\n", r)
	}
}

func TestValidWhiteQueenSideCastlingMove(t *testing.T) {
	b := payloads.NewStandardBoard(
		payloads.BoardWithDeletedPiece(move.Position{File: 1, Rank: 0}),
		payloads.BoardWithDeletedPiece(move.Position{File: 2, Rank: 0}),
		payloads.BoardWithDeletedPiece(move.Position{File: 3, Rank: 0}),
	)

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## ##
	// 4 ## ## ## ## ## ## ## ##
	// 3 ## ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP wP wP wP
	// 1 wR ## ## ## wK wB wN wR
	//    A  B  C  D  E  F  G  H

	m := move.Move{
		From: move.Position{File: 4, Rank: 0},
		To:   move.Position{File: 2, Rank: 0},
	}

	if _, err := b.Move(m); err != nil {
		t.Fatalf("King side castling move invalid, failed with error: %s\n", err)
	}

	k, ok := b.Pieces[m.To]
	if !ok {
		t.Fatalf("There is no piece in destination square for the king")
	}

	if k.GetPieceType() != piece.PieceTypeKing || !k.HasMoved() || k.Position != m.To {
		t.Fatalf("King details at destination square are incorrect, got: %#v\n", k)
	}

	rDest := move.Position{File: 3, Rank: 0}

	r, ok := b.Pieces[rDest]
	if !ok {
		t.Fatalf("There is no piece in destination square for the rook")
	}

	if r.GetPieceType() != piece.PieceTypeRook || !r.HasMoved() || r.Position != rDest {
		t.Fatalf("Rook details at destination square are incorrect, got: %#v\n", r)
	}
}

func TestValidBlackKingSideCastlingMove(t *testing.T) {
	b := payloads.NewStandardBoard(
		payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 7}),
		payloads.BoardWithDeletedPiece(move.Position{File: 6, Rank: 7}),
	)

	// Visualisation of the board
	// 8 bR bN bB bQ bK ## ## bR
	// 7 bP bP bP bP bP bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## ##
	// 4 ## ## ## ## ## ## ## ##
	// 3 ## ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP wP wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	m := move.Move{
		From: move.Position{File: 4, Rank: 7},
		To:   move.Position{File: 6, Rank: 7},
	}

	if _, err := b.Move(m); err != nil {
		t.Fatalf("King side castling move invalid, failed with error: %s\n", err)
	}

	k, ok := b.Pieces[m.To]
	if !ok {
		t.Fatalf("There is no piece in destination square for the king")
	}

	if k.GetPieceType() != piece.PieceTypeKing || !k.HasMoved() || k.Position != m.To {
		t.Fatalf("King details at destination square are incorrect, got: %#v\n", k)
	}

	rDest := move.Position{File: 5, Rank: 7}

	r, ok := b.Pieces[rDest]
	if !ok {
		t.Fatalf("There is no piece in destination square for the rook")
	}

	if r.GetPieceType() != piece.PieceTypeRook || !r.HasMoved() || r.Position != rDest {
		t.Fatalf("Rook details at destination square are incorrect, got: %#v\n", r)
	}
}

func TestValidBlackQueenSideCastlingMove(t *testing.T) {
	b := payloads.NewStandardBoard(
		payloads.BoardWithDeletedPiece(move.Position{File: 1, Rank: 7}),
		payloads.BoardWithDeletedPiece(move.Position{File: 2, Rank: 7}),
		payloads.BoardWithDeletedPiece(move.Position{File: 3, Rank: 7}),
	)

	// Visualisation of the board
	// 8 bR ## ## ## bK bB bN bR
	// 7 bP bP bP bP bP bP bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## ##
	// 4 ## ## ## ## ## ## ## ##
	// 3 ## ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP wP wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	m := move.Move{
		From: move.Position{File: 4, Rank: 7},
		To:   move.Position{File: 2, Rank: 7},
	}

	if _, err := b.Move(m); err != nil {
		t.Fatalf("King side castling move invalid, failed with error: %s\n", err)
	}

	k, ok := b.Pieces[m.To]
	if !ok {
		t.Fatalf("There is no piece in destination square for the king")
	}

	if k.GetPieceType() != piece.PieceTypeKing || !k.HasMoved() || k.Position != m.To {
		t.Fatalf("King details at destination square are incorrect, got: %#v\n", k)
	}

	rDest := move.Position{File: 3, Rank: 7}

	r, ok := b.Pieces[rDest]
	if !ok {
		t.Fatalf("There is no piece in destination square for the rook")
	}

	if r.GetPieceType() != piece.PieceTypeRook || !r.HasMoved() || r.Position != rDest {
		t.Fatalf("Rook details at destination square are incorrect, got: %#v\n", r)
	}
}

func TestIsValidMove(t *testing.T) {
	history := []board.Turn{
		{
			colour.Black: &move.Move{
				From: move.Position{File: 7, Rank: 6},
				To:   move.Position{File: 7, Rank: 4},
			},
		},
	}

	b := payloads.NewStandardBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 0, Rank: 2},
			PieceDetails: piece.NewQueen(),
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:   colour.White,
			Position: move.Position{File: 6, Rank: 4},
			PieceDetails: piece.NewPawn(
				piece.PawnWithHasMoved(true),
			),
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:   colour.Black,
			Position: move.Position{File: 7, Rank: 4},
			PieceDetails: piece.NewPawn(
				piece.PawnWithHasMoved(true),
			),
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:   colour.Black,
			Position: move.Position{File: 5, Rank: 4},
			PieceDetails: piece.NewPawn(
				piece.PawnWithHasMoved(true),
			),
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:   colour.Black,
			Position: move.Position{File: 1, Rank: 4},
			PieceDetails: piece.NewPawn(
				piece.PawnWithHasMoved(true),
			),
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 5, Rank: 3},
			PieceDetails: piece.NewBishop(),
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
		payloads.BoardWithDeletedPiece(move.Position{
			File: 3,
			Rank: 1,
		}),
		payloads.BoardWithHistory(history),
	)

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP ## bP ##
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## wP ## ## ## bP wP bP
	// 4 ## ## ## ## ## bB ## ##
	// 3 wQ ## ## ## ## ## ## ##
	// 2 wP wP wP ## wP wP ## wP
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
		{move.Move{From: move.Position{File: 3, Rank: 1}, To: move.Position{File: 3, Rank: 4}}, false},
		{move.Move{From: move.Position{File: 0, Rank: 1}, To: move.Position{File: 0, Rank: 2}}, false},
		{move.Move{From: move.Position{File: 3, Rank: 6}, To: move.Position{File: 1, Rank: 4}}, false},
		{move.Move{From: move.Position{File: 4, Rank: 0}, To: move.Position{File: 3, Rank: 1}}, false},
	}

	for _, c := range cases {
		c := c // rebind c into this lexical scope
		t.Run(c.move.String(), func(t *testing.T) {
			t.Parallel()
			err := b.IsValidMove(c.move)

			if (err == nil && !c.valid) || (err != nil && c.valid) {
				t.Errorf("IsValidMove(%v) => %v, want %v", c.move, err, c.valid)
			}
		})
	}
}

func TestPawnHasMoved(t *testing.T) {
	t.Parallel()

	b := payloads.NewStandardBoard()

	hasMoved := b.Pieces[move.Position{File: 0, Rank: 1}].PieceDetails.HasMoved()
	if hasMoved {
		t.Errorf("HasMoved is %t, expected %t", hasMoved, false)
	}

	from := move.Position{File: 0, Rank: 1}
	to := move.Position{File: 0, Rank: 3}

	m := move.Move{From: from, To: to}

	_, err := b.Move(m)
	if err != nil {
		t.Errorf("MovePiece with move %v failed", m)
	}

	hasMoved = b.Pieces[to].PieceDetails.HasMoved()
	if !hasMoved {
		t.Errorf("HasMoved is %t, expected %t", hasMoved, true)
	}
}

// TODO: Look at what the hell this does...
func TestCanTakeKingWhenInCheck(t *testing.T) {
	t.Parallel()

	history := []board.Turn{
		{
			colour.White: &move.Move{
				From: move.Position{File: 6, Rank: 3},
				To:   move.Position{File: 7, Rank: 4},
			},
		},
	}

	b := payloads.NewStandardBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 7, Rank: 4},
			PieceDetails: piece.NewBishop(),
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 5, Rank: 6},
			PieceDetails: piece.NewBishop(),
		}),
		payloads.BoardWithHistory(history),
	)

	// Visualisation of the board
	// 8 bR bN bB bQ bK bB bN bR
	// 7 bP bP bP bP bP bB bP bP
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## wB
	// 4 ## ## ## ## ## ## ## ##
	// 3 ## ## ## ## ## ## ## ##
	// 2 wP wP wP wP wP wP wP wP
	// 1 wR wN wB wQ wK wB wN wR
	//    A  B  C  D  E  F  G  H

	p := b.Pieces[move.Position{File: 5, Rank: 6}]

	if err := b.IsValidMove(move.Move{From: p.Position, To: move.Position{File: 4, Rank: 5}}); err == nil {
		t.Fatalf("expected error, got nil, bishop should not be able to move out of a pinned position exposing it's king")
	}

	from := move.Position{File: 5, Rank: 6}
	to := move.Position{File: 7, Rank: 4}
	m := move.Move{From: from, To: to}

	if _, err := b.Move(m); err != nil {
		t.Fatalf("failed with error: %s", err)
	}
}

func TestPawnPromotion(t *testing.T) {
	t.Parallel()

	b := payloads.NewEmptyBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 3, Rank: 6},
			PieceDetails: piece.NewPawn(),
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 3, Rank: 1},
			PieceDetails: piece.NewPawn(),
		}),
	)

	tcs := []struct {
		name    string
		move    move.Move
		details piece.PieceDetails
		want    *piece.Piece
	}{
		{
			name: "WhitePawnPromotion",
			move: move.Move{
				From: move.Position{File: 3, Rank: 6},
				To:   move.Position{File: 3, Rank: 7},
			},
			details: piece.NewQueen(),
			want: &piece.Piece{
				Colour:       colour.White,
				Position:     move.Position{File: 3, Rank: 7},
				PieceDetails: piece.NewQueen(),
			},
		},
		{
			name: "BlackPawnPromotion",
			move: move.Move{
				From: move.Position{File: 3, Rank: 1},
				To:   move.Position{File: 3, Rank: 0},
			},
			details: piece.NewBishop(),
			want: &piece.Piece{
				Colour:       colour.Black,
				Position:     move.Position{File: 3, Rank: 0},
				PieceDetails: piece.NewBishop(),
			},
		},
	}

	// Visualisation of the board
	// 8 ## ## ## ## ## ## ## ##
	// 7 ## ## ## wP ## ## ## ##
	// 6 ## ## ## ## ## ## ## ##
	// 5 ## ## ## ## ## ## ## ##
	// 4 ## ## ## ## ## ## ## ##
	// 3 ## ## ## ## ## ## ## ##
	// 2 ## ## ## bP ## ## ## ##
	// 1 ## ## ## ## ## ## ## ##
	//    A  B  C  D  E  F  G  H

	for _, c := range tcs {
		c := c // rebind t into this lexical scope
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			b.Promote(c.move, c.details)

			if got, ok := b.Pieces[c.move.To]; ok {
				if !reflect.DeepEqual(got, c.want) {
					t.Errorf("Promote(%v, %v) => Piece at position %v == %#v, want %#v",
						c.move, c.details, c.move.To, got, c.want)
				}
			} else {
				t.Errorf("Promote(%v, %v) => Piece at position %v == %#v, want %#v",
					c.move, c.details, c.move.To, got, c.want)
			}
		})
	}
}

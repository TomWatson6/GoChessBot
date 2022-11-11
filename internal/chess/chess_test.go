package chess_test

import (
	"testing"

	"github.com/tomwatson6/chessbot/internal/chess"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/testing/payloads"
)

func TestSimpleMoves(t *testing.T) {
	b := payloads.NewStandardBoard()
	c := payloads.NewStandardChessGame(
		payloads.ChessGameWithBoard(b),
		payloads.ChessGameWithTurn(colour.White),
	)

	tcs := []struct {
		m move.Move
	}{
		{
			move.Move{
				From: move.Position{File: 0, Rank: 1},
				To:   move.Position{File: 0, Rank: 3},
			},
		},
		{
			move.Move{
				From: move.Position{File: 1, Rank: 6},
				To:   move.Position{File: 1, Rank: 4},
			},
		},
	}

	for _, tc := range tcs {
		if err := c.MakeMove(tc.m); err != nil {
			t.Fatalf("MakeMove() returned error: %s", err)
		}
		c.NextTurn()
	}
}

// func TestTranslateNotation(t *testing.T) {
// 	b := payloads.NewStandardBoard(
// 		payloads.BoardWithPiece(&piece.Piece{
// 			Colour:       colour.Black,
// 			Position:     move.Position{File: 2, Rank: 2},
// 			PieceDetails: piece.Queen{},
// 		}),
// 		payloads.BoardWithPiece(&piece.Piece{
// 			Colour:       colour.White,
// 			Position:     move.Position{File: 6, Rank: 4},
// 			PieceDetails: piece.Knight{},
// 		}),
// 		payloads.BoardWithPiece(&piece.Piece{
// 			Colour:       colour.White,
// 			Position:     move.Position{File: 2, Rank: 4},
// 			PieceDetails: piece.Knight{},
// 		}),
// 		payloads.BoardWithPiece(&piece.Piece{
// 			Colour:       colour.Black,
// 			Position:     move.Position{File: 5, Rank: 2},
// 			PieceDetails: piece.Bishop{},
// 		}),
// 		payloads.BoardWithPiece(&piece.Piece{
// 			Colour:   colour.White,
// 			Position: move.Position{File: 2, Rank: 6},
// 			PieceDetails: piece.Pawn{
// 				HasMoved: true,
// 			},
// 		}),
// 		payloads.BoardWithPiece(&piece.Piece{
// 			Colour:   colour.White,
// 			Position: move.Position{File: 7, Rank: 6},
// 			PieceDetails: piece.Pawn{
// 				HasMoved: true,
// 			},
// 		}),
// 		payloads.BoardWithDeletedPiece(move.Position{File: 2, Rank: 7}),
// 	)

// 	ch := payloads.NewStandardChessGame(
// 		payloads.ChessGameWithTurn(colour.White),
// 		payloads.ChessGameWithBoard(b),
// 	)

// 	// Visualisation of the board
// 	// 8 bR bN ## bQ bK bB bN bR
// 	// 7 bP bP wP bP bP bP bP ##
// 	// 6 ## ## ## ## ## ## ## ##
// 	// 5 ## ## wN ## ## ## wN ##
// 	// 4 ## ## ## ## ## ## ## ##
// 	// 3 ## ## bQ ## ## bB ## ##
// 	// 2 wP wP wP wP wP wP wP wP
// 	// 1 wR wN wB wQ wK wB wN wR
// 	//    A  B  C  D  E  F  G  H

// 	cases := []struct {
// 		notation string
// 		want     []move.Move
// 	}{
// 		{
// 			notation: "h3",
// 			want: []move.Move{
// 				{
// 					From: move.Position{File: 7, Rank: 1},
// 					To:   move.Position{File: 7, Rank: 2},
// 				},
// 			},
// 		},
// 		{
// 			notation: "e4",
// 			want: []move.Move{
// 				{
// 					From: move.Position{File: 4, Rank: 1},
// 					To:   move.Position{File: 4, Rank: 3},
// 				},
// 			},
// 		},
// 		{
// 			notation: "dxc3",
// 			want: []move.Move{
// 				{
// 					From: move.Position{File: 3, Rank: 1},
// 					To:   move.Position{File: 2, Rank: 2},
// 				},
// 			},
// 		},
// 		{
// 			notation: "Nc3",
// 			want: []move.Move{
// 				{
// 					From: move.Position{File: 1, Rank: 0},
// 					To:   move.Position{File: 2, Rank: 2},
// 				},
// 			},
// 		},
// 		{
// 			notation: "g5Nxf3",
// 			want: []move.Move{
// 				{
// 					From: move.Position{File: 6, Rank: 4},
// 					To:   move.Position{File: 5, Rank: 2},
// 				},
// 			},
// 		},
// 		{
// 			notation: "c5Ne4",
// 			want: []move.Move{
// 				{
// 					From: move.Position{File: 2, Rank: 4},
// 					To:   move.Position{File: 4, Rank: 3},
// 				},
// 			},
// 		},
// 		{
// 			notation: "c8=Q",
// 			want: []move.Move{
// 				{
// 					From: move.Position{File: 2, Rank: 6},
// 					To:   move.Position{File: 2, Rank: 7},
// 				},
// 			},
// 		},
// 		{
// 			notation: "hxg8=Q",
// 			want: []move.Move{
// 				{
// 					From: move.Position{File: 7, Rank: 6},
// 					To:   move.Position{File: 6, Rank: 7},
// 				},
// 			},
// 		},
// 		{
// 			notation: "O-O",
// 			want: []move.Move{
// 				{
// 					From: move.Position{File: 4, Rank: 0},
// 					To:   move.Position{File: 6, Rank: 0},
// 				},
// 				{
// 					From: move.Position{File: 7, Rank: 0},
// 					To:   move.Position{File: 5, Rank: 0},
// 				},
// 			},
// 		},
// 		{
// 			notation: "O-O-O",
// 			want: []move.Move{
// 				{
// 					From: move.Position{File: 4, Rank: 0},
// 					To:   move.Position{File: 2, Rank: 0},
// 				},
// 				{
// 					From: move.Position{File: 0, Rank: 0},
// 					To:   move.Position{File: 3, Rank: 0},
// 				},
// 			},
// 		},
// 	}

// 	for _, c := range cases {
// 		got, err := ch.TranslateNotation(c.notation)
// 		if err != nil {
// 			t.Errorf("TranslateNotation(%v) returned error: %v", c.notation, err)
// 		}

// 		if len(got) != len(c.want) {
// 			t.Errorf("TranslateNotation(%v) returned %v moves, want %v", c.notation, len(got), len(c.want))
// 		}

// 		for i, g := range got {
// 			if g != c.want[i] {
// 				t.Errorf("TranslateNotation(%v) returned %v, want %v", c.notation, g, c.want[i])
// 			}
// 		}
// 	}
// }

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

// func TestMoveToNotation(t *testing.T) {
// 	b := payloads.NewStandardBoard(
// 		payloads.BoardWithPiece(&piece.Piece{
// 			Colour:       colour.White,
// 			Position:     move.Position{File: 3, Rank: 3},
// 			PieceDetails: piece.Knight{},
// 		}),
// 		payloads.BoardWithPiece(&piece.Piece{
// 			Colour:       colour.Black,
// 			Position:     move.Position{File: 3, Rank: 2},
// 			PieceDetails: piece.Pawn{},
// 		}),
// 		payloads.BoardWithDeletedPiece(move.Position{File: 1, Rank: 7}),
// 		payloads.BoardWithDeletedPiece(move.Position{File: 2, Rank: 7}),
// 		payloads.BoardWithDeletedPiece(move.Position{File: 3, Rank: 7}),
// 		payloads.BoardWithDeletedPiece(move.Position{File: 5, Rank: 7}),
// 		payloads.BoardWithDeletedPiece(move.Position{File: 6, Rank: 7}),
// 	)

// 	c := payloads.NewStandardChessGame(
// 		payloads.ChessGameWithBoard(b),
// 	)

// 	// Visualisation of the board
// 	// 8 bR ## ## ## bK ## ## bR
// 	// 7 bP bP bP bP bP bP bP bP
// 	// 6 ## ## ## ## ## ## ## ##
// 	// 5 ## ## ## ## ## ## ## ##
// 	// 4 ## ## ## wN ## ## ## ##
// 	// 3 ## ## ## bP ## ## ## ##
// 	// 2 wP wP wP wP wP wP wP wP
// 	// 1 wR wN wB wQ wK wB wN wR
// 	//    A  B  C  D  E  F  G  H

// 	tcs := []struct {
// 		ms   []move.Move
// 		want string
// 	}{
// 		{
// 			[]move.Move{
// 				{
// 					From: move.Position{File: 0, Rank: 1},
// 					To:   move.Position{File: 0, Rank: 2},
// 				},
// 			},
// 			"a3",
// 		},
// 		{
// 			[]move.Move{
// 				{
// 					From: move.Position{File: 1, Rank: 0},
// 					To:   move.Position{File: 2, Rank: 2},
// 				},
// 			},
// 			"Nc3",
// 		},
// 		{
// 			[]move.Move{
// 				{
// 					From: move.Position{File: 6, Rank: 0},
// 					To:   move.Position{File: 5, Rank: 2},
// 				},
// 			},
// 			"g1Nf3",
// 		},
// 		{
// 			[]move.Move{
// 				{
// 					From: move.Position{File: 4, Rank: 1},
// 					To:   move.Position{File: 3, Rank: 2},
// 				},
// 			},
// 			"e2xd3",
// 		},
// 		{
// 			[]move.Move{
// 				{
// 					From: move.Position{File: 4, Rank: 7},
// 					To:   move.Position{File: 6, Rank: 7},
// 				},
// 				{
// 					From: move.Position{File: 7, Rank: 7},
// 					To:   move.Position{File: 5, Rank: 7},
// 				},
// 			},
// 			"O-O",
// 		},
// 		{
// 			[]move.Move{
// 				{
// 					From: move.Position{File: 4, Rank: 7},
// 					To:   move.Position{File: 2, Rank: 7},
// 				},
// 				{
// 					From: move.Position{File: 0, Rank: 7},
// 					To:   move.Position{File: 3, Rank: 7},
// 				},
// 			},
// 			"O-O-O",
// 		},
// 	}

// 	for i, tc := range tcs {
// 		tc := tc // rebind to this lexical scope

// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			t.Parallel()
// 			notation, err := c.ToChessNotation(tc.ms)
// 			if err != nil {
// 				t.Errorf("failed to convert to notation: %s", err)
// 			}

// 			if notation != tc.want {
// 				t.Errorf("failed to convert, got: %s, want: %s", notation, tc.want)
// 			}
// 		})
// 	}
// }

// func TestInputGame(t *testing.T) {
// 	b := payloads.NewStandardBoard()
// 	c := payloads.NewStandardChessGame(
// 		payloads.ChessGameWithBoard(b),
// 	)

// 	content, err := os.ReadFile("../../testing/output/2022_07_08 16_19_11.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	split := strings.Split(string(content), "\n\n")

// 	input := split[1]

// 	lines := strings.Split(input, "\n")

// 	for _, line := range lines {
// 		m, err := move.NewMoveFromString(line)
// 		if err != nil {
// 			t.Fatalf("cannot read move: %s", line)
// 		}

// 		if err := c.MakeMove(m); err != nil {
// 			t.Fatalf("invalid move: %+v, failed with error: %s", m, err)
// 		}

// 		c.NextTurn()
// 	}
// }

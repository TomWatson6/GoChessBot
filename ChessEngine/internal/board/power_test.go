package board_test

import (
	"testing"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/testing/payloads"
)

func TestPower(t *testing.T) {
	b := payloads.NewEmptyBoard(
		payloads.BoardWithPiece(
			&piece.Piece{
				Colour:   colour.White,
				Position: move.Position{File: 3, Rank: 7},
				PieceDetails: piece.NewKing(
					piece.KingWithHasMoved(true),
				),
			},
		),
	)

	tcs := []struct {
		name     string
		pos      move.Position
		expected []move.Position
	}{
		{
			name: "King Can Move To Adjacent Squares Only",
			pos:  move.Position{File: 3, Rank: 7},
			expected: []move.Position{
				{File: 2, Rank: 7},
				{File: 4, Rank: 7},
				{File: 2, Rank: 6},
				{File: 3, Rank: 6},
				{File: 4, Rank: 6},
			},
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := b.Power(tc.pos.File, tc.pos.Rank)

			if len(p) != len(tc.expected) {
				t.Errorf("length of power does not match expected, expected: %d, got: %d\n", len(tc.expected), len(p))
			}
		})
	}
}

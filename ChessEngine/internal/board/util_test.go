package board_test

import (
	"testing"

	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/testing/payloads"
)

func TestGetLineExcludingFirstIncludingLast(t *testing.T) {
	t.Parallel()

	b := payloads.NewEmptyBoard()

	tcs := []struct {
		name string
		m    move.Move
		want []move.Position
	}{
		{
			name: "right",
			m: move.Move{
				From: move.Position{File: 0, Rank: 5},
				To:   move.Position{File: 2, Rank: 5},
			},
			want: []move.Position{
				{File: 1, Rank: 5},
				{File: 2, Rank: 5},
			},
		},
		{
			name: "left",
			m: move.Move{
				From: move.Position{File: 2, Rank: 5},
				To:   move.Position{File: 0, Rank: 5},
			},
			want: []move.Position{
				{File: 1, Rank: 5},
				{File: 0, Rank: 5},
			},
		},
		{
			name: "up",
			m: move.Move{
				From: move.Position{File: 2, Rank: 5},
				To:   move.Position{File: 2, Rank: 7},
			},
			want: []move.Position{
				{File: 2, Rank: 6},
				{File: 2, Rank: 7},
			},
		},
		{
			name: "down",
			m: move.Move{
				From: move.Position{File: 2, Rank: 5},
				To:   move.Position{File: 2, Rank: 2},
			},
			want: []move.Position{
				{File: 2, Rank: 4},
				{File: 2, Rank: 3},
				{File: 2, Rank: 2},
			},
		},
		{
			name: "up-left",
			m: move.Move{
				From: move.Position{File: 2, Rank: 5},
				To:   move.Position{File: 0, Rank: 7},
			},
			want: []move.Position{
				{File: 1, Rank: 6},
				{File: 0, Rank: 7},
			},
		},
		{
			name: "up-right",
			m: move.Move{
				From: move.Position{File: 2, Rank: 5},
				To:   move.Position{File: 4, Rank: 7},
			},
			want: []move.Position{
				{File: 3, Rank: 6},
				{File: 4, Rank: 7},
			},
		},
		{
			name: "down-left",
			m: move.Move{
				From: move.Position{File: 2, Rank: 5},
				To:   move.Position{File: 0, Rank: 3},
			},
			want: []move.Position{
				{File: 1, Rank: 4},
				{File: 0, Rank: 3},
			},
		},
		{
			name: "down-right",
			m: move.Move{
				From: move.Position{File: 2, Rank: 5},
				To:   move.Position{File: 4, Rank: 3},
			},
			want: []move.Position{
				{File: 3, Rank: 4},
				{File: 4, Rank: 3},
			},
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := b.GetLine(tc.m.From, tc.m.To, false, true)

			if len(got) != len(tc.want) {
				t.Fatalf("length of slice returned is different, got: %d, want: %d", len(got), len(tc.want))
			}

			for i := range got {
				if got[i].File != tc.want[i].File || got[i].Rank != tc.want[i].Rank {
					t.Errorf("mismatching slice element at index %d, got: %#v, want: %#v\n", i, got, tc.want)
				}
			}
		})
	}
}

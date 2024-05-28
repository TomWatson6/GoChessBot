package generation_test

import (
	"testing"

	"github.com/tomwatson6/chessbot/generation"
)

func TestNewBoard(t *testing.T) {
	tcs := []struct {
		name     string
		value    int
		expected int
	}{
		{
			name:     "five",
			value:    5,
			expected: 5,
		},
		{
			name:     "ten",
			value:    10,
			expected: 10,
		},
		{
			name:     "twenty",
			value:    20,
			expected: 20,
		},
		{
			name:     "thirtytwo",
			value:    32,
			expected: 32,
		},
		{
			name:     "fifty",
			value:    50,
			expected: 32,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := generation.NewBoard(tc.value)
			if len(result.Pieces) != tc.expected {
				t.Errorf("New board has incorrect amount of piece, expected: %d, got: %d\n", tc.expected, len(result.Pieces))
			}
		})
	}
}

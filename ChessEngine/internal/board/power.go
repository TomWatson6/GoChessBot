package board

import (
	"github.com/tomwatson6/chessbot/internal/move"
)

func (b Board) Power(file, rank int) []move.Position {
	power := []move.Position{}

	start := move.Position{File: file, Rank: rank}

	_, ok := b.Pieces[start]
	if !ok {
		return []move.Position{}
	}

	for _, s := range b.Squares {
		m := move.Move{
			From: start,
			To:   s,
		}

		if err := b.IsValidMove(m); err == nil {
			power = append(power, s)
		}
	}

	return power
}

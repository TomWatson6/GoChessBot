package rules

import (
	"github.com/tomwatson6/chessbot/internal/piece"
)

type Rules []Rule

type Rule interface {
	Check(p piece.Piece) bool
}

func (rs Rules) All(p piece.Piece) bool {
	for _, r := range rs {
		if !r.Check(p) {
			return false
		}
	}

	return true
}

func (rs Rules) Any(p piece.Piece) bool {
	for _, r := range rs {
		if r.Check(p) {
			return true
		}
	}

	return false
}

package threat

import (
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

type Threat map[move.Position][]piece.Piece

func (t Threat) Get(pos move.Position) []piece.Piece {
	return t[pos]
}

func (t Threat) GetThreatLevel(pos move.Position) int {
	pieces := t[pos]

	threatLevel := 0

	for _, piece := range pieces {
		threatLevel += piece.GetThreatLevel()
	}

	return threatLevel
}

package piece

import (
	"github.com/tomwatson6/chessbot/internal/piece/rules"
	"math"

	"github.com/tomwatson6/chessbot/internal/move"
)

type Knight struct {
	maxRange int
}

func NewKnight() *Knight {
	return &Knight{
		maxRange: 2,
	}
}

func (k Knight) GetPieceLetter() PieceLetter {
	return PieceLetterKnight
}

func (k Knight) GetPiecePoints() PiecePoints {
	return PiecePointsKnight
}

func (k Knight) GetPieceType() PieceType {
	return PieceTypeKnight
}

func (k *Knight) Move(m move.Move) error {
	rs := rules.Assert(
		rules.IsValidKnightsMove(m),
	)

	if err := rs(); err != nil {
		return err
	}

	return nil
}

func (k Knight) IsValidMove(m move.Move) bool {
	if math.Abs(float64(m.To.File-m.From.File)) == 2 {
		return math.Abs(float64(m.To.Rank-m.From.Rank)) == 1
	} else if math.Abs(float64(m.To.Rank-m.From.Rank)) == 2 {
		return math.Abs(float64(m.To.File-m.From.File)) == 1
	}

	return false
}

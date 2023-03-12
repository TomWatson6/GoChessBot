package piece

import (
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece/rules"
)

type King struct {
	minRange, maxRange int
	hasMoved           bool
}

type KingOption func(k *King)

func NewKing(opts ...KingOption) King {
	k := King{
		minRange: 1,
		maxRange: 1,
	}

	for _, opt := range opts {
		opt(&k)
	}

	return k
}

func KingWithHasMoved(hasMoved bool) KingOption {
	return func(k *King) {
		k.hasMoved = hasMoved
	}
}

func (k King) GetPieceLetter() PieceLetter {
	return PieceLetterKing
}

func (k King) GetPiecePoints() PiecePoints {
	return PiecePointsKing
}

func (k King) GetPieceType() PieceType {
	return PieceTypeKing
}

func (k King) IsValidMove(m move.Move) error {
	max := k.maxRange

	if !k.hasMoved {
		max += 1
	}

	rs := rules.Assert(
		rules.IsValidLine(m),
		rules.IsLargerThanOrEqualToMinRange(k.minRange, m),
		rules.DoesNotExceedMaxRange(max, m),
	)

	if err := rs(); err != nil {
		return err
	}

	return nil
}

func (k King) HasMoved() bool {
	return k.hasMoved
}

//func (k King) IsValidMove(m move.Move) bool {
//	return math.Abs(float64(m.To.File-m.From.File)) <= 1 &&
//		math.Abs(float64(m.To.Rank-m.From.Rank)) <= 1
//}

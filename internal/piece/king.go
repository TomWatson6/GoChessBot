package piece

import (
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece/rules"
)

type King struct {
	minRange, maxRange int
}

func NewKing() King {
	return King{
		minRange: 1,
		maxRange: 1,
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
	rs := rules.Assert(
		rules.IsValidLine(m),
		rules.IsLargerThanOrEqualToMinRange(k.minRange, m),
		rules.DoesNotExceedMaxRange(k.maxRange, m),
	)

	if err := rs(); err != nil {
		return err
	}

	return nil
}

//func (k King) IsValidMove(m move.Move) bool {
//	return math.Abs(float64(m.To.File-m.From.File)) <= 1 &&
//		math.Abs(float64(m.To.Rank-m.From.Rank)) <= 1
//}

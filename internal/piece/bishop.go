package piece

import (
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece/rules"
)

type Bishop struct {
	minRange, maxRange int
}

func NewBishop() Bishop {
	return Bishop{
		minRange: 1,
		maxRange: 8,
	}
}

func (b Bishop) GetPieceLetter() PieceLetter {
	return PieceLetterBishop
}

func (b Bishop) GetPiecePoints() PiecePoints {
	return PiecePointsBishop
}

func (b Bishop) GetPieceType() PieceType {
	return PieceTypeBishop
}

func (b Bishop) IsValidMove(m move.Move) error {
	rs := rules.Assert(
		rules.IsValidLine(m),
		rules.IsLargerThanOrEqualToThanMinRange(b.minRange, m),
		rules.DoesNotExceedMaxRange(b.maxRange, m),
		rules.IsDiagonalLine(m),
	)

	if err := rs(); err != nil {
		return err
	}

	return nil
}

//func (b Bishop) IsValidMove(m move.Move) bool {
//	x := m.To.File - m.From.File
//	y := m.To.Rank - m.From.Rank
//	return (x == y || x == -y) && (x != 0 && y != 0)
//}

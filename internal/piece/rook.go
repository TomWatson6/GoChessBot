package piece

import (
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece/rules"
)

type Rook struct {
	minRange, maxRange int
}

func NewRook() Rook {
	return Rook{
		minRange: 1,
		maxRange: 8,
	}
}

func (r Rook) GetPieceLetter() PieceLetter {
	return PieceLetterRook
}

func (r Rook) GetPiecePoints() PiecePoints {
	return PiecePointsRook
}

func (r Rook) GetPieceType() PieceType {
	return PieceTypeRook
}

func (r Rook) IsValidMove(m move.Move) error {
	rs := rules.Assert(
		rules.IsValidLine(m),
		rules.IsLargerThanOrEqualToMinRange(r.minRange, m),
		rules.DoesNotExceedMaxRange(r.maxRange, m),
		rules.IsNotDiagonalLine(m),
	)

	if err := rs(); err != nil {
		return err
	}

	return nil
}

//func (r Rook) IsValidMove(m move.Move) bool {
//	return m.From.File == m.To.File || m.From.Rank == m.To.Rank
//}

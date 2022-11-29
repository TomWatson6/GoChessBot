package piece

import (
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece/rules"
)

type Queen struct {
	minRange, maxRange int
}

func NewQueen() Queen {
	return Queen{
		minRange: 1,
		maxRange: 8,
	}
}

func (q Queen) GetPieceLetter() PieceLetter {
	return PieceLetterQueen
}

func (q Queen) GetPiecePoints() PiecePoints {
	return PiecePointsQueen
}

func (q Queen) GetPieceType() PieceType {
	return PieceTypeQueen
}

func (q Queen) IsValidMove(m move.Move) error {
	rs := rules.Assert(
		rules.IsValidLine(m),
		rules.IsLargerThanOrEqualToMinRange(q.minRange, m),
		rules.DoesNotExceedMaxRange(q.maxRange, m),
	)

	if err := rs(); err != nil {
		return err
	}

	return nil
}

func (q Queen) HasMoved() bool {
	return true
}

package piece

import (
	"github.com/tomwatson6/chessbot/internal/move"
)

type Bishop struct{}

func (b Bishop) GetPieceLetter() PieceLetter {
	return PieceLetterBishop
}

func (b Bishop) GetPiecePoints() PiecePoints {
	return PiecePointsBishop
}

func (b Bishop) GetPieceType() PieceType {
	return PieceTypeBishop
}

func (b Bishop) IsValidMove(m move.Move) bool {
	x := m.To.File - m.From.File
	y := m.To.Rank - m.From.Rank
	return (x == y || x == -y) && (x != 0 && y != 0)
}

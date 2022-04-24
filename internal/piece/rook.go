package piece

import (
	"github.com/tomwatson6/chessbot/internal/move"
)

type Rook struct{}

func (r Rook) GetPieceLetter() PieceLetter {
	return PieceLetterRook
}

func (r Rook) GetPiecePoints() PiecePoints {
	return PiecePointsRook
}

func (r Rook) GetPieceType() PieceType {
	return PieceTypeRook
}

func (r Rook) IsValidMove(m move.Move) bool {
	return m.From.File == m.To.File || m.From.Rank == m.To.Rank
}

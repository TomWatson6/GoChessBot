package board

import (
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/piece"
)

type Board struct {
	Pieces [][]piece.Piece
}

func (b Board) GetPiece(pos move.Position) piece.Piece {
	return b.Pieces[pos.Rank][pos.File]
}

package board

import (
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

type Board struct {
	Pieces map[move.Position]piece.Piece
}

func (b Board) GetPiece(pos move.Position) piece.Piece {
	return b.Pieces[pos]
}

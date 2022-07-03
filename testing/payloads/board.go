package payloads

import (
	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

type BoardOption func(b *board.Board) error

func NewStandardBoard(opts ...BoardOption) board.Board {
	b := board.New()

	for _, opt := range opts {
		if err := opt(&b); err != nil {
			panic(err)
		}
	}

	b.Update()

	return b
}

func BoardWithPiece(p piece.Piece) BoardOption {
	return BoardOption(func(b *board.Board) error {
		b.Pieces[p.Position] = p
		return nil
	})
}

func BoardWithDeletedPiece(pos move.Position) BoardOption {
	return BoardOption(func(b *board.Board) error {
		delete(b.Pieces, pos)
		return nil
	})
}

func BoardWithMoveNumber(number int) BoardOption {
	return BoardOption(func(b *board.Board) error {
		b.MoveNumber = number
		return nil
	})
}

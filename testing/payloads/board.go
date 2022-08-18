package payloads

import (
	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

type BoardOption func(b *board.Board) error

func NewStandardBoard(opts ...BoardOption) board.Board {
	b := board.New(8, 8)

	for _, opt := range opts {
		if err := opt(&b); err != nil {
			panic(err)
		}
	}

	return b
}

func NewEmptyBoard(opts ...BoardOption) board.Board {
	var b board.Board

	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			b.Squares = append(b.Squares, move.Position{File: f, Rank: r})
		}
	}

	b.Pieces = make(map[move.Position]*piece.Piece)

	for _, opt := range opts {
		if err := opt(&b); err != nil {
			panic(err)
		}
	}

	return b
}

func BoardWithPiece(p *piece.Piece) BoardOption {
	return func(b *board.Board) error {
		b.Pieces[p.Position] = p
		return nil
	}
}

func BoardWithPieces(ps []*piece.Piece) BoardOption {
	return func(b *board.Board) error {
		for _, p := range ps {
			b.Pieces[p.Position] = p
		}

		return nil
	}
}

func BoardWithDeletedPiece(pos move.Position) BoardOption {
	return func(b *board.Board) error {
		delete(b.Pieces, pos)
		return nil
	}
}

func BoardWithHistory(history []board.Turn) BoardOption {
	return func(b *board.Board) error {
		b.History = history
		return nil
	}
}

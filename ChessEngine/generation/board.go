package generation

import (
	"math/rand"

	"github.com/tomwatson6/chessbot/cmd/config"
	"github.com/tomwatson6/chessbot/internal/board"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/pkg/utility/linq"
	"github.com/tomwatson6/chessbot/testing/payloads"
)

func NewBoard(numPieces int) board.Board {
	if numPieces > 32 {
		numPieces = 32
	}

	b := payloads.NewEmptyBoard(
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.White,
			Position:     move.Position{File: 4, Rank: 0},
			PieceDetails: piece.NewKing(),
		}),
		payloads.BoardWithPiece(&piece.Piece{
			Colour:       colour.Black,
			Position:     move.Position{File: 4, Rank: 7},
			PieceDetails: piece.NewKing(),
		}),
	)

	pieces := config.GetStandardPieces()
	pieces = linq.Where(pieces, func(p *piece.Piece) bool {
		return p.GetPieceType() != piece.PieceTypeKing
	})
	pieces = shuffle(pieces)

	for len(pieces) > 32-numPieces {
		r := rand.Intn(8)
		c := rand.Intn(8)

		pos := move.Position{File: c, Rank: r}

		if _, ok := b.Pieces[pos]; !ok {
			p := pieces[len(pieces)-1]
			p.Position = pos
			b.Pieces[pos] = p
			pieces = pieces[:len(pieces)-1]
		}
	}

	return b
}

func shuffle[T any](arr []T) []T {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

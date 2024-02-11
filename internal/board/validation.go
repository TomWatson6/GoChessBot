package board

import (
	"errors"
	"fmt"

	"github.com/tomwatson6/chessbot/internal/board/rules"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

// ErrorInvalidPieceType is returned when a piece is not of a valid type.
var ErrorInvalidPieceType = errors.New("invalid piece type")

var (
	// ErrorInvalidPawnMove is an error returned for an invalid pawn move
	ErrorInvalidPawnMove = errors.New("invalid move specified for a pawn")
	// ErrorInvalidKnightMove is an error returned for an invalid knight move
	ErrorInvalidKnightMove = errors.New("invalid move specified for a knight")
	// ErrorInvalidBishopMove is an error returned for an invalid bishop move
	ErrorInvalidBishopMove = errors.New("invalid move specified for a bishop")
	// ErrorInvalidRookMove is an error returned for an invalid rook move
	ErrorInvalidRookMove = errors.New("invalid move specified for a rook")
	// ErrorInvalidQueenMove is an error returned for an invalid queen move
	ErrorInvalidQueenMove = errors.New("invalid move specified for a queen")
	// ErrorInvalidKingMove is an error returned for an invalid king move
	ErrorInvalidKingMove = errors.New("invalid move specified for a king")
)

func (b Board) ValidatePieceMove(p piece.Piece, m move.Move) error {
	switch p.PieceDetails.(type) {
	case piece.Pawn:
		return b.validatePawnMove(p, m)
	case piece.Rook:
		return b.validateRookMove(p, m)
	case piece.Knight:
		return b.validateKnightMove(p, m)
	case piece.Bishop:
		return b.validateBishopMove(p, m)
	case piece.Queen:
		return b.validateQueenMove(p, m)
	case piece.King:
		if m.Distance() == 1 {
			return b.validateKingMove(p, m)
		}

		return b.validateKingCastlingMove(p, m)
	default:
		return ErrorInvalidPieceType
	}
}

func (b Board) validatePawnMove(p piece.Piece, m move.Move) error {
	if err := p.IsValidMove(m); err != nil {
		return ErrorInvalidPawnMove
	}

	if m.To.File-m.From.File == 0 {
		// Handle pawn forward move
		validBoardMove := rules.Assert(
			rules.IsLineClear(b.Pieces, m),
			rules.IsNotPieceInEndPosition(b.Pieces, m.To),
		)

		if err := validBoardMove(); err != nil {
			return fmt.Errorf("%w: %s", ErrorInvalidPawnMove, err)
		}

		return nil
	}

	lastMove := b.History[len(b.History)-1]

	whiteMove := lastMove[colour.White]
	blackMove := lastMove[colour.Black]

	// Handle pawn capture
	validBoardMove := rules.Assert(
		rules.IsLineClear(b.Pieces, m),
		rules.IsValidIfPawnCapture(b.Pieces, whiteMove, blackMove, m),
	)

	if err := validBoardMove(); err != nil {
		return err
	}

	return nil
}

func (b Board) validateRookMove(p piece.Piece, m move.Move) error {
	if err := p.IsValidMove(m); err != nil {
		return ErrorInvalidRookMove
	}

	validBoardMove := rules.Assert(
		rules.IsLineClear(b.Pieces, m),
	)

	if err := validBoardMove(); err != nil {
		return err
	}

	return nil
}

func (b Board) validateKnightMove(p piece.Piece, m move.Move) error {
	if err := p.IsValidMove(m); err != nil {
		return ErrorInvalidKnightMove
	}

	return nil
}

func (b Board) validateBishopMove(p piece.Piece, m move.Move) error {
	// I think this is already done at a higher level, in IsCheck(...) <-- this needs to be checked TODO
	if err := p.IsValidMove(m); err != nil {
		return ErrorInvalidBishopMove
	}

	validBoardMove := rules.Assert(
		rules.IsLineClear(b.Pieces, m),
	)

	if err := validBoardMove(); err != nil {
		return err
	}

	return nil
}

func (b Board) validateQueenMove(p piece.Piece, m move.Move) error {
	if err := p.IsValidMove(m); err != nil {
		return ErrorInvalidQueenMove
	}

	validBoardMove := rules.Assert(
		rules.IsLineClear(b.Pieces, m),
	)

	if err := validBoardMove(); err != nil {
		return err
	}

	return nil
}

func (b Board) validateKingMove(p piece.Piece, m move.Move) error {
	if err := p.IsValidMove(m); err != nil {
		return ErrorInvalidKingMove
	}

	validBoardMove := rules.IsLineClear(b.Pieces, m)

	if err := validBoardMove(); err != nil {
		return err
	}

	return nil
}

func (b Board) validateKingCastlingMove(p piece.Piece, m move.Move) error {
	if err := p.IsValidMove(m); err != nil {
		return ErrorInvalidKingMove
	}

	validBoardMove := rules.IsValidIfCastlingMove(b.Pieces, m)

	if err := validBoardMove(); err != nil {
		return err
	}

	return nil
}

package rules

import "errors"

var (
	// ErrorIsOutOfBoundsOfBoard is thrown when the move is out of bounds
	ErrorIsOutOfBoundsOfBoard = errors.New("the move specified is out of the bounds of the board")
	// ErrorIsNotValidLine is thrown when the move specified is not for a valid line for a non-knight piece
	ErrorIsNotValidLine = errors.New("the move specified is not a valid line for a non-knight piece")
	// ErrorLineIsNotClear is thrown when the move specified is a move through a piece on the board
	ErrorLineIsNotClear = errors.New("the move specified is a move through a piece on the board")
	// ErrorIsFriendlyCapture is thrown when the move specified is attempting to take a piece of the same colour
	ErrorIsFriendlyCapture = errors.New("the move specified is a move that is attempting to take a piece of the same colour")
	// ErrorIsNotValidPawnCapture is thrown when the piece in the start position is a pawn, the pawn is moving forward vertically, and there is a piece in the end position
	ErrorIsNotValidPawnCapture = errors.New("the piece in the start position is a pawn, the pawn is moving forward vertically, and there is a piece in the end position")
	// ErrorIsNotValidDiagonalPawnCapture is thrown when the piece in the start position is a pawn and another piece of the opposite colour is not where it is diagonally moving
	ErrorIsNotValidDiagonalPawnCapture = errors.New("the piece in the start position is a pawn and another piece of the opposite colour is not where it is diagonally moving")
	// ErrorKingNotFound is thrown when there is no king for the colour specified in the pieces on the board
	ErrorKingNotFound = errors.New("there is no king for the colour specified in the pieces on the board")
	// ErrorIsPinned is thrown when the piece that is trying to be moved is pinned to their king
	ErrorIsPinned = errors.New("the piece that is trying to be moved is pinned the their king")
	// ErrorIsInCheck is thrown when the king is in check for the colour of the piece at the position provided
	ErrorIsInCheck = errors.New("the king is in check for the colour of the piece at the position provided")
	// ErrorPieceNotInStartPosition is thrown when there is no piece in the start position provided
	ErrorPieceNotInStartPosition = errors.New("there is no piece in the start position provided")
	// ErrorResultsInCheck is thrown when the move specified results in having the friendly king in check
	ErrorResultsInCheck = errors.New("the move specified results in having the friendly king in check")
	// ErrorInvalidCastlingMove is thrown when the move specified is not a valid castling move
	ErrorInvalidCastlingMove = errors.New("the move specified is not a valid castling move")
)

type Assertion func() error

// Assert chains all assertions together into one assertion
func Assert(assertions ...Assertion) Assertion {
	return func() error {
		for _, a := range assertions {
			if err := a(); err != nil {
				return err
			}
		}
		return nil
	}
}

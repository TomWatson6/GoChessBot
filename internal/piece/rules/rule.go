package rules

import "errors"

var (
	// ErrorIsNotValidLine is thrown when the move is not a valid line
	ErrorIsNotValidLine = errors.New("the move provided is not a valid line")
	// ErrorIsSmallerThanMinRange is thrown when the move provided is smaller than the min range
	ErrorIsSmallerThanMinRange = errors.New("the move provided is smaller than the min range")
	// ErrorDiagonalLineExceedsMaxRange is thrown when the move exceeds the max range and is diagonal
	ErrorDiagonalLineExceedsMaxRange = errors.New("the move provided exceeds the max range and is diagonal")
	// ErrorExceedsMaxRange is thrown when the move is larger than the max range of the piece
	ErrorExceedsMaxRange = errors.New("the move exceeds the max range of the piece")
	// ErrorIsNotValidKnightsMove is thrown when the move is not a valid knights move (2 forward, 1 to the side)
	ErrorIsNotValidKnightsMove = errors.New("the move is not a valid knights move")
	// ErrorIsNotCorrectDirection is thrown when the move is invalid due to the direction being incorrect
	ErrorIsNotCorrectDirection = errors.New("the move is invalid due to the direction being incorrect")
	// ErrorIsDiagonalLine is thrown when the line should not be diagonal, but is
	ErrorIsDiagonalLine = errors.New("the move is a diagonal line, but should not be")
	// ErrorIsNotDiagonalLine is thrown when the line is diagonal, and should not be
	ErrorIsNotDiagonalLine = errors.New("the move is not a diagonal line, but should be")
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

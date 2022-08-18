package rules

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

func IsValidLine(m move.Move) func() error {
	return func() error {
		x, y := splitMove(m)

		// If diagonal move, it is a valid line
		if x == y || x == -y {
			return nil
		}

		// If it is a horizontal or vertical move, it is a valid line
		if (x == 0 && y != 0) || (x != 0 && y == 0) {
			return nil
		}

		return ErrorIsNotValidLine
	}
}

func DoesNotExceedMaxRange(r int, m move.Move) func() error {
	return func() error {
		x, y := splitMove(m)

		// If diagonal move
		if x == y || x == -y {
			if x > 0 && r >= x {
				return nil
			}

			if y > 0 && r >= y {
				return nil
			}

			return ErrorExceedsMaxRange
		}

		// If vertical move
		if x == 0 {
			if r >= y {
				return nil
			}
		}

		// If horizontal move
		if y == 0 {
			if r >= x {
				return nil
			}
		}

		return ErrorExceedsMaxRange
	}
}

func DoesNotExceedMaxRangeIfDiagonal(r int, m move.Move) func() error {
	return func() error {
		x, y := splitMove(m)

		if x == y || x == -y {
			if x > 0 {
				if r < x {
					return ErrorDiagonalLineExceedsMaxRange
				}
			}

			if y > 0 {
				if r < y {
					return ErrorDiagonalLineExceedsMaxRange
				}
			}
		}

		return nil
	}
}

func IsValidKnightsMove(m move.Move) func() error {
	return func() error {
		x, y := splitMove(m)

		if x == 1 && y == 2 {
			return nil
		}

		if x == 2 && y == 1 {
			return nil
		}

		return ErrorIsNotValidKnightsMove
	}
}

func IsCorrectDirection(c colour.Colour, m move.Move) func() error {
	return func() error {
		y := m.To.Rank - m.From.Rank

		if c == colour.White && y > 0 {
			return nil
		}

		if c == colour.Black && y < 0 {
			return nil
		}

		return ErrorIsNotCorrectDirection
	}
}

func IsDiagonalLine(m move.Move) func() error {
	return func() error {
		x, y := splitMove(m)

		if x == y || x == -y {
			return nil
		}

		return ErrorIsNotDiagonalLine
	}
}

func IsNotDiagonalLine(m move.Move) func() error {
	return func() error {
		x, y := splitMove(m)

		if x == y || x == -y {
			return ErrorIsDiagonalLine
		}

		return nil
	}
}

func splitMove(m move.Move) (int, int) {
	x := m.To.File - m.From.File
	y := m.To.Rank - m.From.Rank

	return x, y
}

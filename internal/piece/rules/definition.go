package rules

import (
	"math"

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

func IsLargerThanOrEqualToMinRange(r int, m move.Move) func() error {
	return func() error {
		x0, y0 := splitMove(m)
		x := math.Abs(float64(x0))
		y := math.Abs(float64(y0))
		m := math.Max(x, y)
		// xf := math.Abs(float64(x))
		// yf := math.Abs(float64(y))
		rf := float64(r)

		if m >= rf {
			return nil
		}

		return ErrorIsSmallerThanMinRange
	}
}

func DoesNotExceedMaxRange(r int, m move.Move) func() error {
	return func() error {
		x, y := splitMove(m)
		xf := math.Abs(float64(x))
		yf := math.Abs(float64(y))
		rf := float64(r)

		// If diagonal move
		if xf == yf {
			if rf >= xf {
				return nil
			}

			if rf >= yf {
				return nil
			}

			return ErrorExceedsMaxRange
		}

		// If vertical move
		if xf == 0 {
			if rf >= yf {
				return nil
			}
		}

		// If horizontal move
		if yf == 0 {
			if rf >= xf {
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
		xf := math.Abs(float64(x))
		yf := math.Abs(float64(y))

		if xf == 1 && yf == 2 {
			return nil
		}

		if xf == 2 && yf == 1 {
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

package board

import (
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"math"
)

func (b Board) getLine(start, end move.Position, includingLast bool) []move.Position {
	// If x and y not equal to 0 (not horiz or vert), then if abs(x) != abs(y) (also not diagonal), return
	// Only side case is Knight, and you would never check a line for a Knight
	if end.File-start.File != 0 && end.Rank-start.Rank != 0 {
		if math.Abs(float64(end.File-start.File)) != math.Abs(float64(end.Rank-start.Rank)) {
			return []move.Position{}
		}
	}

	var line []move.Position

	//If line is diagonal
	if math.Abs(float64(end.File-start.File)) == math.Abs(float64(end.Rank-start.Rank)) {
		xStep := 1
		yStep := 1

		if start.File > end.File {
			xStep = -1
		}

		if start.Rank > end.Rank {
			yStep = -1
		}

		x := start.File
		y := start.Rank

		line = append(line, move.Position{File: x, Rank: y})

		for x != end.File || y != end.Rank {
			// if x == end.File && y == end.Rank {
			// 	break
			// }

			x += xStep
			y += yStep
			line = append(line, move.Position{File: x, Rank: y})
		}
	} else {
		//If line is vertical
		if start.File == end.File {
			if start.Rank < end.Rank {
				for i := start.Rank; i <= end.Rank; i++ {
					line = append(line, move.Position{File: start.File, Rank: i})
				}
			} else {
				for i := start.Rank; i >= end.Rank; i-- {
					line = append(line, move.Position{File: start.File, Rank: i})
				}
			}
		} else if start.Rank == end.Rank {
			//If line is horizontal
			if start.File < end.File {
				for i := start.File; i <= end.File; i++ {
					line = append(line, move.Position{File: i, Rank: start.Rank})
				}
			} else {
				for i := start.File; i >= end.File; i-- {
					line = append(line, move.Position{File: i, Rank: start.Rank})
				}
			}
		}
	}

	if includingLast {
		return line[1:]
	} else {
		return line[1 : len(line)-1]
	}
}

func (b Board) isLineClear(line []move.Position) bool {
	for _, pos := range line {
		if _, ok := b.Pieces[pos]; ok {
			return false
		}
	}

	return true
}

func (b Board) getRemainingPieces() []piece.Piece {
	var pieces []piece.Piece

	for _, p := range b.Pieces {
		pieces = append(pieces, p)
	}

	return pieces
}

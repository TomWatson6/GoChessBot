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

func (b Board) IsValidMove(m move.Move) bool {
	p := b.GetPiece(m.From)

	if valid := p.IsValidMove(m.To); valid {
		if _, ok := p.(piece.Knight); ok {
			return true
		}

	}

	return false
}

func (b Board) GetLine(start, end move.Position) []move.Position {
	//If line is diagonal
	if start.File != end.File && start.Rank != end.Rank {
		xStep := 1
		yStep := 1

		if start.File > end.File {
			xStep = -1
		}

		if start.Rank > end.Rank {
			yStep = -1
		}

		var line []move.Position
		x := start.File
		y := start.Rank

		for {
			if start.File == end.File && start.Rank == end.Rank {
				break
			}

			line = append(line, move.Position{File: x, Rank: y})
			x += xStep
			y += yStep
		}

		return line
	} else {
		//If line is vertical
		if start.File == end.File {
			var line []move.Position
			if start.Rank < end.Rank {
				for i := start.Rank; i <= end.Rank; i++ {
					line = append(line, move.Position{File: start.File, Rank: i})
				}
			} else {
				for i := start.Rank; i >= end.Rank; i-- {
					line = append(line, move.Position{File: start.File, Rank: i})
				}
			}

			return line
		} else if start.Rank == end.Rank {
			//If line is horizontal
			var line []move.Position
			if start.File < end.File {
				for i := start.File; i <= end.File; i++ {
					line = append(line, move.Position{File: i, Rank: start.Rank})
				}
			} else {
				for i := start.File; i >= end.File; i-- {
					line = append(line, move.Position{File: i, Rank: start.Rank})
				}
			}

			return line
		}
	}

	return []move.Position{}
}

func (b Board) IsLineClear(m move.Move) bool {
	line := b.GetLine(m.From, m.To)

	for _, pos := range line {
		if _, ok := b.Pieces[pos]; ok {
			return false
		}
	}

	return true
}

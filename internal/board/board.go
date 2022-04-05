package board

import (
	"fmt"
	"math"

	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

type Board struct {
	Pieces map[move.Position]piece.Piece
}

func (b Board) GetPiece(pos move.Position) piece.Piece {
	return b.Pieces[pos]
}

func (b Board) MovePiece(m move.Move) error {
	if b.IsValidMove(m) {
		b.Pieces[m.From] = b.Pieces[m.From].SetPosition(m.To)
		b.Pieces[m.To] = b.Pieces[m.From]
		delete(b.Pieces, m.From)
		return nil
	}

	return fmt.Errorf("invalid move: %v", m)
}

func (b Board) IsValidMove(m move.Move) bool {
	if m.From == m.To {
		return false
	}

	p := b.GetPiece(m.From)

	if valid := p.IsValidMove(m.To); valid {
		if _, ok := p.(piece.Knight); ok {
			return true
		}

		//Get line apart from last position
		//If line is clear, check that last square has no piece, or piece of opposite colour
		line := b.GetLine(m.From, m.To)

		if b.IsLineClear(line) {
			if opp := b.Pieces[m.To]; opp != nil {
				return opp.GetColour() != p.GetColour()
			}

			return true
		}
	}

	return false
}

func (b Board) GetLine(start, end move.Position) []move.Position {
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

		var line []move.Position
		x := start.File
		y := start.Rank

		for {
			if x == end.File && y == end.Rank {
				break
			}

			line = append(line, move.Position{File: x, Rank: y})
			x += xStep
			y += yStep
		}

		return line[1:]
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

			return line[1 : len(line)-1]
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

			return line[1 : len(line)-1]
		}
	}

	return []move.Position{}
}

func (b Board) IsLineClear(line []move.Position) bool {
	for _, pos := range line {
		if p := b.Pieces[pos]; p != nil {
			return false
		}
	}

	return true
}

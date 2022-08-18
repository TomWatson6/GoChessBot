package rules

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func InBoundsOfBoard(w, h int, m move.Move) func() error {
	return func() error {
		if m.From.File < 0 || m.From.File >= w {
			return ErrorIsOutOfBoundsOfBoard
		}

		if m.To.File < 0 || m.To.File >= w {
			return ErrorIsOutOfBoundsOfBoard
		}

		if m.From.Rank < 0 || m.From.Rank >= h {
			return ErrorIsOutOfBoundsOfBoard
		}

		if m.To.Rank < 0 || m.To.Rank >= h {
			return ErrorIsOutOfBoundsOfBoard
		}

		return nil
	}
}

func IsLineClear(ps map[move.Position]*piece.Piece, m move.Move) func() error {
	return func() error {
		line, err := getLine(m)
		if err != nil {
			return err
		}

		for _, pos := range line {
			if _, ok := ps[pos]; ok {
				return ErrorLineIsNotClear
			}
		}

		return nil
	}
}

func IsNotFriendlyCapture(ps map[move.Position]*piece.Piece, m move.Move) func() error {
	return func() error {
		p := ps[m.From]

		if p2, ok := ps[m.To]; ok {
			if p.Colour == p2.Colour {
				return ErrorIsFriendlyCapture
			}
		}

		return nil
	}
}

func IsNotPinned(w, h int, ps map[move.Position]*piece.Piece, m move.Move) func() error {
	return func() error {
		p := ps[m.From]
		k, err := getKing(ps, p.Colour)
		if err != nil {
			return err
		}

		// If the piece being moved is the king, then it is not pinned
		if k.Position == p.Position {
			return nil
		}

		expected := []*piece.Piece{
			k,
			p,
		}

		dx := p.Position.File - k.Position.File
		dy := p.Position.Rank - k.Position.Rank

		sx, sy := getSteps(dx, dy)

		index := 0

		x := k.Position.File
		y := k.Position.Rank

		// While still in bounds of the board
		for x >= 0 && x < w && y >= 0 && y < h {
			dest := move.Position{File: x, Rank: y}

			if p2, ok := ps[dest]; ok {
				// If destination is hit, then piece still blocks against potential attacker
				if dest == m.To {
					return nil
				}

				if index < len(expected) {
					if !p2.Equals(*expected[index]) {
						return nil
					}
					index++
					continue
				}

				if p.Colour != p2.Colour {
					if p2.IsValidMove(move.Move{From: p2.Position, To: k.Position}) {
						return ErrorIsPinned
					}
				}
			}

			x += sx
			y += sy
		}

		return nil
	}
}

func IsPawnCapture(ps map[move.Position]*piece.Piece, whiteMove, blackMove, m move.Move) func() error {
	return func() error {
		dx := m.To.File - m.From.File
		dy := m.To.Rank - m.From.Rank

		if dx != 1 || dy != 1 {
			return ErrorIsNotPawnCapture
		}

		p := ps[m.From]

		if p2, ok := ps[m.To]; ok {
			if p.Colour != p2.Colour {
				return nil
			}
		}

		// En passant
		if p2, ok := ps[move.Position{File: m.From.File + dx, Rank: m.From.Rank}]; ok {
			if p2.GetPieceType() != piece.PieceTypePawn {
				return ErrorIsNotPawnCapture
			}

			lastMove := blackMove
			if p2.Colour == colour.White {
				lastMove = whiteMove
			}

			dy := lastMove.To.Rank - lastMove.From.Rank

			// If the last move was the attacked pawns move and that the move was a 2 space move
			if lastMove.To == p2.Position && (dy == 2 || dy == -2) {
				if p.Colour != p2.Colour {
					return nil
				}
			}
		}

		return ErrorIsNotPawnCapture
	}
}

func IsPieceInStartPosition(ps map[move.Position]*piece.Piece, pos move.Position) func() error {
	return func() error {
		if _, ok := ps[pos]; ok {
			return nil
		}

		return ErrorPieceNotInStartPosition
	}
}

func IsNotInCheck(ps map[move.Position]*piece.Piece, pos move.Position) func() error {
	return func() error {
		p := ps[pos]

		col := p.Colour

		k, err := getKing(ps, col)
		if err != nil {
			return err
		}

		for _, p2 := range ps {
			if p2.IsValidMove(move.Move{From: p2.Position, To: k.Position}) {
				return ErrorIsInCheck
			}
		}

		return nil
	}
}

func getKing(ps map[move.Position]*piece.Piece, c colour.Colour) (*piece.Piece, error) {
	for _, k := range ps {
		if k.GetPieceType() == piece.PieceTypeKing && k.Colour == c {
			return k, nil
		}
	}

	return nil, ErrorKingNotFound
}

// getLine returns line from start up to and not including the destination
func getLine(m move.Move) ([]move.Position, error) {
	dx := m.To.File - m.From.File
	dy := m.To.Rank - m.From.Rank

	if dx != dy && dx != 0 && dy != 0 {
		return []move.Position{}, ErrorIsNotValidLine
	}

	sx, sy := getSteps(dx, dy)

	var line []move.Position

	x := m.From.File
	y := m.From.Rank

	for x != m.To.File && y != m.To.Rank {
		line = append(line, move.Position{File: x, Rank: y})

		x += sx
		y += sy
	}

	return line, nil
}

func getSteps(dx, dy int) (int, int) {
	sx := 1
	sy := 1

	if dx < 0 {
		sx = -1
	}

	if dx == 0 {
		sx = 0
	}

	if dy < 0 {
		sy = -1
	}

	if dy == 0 {
		sy = 0
	}

	return sx, sy
}

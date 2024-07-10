package rules

import (
	"math"
	"reflect"

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
			if pos == m.From {
				continue
			}
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

		if k.Position == p.Position {
			return nil
		}

		dx := p.Position.File - k.Position.File
		dy := p.Position.Rank - k.Position.Rank

		sx, sy := getSteps(dx, dy)

		var seq []*piece.Piece
		seq = append(seq, k)

		x := k.Position.File
		y := k.Position.Rank

		var line []move.Position
		line = append(line, move.Position{File: x, Rank: y})

		for x+sx >= 0 && x+sx < w && y+sy >= 0 && y+sy < h {
			x += sx
			y += sy

			line = append(line, move.Position{File: x, Rank: y})

			if p2, ok := ps[move.Position{File: x, Rank: y}]; ok {
				seq = append(seq, p2)

				if len(seq) == 2 && !reflect.DeepEqual(p, p2) {
					return nil
				}

				if len(seq) == 3 && p2.Colour != p.Colour {
					if err := p2.IsValidMove(move.Move{From: p2.Position, To: k.Position}); err == nil {
						// Check to see if destination of the move is on the line between the attacker and the king
						for _, l := range line {
							if reflect.DeepEqual(m.To, l) {
								return nil
							}
						}

						return ErrorIsPinned
					}

					return nil
				}
			}
		}

		return nil
	}
}

func IsValidIfPawnCapture(ps map[move.Position]*piece.Piece, whiteMove, blackMove, m *move.Move) func() error {
	return func() error {
		dx := m.To.File - m.From.File

		// If pawn is not moving diagonally then it isn't a pawn capture move
		if dx == 0 {
			return nil
		}

		p := ps[m.From]

		if p2, ok := ps[m.To]; ok {
			if p.Colour != p2.Colour {
				return nil
			}
		}

		if whiteMove == nil && blackMove == nil {
			return ErrorIsNotValidDiagonalPawnCapture
		}

		// En passant
		if p2, ok := ps[move.Position{File: m.From.File + dx, Rank: m.From.Rank}]; ok {
			if p2.GetPieceType() != piece.PieceTypePawn {
				return ErrorIsNotValidDiagonalPawnCapture
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

		return ErrorIsNotValidDiagonalPawnCapture
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

func IsNotPieceInEndPosition(ps map[move.Position]*piece.Piece, pos move.Position) func() error {
	return func() error {
		if _, ok := ps[pos]; ok {
			return ErrorIsNotValidPawnCapture
		}

		return nil
	}
}

func IsNotInCheck(ps map[move.Position]*piece.Piece, movingPiece *piece.Piece, m move.Move) func() error {
	return func() error {
		col := movingPiece.Colour

		k, err := getKing(ps, col)
		if err != nil {
			return err
		}

		attackedPosition := k.Position

		if movingPiece.GetPieceType() == piece.PieceTypeKing {
			attackedPosition = m.To
		}

		for _, p2 := range ps {
			if p2.Colour == col {
				continue
			}

			if err := p2.IsValidMove(move.Move{From: p2.Position, To: attackedPosition}); err == nil {
				return ErrorIsInCheck
			}
		}

		return nil
	}
}

func IsValidIfCastlingMove(width, height int, ps map[move.Position]*piece.Piece, m move.Move) func() error {
	return func() error {
		dx := m.To.File - m.From.File
		dy := m.To.Rank - m.From.Rank

		// If it is not a castling move, then it is a valid move
		// Once established that it's a 2 space move, we know it's a castling attempt
		if math.Abs(float64(dx)) < 2 && math.Abs(float64(dy)) < 2 {
			return nil
		}

		if dy != 0 {
			return ErrorInvalidCastlingMove
		}

		k := ps[m.From]
		if k.HasMoved() {
			return ErrorInvalidCastlingMove
		}

		sx, sy := getSteps(dx, dy)
		var line []move.Position

		x := m.From.File + sx
		y := m.From.Rank + sy

		rookFound := false

		for x >= 0 && x < width && y >= 0 && y < height {
			line = append(line, move.Position{File: x, Rank: y})

			if p, ok := ps[move.Position{File: x, Rank: y}]; ok {
				if p.GetPieceType() == piece.PieceTypeRook && !p.HasMoved() {
					rookFound = true
					break
				} else {
					return ErrorInvalidCastlingMove
				}
			}

			x += sx
			y += sy
		}

		if !rookFound {
			return ErrorInvalidCastlingMove
		}

		for _, pos := range line[:len(line)-1] {
			if isThreatened(ps, k, pos) {
				return ErrorInvalidCastlingMove
			}

			// Only need to check if king is castling through check, not the rook
			if pos == m.To {
				break
			}
		}

		return nil
	}
}

func IsNotMovingIntoDanger(ps map[move.Position]*piece.Piece, m move.Move) func() error {
	return func() error {
		p := ps[m.From]

		if isThreatened(ps, p, m.To) {
			return ErrorIsMovingIntoDanger
		}

		return nil
	}
}

func isThreatened(ps map[move.Position]*piece.Piece, p *piece.Piece, pos move.Position) bool {
	for _, pi := range ps {
		if pi.Colour == p.Colour {
			continue
		}

		attack := move.Move{From: pi.Position, To: pos}
		if err := p.IsValidMove(attack); err == nil {
			if p.GetPieceType() == piece.PieceTypeKnight {
				return true
			}

			if err := IsLineClear(ps, attack)(); err == nil {
				return true
			}
		}
	}

	return false
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

	if math.Abs(float64(dx)) != math.Abs(float64(dy)) && dx != 0 && dy != 0 {
		return []move.Position{}, ErrorIsNotValidLine
	}

	sx, sy := getSteps(dx, dy)

	var line []move.Position

	x := m.From.File
	y := m.From.Rank

	for x != m.To.File || y != m.To.Rank {
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

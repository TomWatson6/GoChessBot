package chess

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

func Contains(runes []rune, char rune) bool {
	for _, r := range runes {
		if r == char {
			return true
		}
	}

	return false
}

func (c Chess) TranslateNotation(n string) ([]move.Move, error) {
	var ms []move.Move
	var m move.Move

	runes := []rune(n)

	// If string contains 'x', it's a capture
	if Contains(runes, 'x') {
		// Split string into two parts
		parts := strings.Split(n, "x")
		left := []rune(parts[0])
		right := []rune(parts[1])
		m.To = move.Position{File: FileToNumber(right[0]), Rank: RankToNumber(right[1])}
		attackingPieces := c.Board.GetAttackingPieces(m.To)
		instantiated := false

		// Different variations: Nxe3, e2xf3, e3Nxf5, etc.
		if len(left) == 1 {
			// Piece other than a pawn is the capturing piece and is non-ambiguous
			for _, p := range attackingPieces {
				if p.GetLetter() == piece.PieceLetter(left[0]) {
					m.From = p.GetPosition()
					instantiated = true
					break
				}
			}
		} else if len(left) == 2 || len(left) == 3 {
			// Pawn is the capturing piece or piece other than a pawn is the capturing piece and is ambiguous
			m.From = move.Position{File: FileToNumber(left[0]), Rank: RankToNumber(left[1])}
			instantiated = true
		}

		if instantiated {
			ms = append(ms, m)
		} else {
			return nil, fmt.Errorf("invalid move: %v", n)
		}
	} else {
		// If first character is lower case, then it's a pawn move
		if unicode.IsLower(runes[0]) {
			m.To = move.Position{File: FileToNumber(runes[0]), Rank: RankToNumber(runes[1])}
			attackingPieces := c.Board.GetAttackingPieces(m.To)
			instantiated := false

			for _, p := range attackingPieces {
				if p.GetPieceType() == piece.PieceTypePawn && p.GetColour() == c.Turn {
					m.From = p.GetPosition()
					instantiated = true
					break
				}
			}

			if instantiated {
				ms = append(ms, m)
				return ms, nil
			} else {
				return ms, fmt.Errorf("invalid move: %s", n)
			}
		} else if runes[0] == 'O' {
			// If first character is 'O', it's a castling move (O-O)
			// If length of notation is 3, then it's a king side castling move
			if len(runes) == 3 {
				if c.Turn == colour.White {
					m.From = move.Position{File: 4, Rank: 0}
					m.To = move.Position{File: 6, Rank: 0}
					ms = append(ms, m)
					m.From = move.Position{File: 7, Rank: 0}
					m.To = move.Position{File: 5, Rank: 0}
					ms = append(ms, m)
				} else {
					m.From = move.Position{File: 4, Rank: 7}
					m.To = move.Position{File: 6, Rank: 7}
					ms = append(ms, m)
					m.From = move.Position{File: 7, Rank: 7}
					m.To = move.Position{File: 5, Rank: 7}
					ms = append(ms, m)
				}
			} else {
				// If length of notation is 5, then it's a queen side castling move (O-O-O)
				if c.Turn == colour.White {
					m.From = move.Position{File: 4, Rank: 0}
					m.To = move.Position{File: 2, Rank: 0}
					ms = append(ms, m)
					m.From = move.Position{File: 0, Rank: 0}
					m.To = move.Position{File: 3, Rank: 0}
					ms = append(ms, m)
				} else {
					m.From = move.Position{File: 4, Rank: 7}
					m.To = move.Position{File: 2, Rank: 7}
					ms = append(ms, m)
					m.From = move.Position{File: 0, Rank: 7}
					m.To = move.Position{File: 3, Rank: 7}
					ms = append(ms, m)
				}
			}
			return ms, nil
		} else if unicode.IsUpper(runes[0]) {
			// If first character is upper case, then it's a piece move
			m.To = move.Position{File: FileToNumber(runes[1]), Rank: RankToNumber(runes[2])}
			attackingPieces := c.Board.GetAttackingPieces(m.To)
			instantiated := false

			for _, p := range attackingPieces {
				if p.GetLetter() == piece.PieceLetter(runes[0]) && p.GetColour() == c.Turn {
					m.From = p.GetPosition()
					instantiated = true
					break
				}
			}

			if instantiated {
				ms = append(ms, m)
				return ms, nil
			} else {
				return ms, fmt.Errorf("invalid move: %s", n)
			}
		}
	}

	return ms, fmt.Errorf("invalid move: %s", n)
}

func FileToNumber(file rune) int {
	return int(file - 'a')
}

func RankToNumber(rank rune) int {
	return int(rank - '1')
}

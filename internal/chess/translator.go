package chess

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

// function for handling pawn standard move e.g. e4
func (c Chess) translatePawnMove(n string) (move.Move, error) {
	var m move.Move

	runes := []rune(n)

	m.To = move.Position{File: fileToNumber(runes[0]), Rank: rankToNumber(runes[1])}
	movePieces := c.Board.MoveMap[m.To]
	instantiated := false

	for _, movePiece := range movePieces {
		if movePiece.GetPieceType() == piece.PieceTypePawn {
			m.From = movePiece.Position
			instantiated = true
			break
		}
	}

	if instantiated {
		return m, nil
	}

	return move.Move{}, fmt.Errorf("invalid move: %s", n)
}

// function for handling piece standard move e.g. Qa3
func (c Chess) translatePieceMove(n string) (move.Move, error) {
	var m move.Move

	runes := []rune(n)

	m.To = move.Position{File: fileToNumber(runes[1]), Rank: rankToNumber(runes[2])}
	movePieces := c.Board.MoveMap[m.To]
	instantiated := false

	for _, movePiece := range movePieces {
		if movePiece.GetPieceLetter() == piece.PieceLetter(runes[0]) {
			m.From = movePiece.Position
			instantiated = true
			break
		}
	}

	if instantiated {
		return m, nil
	}

	return move.Move{}, fmt.Errorf("invalid move: %s", n)
}

func (c Chess) translateAmbiguousPieceMove(n string) move.Move {
	var m move.Move

	runes := []rune(n)

	m.From = move.Position{File: fileToNumber(runes[0]), Rank: rankToNumber(runes[1])}
	m.To = move.Position{File: fileToNumber(runes[3]), Rank: rankToNumber(runes[4])}

	return m
}

// function for handling captures e.g. Nxe3, e2xf3, e3Nxf5
func (c Chess) translatePawnCapture(n string) (move.Move, error) {
	var m move.Move

	parts := strings.Split(n, "x")
	left := []rune(parts[0])
	right := []rune(parts[1])

	m.To = move.Position{File: fileToNumber(right[0]), Rank: rankToNumber(right[1])}

	attacked := c.Board.Pieces[m.To]

	opp := colour.White

	if attacked.Colour == colour.White {
		opp = colour.Black
	}

	attackingPieces := c.Board.GetAttackingPiecesForColour(m.To, opp)

	for _, attackingPiece := range attackingPieces {
		if attackingPiece.Position.File == fileToNumber(left[0]) &&
			attackingPiece.GetPieceType() == piece.PieceTypePawn {
			m.From = attackingPiece.Position
			return m, nil
		}
	}

	return m, fmt.Errorf("invalid move: %s", n)
}

func (c Chess) translatePieceCapture(n string) (move.Move, error) {
	var m move.Move

	parts := strings.Split(n, "x")
	left := []rune(parts[0])
	right := []rune(parts[1])
	m.To = move.Position{File: fileToNumber(right[0]), Rank: rankToNumber(right[1])}
	attackingPieces := c.Board.ThreatMap[m.To]
	instantiated := false

	for _, p := range attackingPieces {
		if p.GetPieceLetter() == piece.PieceLetter(left[0]) {
			m.From = p.Position
			instantiated = true
			break
		}
	}

	if instantiated {
		return m, nil
	}

	return move.Move{}, fmt.Errorf("invalid move: %s", n)
}

func (c Chess) translateAmbiguousPieceCapture(n string) move.Move {
	var m move.Move

	parts := strings.Split(n, "x")
	left := []rune(parts[0])
	right := []rune(parts[1])

	m.From = move.Position{File: fileToNumber(left[0]), Rank: rankToNumber(left[1])}
	m.To = move.Position{File: fileToNumber(right[0]), Rank: rankToNumber(right[1])}

	return m
}

// function for handling castling e.g. O-O, O-O-O
func (c Chess) translateCastlingMove(n string) ([]move.Move, error) {
	var ms []move.Move
	var m move.Move

	runes := []rune(n)

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
	} else if len(runes) == 5 {
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

	if len(ms) == 0 {
		return ms, fmt.Errorf("invalid move: %s", n)
	}

	return ms, nil
}

// function for handling pawn promotion e.g. e8=Q, dxe8=Q
func (c Chess) translatePawnPromotionMove(n string) (move.Move, error) {
	parts := strings.Split(n, "=")
	to := parts[0]

	if len(to) == 2 {
		m, err := c.translatePawnMove(to)
		if err != nil {
			return m, err
		}
		return m, nil
	} else {
		m, err := c.translatePawnCapture(to)
		if err != nil {
			return m, err
		}
		return m, nil
	}
}

func (c Chess) TranslateNotation(n string) ([]move.Move, error) {
	// TODO: Pawn promotion could be all of the following:
	// e8=Q, dxe8=Q

	var ms []move.Move

	if strings.Contains(n, "x") {
		// Piece capture e.g. Nxf3, e2xf3, e3Nxf5
		if len(n) == 4 {
			if unicode.IsUpper([]rune(n)[0]) {
				n, err := c.translatePieceCapture(n)
				if err != nil {
					return ms, err
				}
				ms = append(ms, n)
			} else {
				n, err := c.translatePawnCapture(n)
				if err != nil {
					return ms, err
				}
				ms = append(ms, n)
			}
			return ms, nil
		} else if len(n) == 6 {
			if strings.Contains(n, "=") {
				//Pawn capture into promotion
				m, err := c.translatePawnPromotionMove(n)
				if err != nil {
					return ms, err
				}
				ms = append(ms, m)
				return ms, nil
			} else {
				n := c.translateAmbiguousPieceCapture(n)
				ms = append(ms, n)
				return ms, nil
			}
		}
		return ms, fmt.Errorf("invalid move: %s", n)
	} else {
		if []rune(n)[0] == 'O' {
			return c.translateCastlingMove(n)
		} else if len(n) == 2 {
			// Pawn move e.g. e4
			m, err := c.translatePawnMove(n)
			if err != nil {
				return ms, err
			}
			return append(ms, m), nil
		} else if len(n) == 3 {
			// Piece move e.g. Nf3
			m, err := c.translatePieceMove(n)
			if err != nil {
				return ms, err
			}
			return append(ms, m), nil
		} else if len(n) == 4 {
			// Pawn promotion e.g. e8=Q
			m, err := c.translatePawnPromotionMove(n)
			if err != nil {
				return ms, err
			}
			ms = append(ms, m)
			return ms, nil
		} else if len(n) == 5 {
			m := c.translateAmbiguousPieceMove(n)
			ms = append(ms, m)
			return ms, nil
		}
	}
	return ms, fmt.Errorf("invalid move: %s", n)
}

func fileToNumber(file rune) int {
	return int(file - 'a')
}

func rankToNumber(rank rune) int {
	return int(rank - '1')
}

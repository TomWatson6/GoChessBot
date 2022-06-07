package board

import (
	"fmt"
	"math"
	"sync"

	"github.com/tomwatson6/chessbot/cmd/config"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

type Board struct {
	Squares   []move.Position
	Pieces    map[move.Position]piece.Piece
	MoveMap   map[move.Position][]piece.Piece
	ThreatMap map[move.Position][]piece.Piece
	EnPassant map[move.Position]colour.Colour
}

func New() Board {
	var b Board

	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			b.Squares = append(b.Squares, move.Position{File: f, Rank: r})
		}
	}

	b.Pieces = make(map[move.Position]piece.Piece)

	for _, p := range config.GetStandardPieces() {
		b.Pieces[p.Position] = p
	}

	b.EnPassant = make(map[move.Position]colour.Colour)

	b.Update()

	return b
}

// Needs updating to include colour parameter, pawns can only move in 1 direction
func (b *Board) Update() {
	b.GenerateMoveMap()
	b.GenerateThreatMap()
}

func (b Board) IsCheck(c colour.Colour) bool {
	if king, err := b.getKing(c); err == nil {
		return len(b.ThreatMap[king.Position]) > 0
	}

	return false
}

// TODO: Look into making this concurrent
func (b Board) IsCheckMate(c colour.Colour) bool {
	if king, err := b.getKing(c); err == nil {
		if b.IsCheck(c) {
			for pos := range king.ValidMoves {
				opp := colour.White
				if c == colour.White {
					opp = colour.Black
				}
				threat := b.GetAttackingPiecesForColour(pos, opp)
				if len(threat) == 0 {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (b Board) getKing(c colour.Colour) (piece.Piece, error) {
	for _, p := range b.Pieces {
		if p.Colour == c && p.GetPieceType() == piece.PieceTypeKing {
			return p, nil
		}
	}

	return piece.Piece{}, fmt.Errorf("cannot find king")
}

// Needs updating to include colour parameter, pawns can only move in 1 direction
func (b *Board) GenerateMoveMap() {
	b.MoveMap = make(map[move.Position][]piece.Piece)
	pieces := b.GetRemainingPieces()
	// wg := &sync.WaitGroup{}
	// mu := &sync.Mutex{}

	// wg.Add(len(b.Squares) * len(pieces))
	for _, pos := range b.Squares {
		for i := range pieces {
			if b.IsValidMove(move.Move{From: pieces[i].Position, To: pos}) {
				b.MoveMap[pos] = append(b.MoveMap[pos], pieces[i])
				p := pieces[i]
				p.ValidMoves[pos] = true
				b.Pieces[p.Position] = p
			}
			// go func(pos move.Position, pieces []piece.Piece, i int) {
			// 	if b.IsValidMove(move.Move{From: pieces[i].Position, To: pos}) {
			// 		mu.Lock()
			// 		b.MoveMap[pos] = append(b.MoveMap[pos], pieces[i])
			// 		p := pieces[i]
			// 		p.ValidMoves = append(p.ValidMoves, pos)
			// 		b.Pieces[p.Position] = p
			// 		mu.Unlock()
			// 	}
			// 	wg.Done()
			// }(pos, pieces, i)
		}
	}
	// wg.Wait()
}

// Needs updating to include colour parameter, pawns can only move in 1 direction
func (b *Board) GenerateThreatMap() {
	b.ThreatMap = make(map[move.Position][]piece.Piece)
	pieces := b.GetRemainingPieces()
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Add(len(b.Squares) * len(pieces))
	for _, pos := range b.Squares {
		for _, p := range pieces {
			go func(pos move.Position, p piece.Piece) {
				defer wg.Done()
				if b.IsValidMove(move.Move{From: p.Position, To: pos}) {
					if p.GetPieceType() == piece.PieceTypePawn {
						x := math.Abs(float64(p.Position.File - pos.File))
						y := math.Abs(float64(p.Position.Rank - pos.Rank))

						// Diagonal move means attacking move
						// Horizontal & Vertical moves aren't attacking moves
						if x == 0 && (y == 1 || y == 2) {
							return
						}
					}

					mu.Lock()
					b.ThreatMap[pos] = append(b.ThreatMap[pos], p)
					mu.Unlock()
				}
			}(pos, p)
		}
	}

	wg.Wait()
}

func (b Board) GetRemainingPieces() []piece.Piece {
	var pieces []piece.Piece

	for _, p := range b.Pieces {
		pieces = append(pieces, p)
	}

	return pieces
}

func (b Board) GetMoveMapForColour(pos move.Position, c colour.Colour) []piece.Piece {
	var pieces []piece.Piece

	for _, p := range b.MoveMap[pos] {
		if p.Colour == c {
			pieces = append(pieces, p)
		}
	}

	return pieces
}

func (b Board) GetAttackingPiecesForColour(pos move.Position, c colour.Colour) []piece.Piece {
	var pieces []piece.Piece

	for _, p := range b.ThreatMap[pos] {
		if p.Colour == c {
			pieces = append(pieces, p)
		}
	}

	return pieces
}

func (b *Board) MovePiece(m move.Move) error {
	p := b.Pieces[m.From]
	if _, ok := p.ValidMoves[m.To]; ok {
		if b2, err := b.isPinned(m); err == nil {
			*b = b2
			return nil
		} else {
			return fmt.Errorf("failed to move piece: %w", err)
		}
	}

	return fmt.Errorf("invalid move: %v", m)
}

// CHECK WHAT IS HAPPENING BELOW
// 1 wR wN wB wK ## wB wN wR
// 2 wP wP wP ## ## wP wP wP
// 3 ## ## ## ## ## ## ## ##
// 4 ## ## ## wQ ## ## ## ##
// 5 ## ## ## ## ## ## ## ##
// 6 ## bQ ## ## bP ## ## ##
// 7 bP bP bP ## ## bP bP bP
// 8 bR bN bB bK ## bB bN bR
//    H  G  F  E  D  C  B  A
// Black's move: Ke7
// [(4,7)->(4,6)]
// invalid move: (4,7)->(4,6), err: failed to move piece: {(4,7) Black map[(3,6):true (3,7):true (4,5):true (4,6):true (4,7):true (5,5):true] {}} is pinned
// 1 wR wN wB wK ## wB wN wR
// 2 wP wP wP ## ## wP wP wP
// 3 ## ## ## ## ## ## ## ##
// 4 ## ## ## wQ ## ## ## ##
// 5 ## ## ## ## ## ## ## ##
// 6 ## bQ ## ## bP ## ## ##
// 7 bP bP bP bK ## bP bP bP
// 8 bR bN bB bK ## bB bN bR
//    H  G  F  E  D  C  B  A

func (b Board) isPinned(m move.Move) (Board, error) {
	p := b.Pieces[m.From]

	p.Position = m.To
	b.Pieces[m.From] = p
	b.Pieces[m.To] = b.Pieces[m.From]
	delete(b.Pieces, m.From)

	b.Update()

	if !b.IsCheck(p.Colour) {
		if p.GetPieceType() == piece.PieceTypePawn {
			details := p.PieceDetails.(piece.Pawn)
			details.HasMoved = true
			p.PieceDetails = details
			b.Pieces[m.To] = p
		}

		return b, nil
	}

	p.Position = m.From
	b.Pieces[m.From] = p
	return Board{}, fmt.Errorf("%v is pinned", p)
}

func (b Board) IsValidMove(m move.Move) bool {
	if m.From == m.To {
		return false
	}

	p := b.Pieces[m.From]

	// TODO: Update to check for constraints of the board size (8x8)
	if p.IsValidMove(m) {
		if p.GetPieceType() == piece.PieceTypeKnight {
			if p2, ok := b.Pieces[m.To]; ok {
				return p2.Colour != p.Colour
			}

			return true
		}

		// if p.GetPieceType() == piece.PieceTypePawn && m.From.File == m.To.File {
		// 	m.To = move.Position{File: m.To.File, Rank: m.To.Rank + 1}
		// }

		//Get line apart from last position
		//If line is clear, check that last square has no piece, or piece of opposite colour
		// line := b.GetLine(m.From, m.To, false)
		line := b.GetLine(m.From, m.To, p.GetPieceType() == piece.PieceTypePawn && m.From.File == m.To.File)

		if b.IsLineClear(line) {
			if opp, ok := b.Pieces[m.To]; ok {
				return opp.Colour != p.Colour
			} else if m.From.File != m.To.File && p.GetPieceType() == piece.PieceTypePawn {
				return false
			}

			return true
		}
	}

	return false
}

func (b Board) GetLine(start, end move.Position, includingLast bool) []move.Position {
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

func (b Board) IsLineClear(line []move.Position) bool {
	for _, pos := range line {
		if _, ok := b.Pieces[pos]; ok {
			return false
		}
	}

	return true
}

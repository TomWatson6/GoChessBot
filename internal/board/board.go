package board

import (
	"fmt"
	"github.com/tomwatson6/chessbot/internal/rules"
	"math"
	"sync"

	"github.com/tomwatson6/chessbot/cmd/config"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

// Board is a struct to hold the current state of the chess game board
type Board struct {
	Squares    []move.Position
	Pieces     map[move.Position]piece.Piece
	MoveMap    map[move.Position][]piece.Piece
	ThreatMap  map[move.Position][]piece.Piece
	EnPassant  map[move.Position]colour.Colour
	MoveNumber int
}

// New makes a new instance of a board with a default state
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

	b.MoveNumber = 0
	b.updatePieceHistory()

	b.Update()

	return b
}

// MovePiece moves a piece on the board if possible, else it returns an error stating why it cannot
func (b *Board) MovePiece(m move.Move) error {
	p := b.Pieces[m.From]
	if p.ValidMoves[m.To] {
		if b2, err := b.isPinned(m); err == nil {
			*b = b2
			b.updatePieceHistory()

			return nil
		} else {
			return fmt.Errorf("failed to move piece: %w", err)
		}
	}

	return fmt.Errorf("invalid move: %v", m)
}

// IsValidMove checks to see that the move specified is valid given the current state of the board
func (b Board) IsValidMove(m move.Move) bool {
	if m.From == m.To {
		return false
	}

	p := b.Pieces[m.From]

	// TODO: Update to check for constraints of the board size (8x8)
	if p.IsValidMove(m) {
		// With the knight being the only piece that does not move in a straight line,
		// you need to check explicitly if there is a piece where it is trying to move to,
		// as this logic is normally done in the IsLineClear function
		if p.GetPieceType() == piece.PieceTypeKnight {
			if p2, ok := b.Pieces[m.To]; ok {
				return p2.Colour != p.Colour
			}

			return true
		}

		// Pawns cannot take when moving forward, so we must include p.GetPieceType() == piece.PieceTypePawn && not diagonal move for includingLast
		line := b.getLine(m.From, m.To, p.GetPieceType() == piece.PieceTypePawn && m.From.File == m.To.File)

		if b.isLineClear(line) {
			if opp, ok := b.Pieces[m.To]; ok {
				return opp.Colour != p.Colour
			} else if m.From.File != m.To.File && p.GetPieceType() == piece.PieceTypePawn {
				// Check for en passant
				if attacked, ok := b.Pieces[move.Position{File: m.To.File, Rank: m.From.Rank}]; ok {
					fileDiff := attacked.Position.File - attacked.History[b.MoveNumber-1].File
					rankDiff := attacked.Position.Rank - attacked.History[b.MoveNumber-1].Rank

					attackedRules := rules.Rules{
						rules.FirstMoveRule{},
						rules.MovedLastTurnRule{},
						rules.LastTurnPositionChangeRule{
							Diff: move.Position{
								File: int(math.Abs(float64(fileDiff))),
								Rank: int(math.Abs(float64(rankDiff))),
							},
						},
					}

					if attackedRules.All(attacked) {
						return true
					}
				}

				return false
			}

			return true
		}
	}

	return false
}

// updatePieceHistory updates all pieces to keep a track of their moves throughout the current game
func (b *Board) updatePieceHistory() {
	wg := &sync.WaitGroup{}
	wg.Add(len(b.Pieces))

	for _, p := range b.Pieces {
		go func(wg *sync.WaitGroup, p piece.Piece) {
			p.History[b.MoveNumber] = p.Position
			wg.Done()
		}(wg, p)
	}

	wg.Wait()
}

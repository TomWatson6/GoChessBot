package board

import (
	"fmt"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"math"
	"sync"
)

// Update refreshes the state of the board,
// and loads new possible moves into memory based on the state change of the board
func (b *Board) Update() {
	b.GenerateMoveMap()
	b.GenerateThreatMap()
}

// IsCheckMate checks for the state of the board being check mate for the colour provided
// TODO: Look into making this concurrent
func (b Board) IsCheckMate(c colour.Colour) bool {
	king, err := b.getKing(c)
	if err != nil {
		return false
	}

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

	return false
}

// GenerateMoveMap takes the current state of the board to generate all possible moves from pieces
func (b *Board) GenerateMoveMap() {
	b.MoveMap = make(map[move.Position][]piece.Piece)
	pieces := b.getRemainingPieces()

	for _, pos := range b.Squares {
		for _, piece := range pieces {
			if b.IsValidMove(move.Move{From: piece.Position, To: pos}) {
				b.MoveMap[pos] = append(b.MoveMap[pos], piece)
				p := piece
				p.ValidMoves[pos] = true
				b.Pieces[p.Position] = p
			}
		}
	}
}

// GenerateThreatMap takes the current state of the board to calculate all threats to each square on the board
func (b *Board) GenerateThreatMap() {
	b.ThreatMap = make(map[move.Position][]piece.Piece)
	pieces := b.getRemainingPieces()
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Add(len(b.Squares) * len(pieces))
	for _, pos := range b.Squares {
		for _, p := range pieces {
			go func(pos move.Position, p piece.Piece) {
				defer wg.Done()
				if b.IsValidMove(move.Move{From: p.Position, To: pos}) {
					if p.GetPieceType() == piece.PieceTypePawn {
						file := p.Position.File - pos.File
						rank := p.Position.Rank - pos.Rank

						x := math.Abs(float64(file))
						y := math.Abs(float64(rank))

						// TODO: Remove complexity by adding this to it's own method
						// TODO: Look into moving this into the b.IsValidMove(move) method in this file
						// TODO: Look into separating this large file into separate files based on functionality
						// Check for en passant

						// Make set of rules to loop through, and if returns true, then passes rules for en-passant move, and adds to the threat map
						// TODO: Rules are the following:
						// 	Pawn to the left or right - check both for potentially 2 additional moves
						// 	Pawn history having 2 unique moves (means that it's only moved once)
						// 	MoveA == MoveB where difference in y == 2
						// 	MoveA -> MoveB was the last move made (end of history)
						//pieceRules := rules.Rules{
						//	rules.PawnCaptureMoveRule{
						//		Dest: pos,
						//	},
						//}

						//adjRules := rules.Rules{
						//	rules.FirstMoveRule{},
						//	rules.MovedLastTurnRule{},
						//	rules.LastTurnPositionChangeRule{
						//		Diff: move.Position{
						//			File: 0,
						//			Rank: 2,
						//		},
						//	},
						//}

						// Check to see if the move being made is a pawn capture
						// TODO: This needs some looking into, make sure we don't remove unnecessary pawn moves

						// Check on left side of pawn
						//if adj, ok := b.Pieces[move.Position{File: file - 1, Rank: rank}]; ok {
						//	if !adjRules.All(adj) {
						//		return
						//	}
						//}

						// Check on right side of pawn
						//if adj, ok := b.Pieces[move.Position{File: file + 1, Rank: rank}]; ok {
						//	if !adjRules.All(adj) {
						//		return
						//	}
						//}

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

// GetMoveMapForColour gets all possible moves for the position and colour specified
func (b Board) GetMoveMapForColour(pos move.Position, c colour.Colour) []piece.Piece {
	var pieces []piece.Piece

	for _, p := range b.MoveMap[pos] {
		if p.Colour == c {
			pieces = append(pieces, p)
		}
	}

	return pieces
}

// GetAttackingPiecesForColour gets all the pieces that are threatening a square based on the colour provided
func (b Board) GetAttackingPiecesForColour(pos move.Position, c colour.Colour) []piece.Piece {
	var pieces []piece.Piece

	for _, p := range b.ThreatMap[pos] {
		if p.Colour == c {
			pieces = append(pieces, p)
		}
	}

	return pieces
}

// IsCheck checks to see whether the king for the colour provided is currently in check
func (b Board) IsCheck(c colour.Colour) bool {
	if king, err := b.getKing(c); err == nil {
		return len(b.ThreatMap[king.Position]) > 0
	}

	return false
}

// isPinned checks to see if the move provided is possible based on whether it is pinned to it's king or not
func (b Board) isPinned(m move.Move) (Board, error) {
	p := b.Pieces[m.From]

	attackedPiece, isAttacking := b.Pieces[m.To]

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

	if isAttacking {
		b.Pieces[m.To] = attackedPiece
	}

	p.Position = m.From
	b.Pieces[m.From] = p
	delete(b.Pieces, m.To)
	return Board{}, fmt.Errorf("cannot make move: %v, as you are putting the king in check", m)
}

// getKing gets the king piece for the colour provided
func (b Board) getKing(c colour.Colour) (piece.Piece, error) {
	for _, p := range b.Pieces {
		if p.Colour == c && p.GetPieceType() == piece.PieceTypeKing {
			return p, nil
		}
	}

	return piece.Piece{}, fmt.Errorf("cannot find king")
}

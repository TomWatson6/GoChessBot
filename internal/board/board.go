package board

import (
	"github.com/tomwatson6/chessbot/cmd/config"
	"github.com/tomwatson6/chessbot/internal/board/rules"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

// Board is a struct to hold the current state of the chess game board
type Board struct {
	Width   int
	Height  int
	Squares []move.Position
	Pieces  map[move.Position]*piece.Piece
	History []Turn
}

type Turn map[colour.Colour]move.Move

// New makes a new instance of a board with a default state
func New(w, h int) Board {
	var b Board

	b.Width = w
	b.Height = h

	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			b.Squares = append(b.Squares, move.Position{File: f, Rank: r})
		}
	}

	b.Pieces = make(map[move.Position]*piece.Piece)

	for _, p := range config.GetStandardPieces() {
		b.Pieces[p.Position] = p
	}

	b.History = append(b.History, make(Turn))

	return b
}

func (b Board) IsValidMove(m move.Move) error {
	// Firstly check the rules that will always need to be checked
	rs := rules.Assert(
		rules.InBoundsOfBoard(b.Width, b.Height, m),
		rules.IsPieceInStartPosition(b.Pieces, m.From),
		rules.IsNotPinned(b.Width, b.Height, b.Pieces, m),
		rules.IsNotFriendlyCapture(b.Pieces, m),
	)

	if err := rs(); err != nil {
		return err
	}

	p := b.Pieces[m.From]

	// If in check, then from/to must be on the line between attacker and king unless a knight in which case the from must be the king
	// First need to check if in check

	// Check for the king being in check, that will narrow the amount of moves to a subset if so

	// If in check, then a different ruleset must be applied

	if err := b.ValidatePieceMove(*p, m); err != nil {
		return err
	}

	return nil
}

func (b *Board) Move(m move.Move) error {
	if err := b.IsValidMove(m); err != nil {
		return err
	}

	p := b.Pieces[m.From]
	p.Position = m.To

	if _, ok := b.Pieces[m.To]; ok {
		b.Pieces[m.To] = p
		b.Pieces[m.From] = nil
	} else {
		// En passant
		dx := m.To.File - m.From.File

		b.Pieces[m.To] = p
		b.Pieces[move.Position{File: m.From.File + dx, Rank: m.From.Rank}] = nil
	}

	b.History[len(b.History)-1][p.Colour] = m

	// Create new entry if black's move is successful
	if p.Colour == colour.Black {
		b.History = append(b.History, make(Turn))
	}

	return nil
}

// Make rules for IsCheck and IsCheckMate, should be able to check them per successful turn, as an update on state of the board?
// Maybe make a map[colour.Colour]{{None|Check|CheckMate|StaleMate}}

// MovePiece moves a piece on the board if possible, else it returns an error stating why it cannot
//func (b *Board) MovePiece(m move.Move) error {
//	p := b.Pieces[m.From]
//	if p.ValidMoves[m.To] {
//		if b2, err := b.isPinned(m); err == nil {
//			*b = b2
//			b.updatePieceHistory()
//
//			return nil
//		} else {
//			return fmt.Errorf("failed to move piece: %w", err)
//		}
//	}
//
//	return fmt.Errorf("not in valid move slice for %+v", p)
//}

// IsValidMove checks to see that the move specified is valid given the current state of the board
//func (b Board) IsValidMove(m move.Move) bool {
//	if m.From == m.To {
//		return false
//	}
//
//	p := b.Pieces[m.From]
//
//	// TODO: Update to check for constraints of the board size (8x8)
//	if p.IsValidMove(m) {
//		// With the knight being the only piece that does not move in a straight line,
//		// you need to check explicitly if there is a piece where it is trying to move to,
//		// as this logic is normally done in the IsLineClear function
//		if p.GetPieceType() == piece.PieceTypeKnight {
//			if p2, ok := b.Pieces[m.To]; ok {
//				return p2.Colour != p.Colour
//			}
//
//			return true
//		}
//
//		// Pawns cannot take when moving forward, so we must include p.GetPieceType() == piece.PieceTypePawn && not diagonal move for includingLast
//		line := b.getLine(m.From, m.To, p.GetPieceType() == piece.PieceTypePawn && m.From.File == m.To.File)
//
//		if b.isLineClear(line) {
//			if opp, ok := b.Pieces[m.To]; ok {
//				return opp.Colour != p.Colour
//			} else if m.From.File != m.To.File && p.GetPieceType() == piece.PieceTypePawn {
//				// Check for en passant
//				if attacked, ok := b.Pieces[move.Position{File: m.To.File, Rank: m.From.Rank}]; ok {
//					fileDiff := attacked.Position.File - attacked.History[b.MoveNumber-1].File
//					rankDiff := attacked.Position.Rank - attacked.History[b.MoveNumber-1].Rank
//
//					attackedRules := rules.Rules{
//						rules.FirstMoveRule{},
//						rules.MovedLastTurnRule{},
//						rules.LastTurnPositionChangeRule{
//							Diff: move.Position{
//								File: int(math.Abs(float64(fileDiff))),
//								Rank: int(math.Abs(float64(rankDiff))),
//							},
//						},
//					}
//
//					if attackedRules.All(attacked) {
//						return true
//					}
//				}
//
//				return false
//			}
//
//			return true
//		}
//	}
//
//	return false
//}
//
//// updatePieceHistory updates all pieces to keep a track of their moves throughout the current game
//func (b *Board) updatePieceHistory() {
//	wg := &sync.WaitGroup{}
//	wg.Add(len(b.Pieces))
//
//	for _, p := range b.Pieces {
//		go func(wg *sync.WaitGroup, p piece.Piece) {
//			p.History[b.MoveNumber] = p.Position
//			wg.Done()
//		}(wg, p)
//	}
//
//	wg.Wait()
//}

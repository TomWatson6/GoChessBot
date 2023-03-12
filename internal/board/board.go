package board

import (
	"errors"

	"github.com/tomwatson6/chessbot/cmd/config"
	"github.com/tomwatson6/chessbot/internal/board/rules"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
)

var ErrorIsCheckMate = errors.New("king is in check mate, game over")

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

	for r := 0; r < b.Height; r++ {
		for f := 0; f < b.Width; f++ {
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

	if err := b.ValidatePieceMove(*p, m); err != nil {
		return err
	}

	return nil
}

// TODO: look at logic for this - en passant is currently being considered for any piece moving without taking
func (b *Board) Move(m move.Move) error {
	if err := b.IsValidMove(m); err != nil {
		return err
	}

	var toDelete move.Position

	p := b.Pieces[m.From]
	p.Position = m.To

	if p.GetPieceType() == piece.PieceTypePawn {
		p.PieceDetails = piece.NewPawn(
			piece.PawnWithColour(p.Colour),
			piece.PawnWithHasMoved(true),
		)

		if _, ok := b.Pieces[m.To]; ok {
			b.Pieces[m.To] = p
			delete(b.Pieces, m.From)
		} else {
			// En passant
			dx := m.To.File - m.From.File

			b.Pieces[m.To] = p
			toDelete = move.Position{File: m.From.File + dx, Rank: m.From.Rank}
		}
	} else if p.GetPieceType() == piece.PieceTypeKing {
		p.PieceDetails = piece.NewKing(
			piece.KingWithHasMoved(true),
		)

		// Castling move, so needs to also move the rook
		if m.Distance() == 2 {
			rookMove := b.getRookCastlingMove(m)
			r := b.Pieces[rookMove.From]

			r.Position = rookMove.To
			r.PieceDetails = piece.NewRook(
				piece.RookWithHasMoved(true),
			)

			b.Pieces[rookMove.To] = r

			delete(b.Pieces, rookMove.From)
		}

		b.Pieces[m.To] = p
		toDelete = m.From
	} else {
		b.Pieces[m.To] = p
		toDelete = m.From
	}

	delete(b.Pieces, toDelete)

	b.History[len(b.History)-1][p.Colour] = m

	// Create new entry if black's move is successful
	if p.Colour == colour.Black {
		b.History = append(b.History, make(Turn))
	}

	return nil
}

func (b Board) GetPiecesThatMoveToDestWithColour(dest move.Position, col colour.Colour) ([]*piece.Piece, error) {
	var output []*piece.Piece

	for _, p := range b.Pieces {
		if p.Colour != col {
			continue
		}

		if err := p.IsValidMove(move.Move{
			From: p.Position,
			To:   dest,
		}); err == nil {
			output = append(output, p)
		} else {
			return []*piece.Piece{}, err
		}
	}

	return output, nil
}

func (b Board) GetAttackingPiecesForColour(dest move.Position, col colour.Colour) ([]*piece.Piece, error) {
	var output []*piece.Piece

	for _, p := range b.Pieces {
		if p.Colour == col {
			continue
		}

		m := move.Move{
			From: p.Position,
			To:   dest,
		}

		if err := b.IsValidMove(m); err == nil {
			output = append(output, p)
		}
	}

	return output, nil
}

func (b Board) IsCheck(c colour.Colour) (*piece.Piece, bool, error) {
	k, err := b.getKing(c)
	if err != nil {
		return nil, false, err
	}

	ps, err := b.GetAttackingPiecesForColour(k.Position, c)
	if err != nil {
		return nil, false, err
	}

	for _, p := range ps {
		if err := b.IsValidMove(move.Move{From: p.Position, To: k.Position}); err != nil {
			continue
		}

		return p, true, nil
	}

	return nil, false, nil
}

func (b Board) IsCheckMate(c colour.Colour) (bool, error) {
	if p, check, err := b.IsCheck(c); check && err == nil {
		k, err := b.getKing(c)
		if err != nil {
			return false, err
		}

		if possible := b.kingCanMoveToSafety(k); possible {
			return false, nil
		}

		attackLine := b.getLine(p.Position, k.Position, true, false)

		if possible := b.checkIfPiecesCanMoveToLine(c, attackLine); possible {
			return false, nil
		}

		return true, nil

		// Can a piece take the attacking piece for the king, how will we find a piece that's attacking the king?
		// IsCheck could return the piece that is attacking the king maybe? In which case we can then check if any of the piece of the king's
		// colour are able to take it
		// Can anything block the piece, we also need to know the line of the attacking piece if it exists, which we can use getLine(...) from util.go to do so
		// Can any piece of the attacked king's colour move onto any of the line piece
		// If we include the starting position (where the attacking piece is) in the line, then we don't have to do the first check, we can just check if we can take it
		// May have to consider Pawn moves carefully, as forward moves are not a capture but IsValidMove(...) in board should cover this case already (I'm hoping...)
	}

	return false, nil
}

func (b Board) getRookCastlingMove(m move.Move) move.Move {
	switch {
	case m.To == move.Position{File: 2, Rank: 0}:
		return move.Move{
			From: move.Position{File: 0, Rank: 0},
			To:   move.Position{File: 3, Rank: 0},
		}
	case m.To == move.Position{File: 6, Rank: 0}:
		return move.Move{
			From: move.Position{File: 7, Rank: 0},
			To:   move.Position{File: 5, Rank: 0},
		}
	case m.To == move.Position{File: 2, Rank: 7}:
		return move.Move{
			From: move.Position{File: 0, Rank: 7},
			To:   move.Position{File: 3, Rank: 7},
		}
	case m.To == move.Position{File: 6, Rank: 7}:
		return move.Move{
			From: move.Position{File: 7, Rank: 7},
			To:   move.Position{File: 5, Rank: 7},
		}
	default:
		return move.Move{}
	}
}

func (b Board) kingCanMoveToSafety(k *piece.Piece) bool {
	iter := []int{-1, 0, 1}

	for _, f := range iter {
		for _, r := range iter {
			if f == 0 && r == 0 {
				continue
			}

			file := k.Position.File + f
			rank := k.Position.Rank + r

			move := move.Move{
				From: k.Position,
				To:   move.Position{File: file, Rank: rank},
			}

			if err := b.IsValidMove(move); err == nil {
				return true
			}
		}
	}

	return false
}

func (b Board) checkIfPiecesCanMoveToLine(c colour.Colour, line []move.Position) bool {
	ps := b.getRemainingPieces(c)

	for _, dest := range line {
		for _, p := range ps {
			if err := b.IsValidMove(move.Move{From: p.Position, To: dest}); err == nil {
				return true
			}
		}
	}

	return false
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

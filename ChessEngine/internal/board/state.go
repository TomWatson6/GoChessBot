package board

// Update refreshes the state of the board,
// and loads new possible moves into memory based on the state change of the board
// func (b *Board) Update() {
// 	b.GenerateMoveMap()
// 	b.GenerateThreatMap()
// }

// IsCheckMate checks for the state of the board being check mate for the colour provided
// TODO: Look into making this concurrent
// func (b Board) IsCheckMate(c colour.Colour) bool {
// 	king, err := b.getKing(c)
// 	if err != nil {
// 		return false
// 	}

// 	if b.IsCheck(c) {
// 		for pos := range king.ValidMoves {
// 			opp := colour.White
// 			if c == colour.White {
// 				opp = colour.Black
// 			}
// 			threat := b.GetAttackingPiecesForColour(pos, opp)
// 			if len(threat) == 0 {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	return false
// }

// GenerateMoveMap takes the current state of the board to generate all possible moves from pieces
// func (b *Board) GenerateMoveMap() {
// 	b.MoveMap = make(map[move.Position][]piece.Piece)
// 	pieces := b.getRemainingPieces()

// 	// Reset ValidMoves before assigning to it
// 	for pos, piece := range b.Pieces {
// 		piece.ValidMoves = make(map[move.Position]bool)
// 		b.Pieces[pos] = piece
// 	}

// 	for _, pos := range b.Squares {
// 		for _, piece := range pieces {
// 			if b.IsValidMove(move.Move{From: piece.Position, To: pos}) {
// 				b.MoveMap[pos] = append(b.MoveMap[pos], piece)
// 				p := piece
// 				p.ValidMoves[pos] = true
// 				b.Pieces[p.Position] = p
// 			}
// 		}
// 	}
// }

// // GenerateThreatMap takes the current state of the board to calculate all threats to each square on the board
// func (b *Board) GenerateThreatMap() {
// 	b.ThreatMap = make(map[move.Position][]piece.Piece)
// 	pieces := b.getRemainingPieces()
// 	//wg := &sync.WaitGroup{}
// 	//mu := &sync.Mutex{}

// 	//wg.Add(len(b.Squares) * len(pieces))
// 	for _, pos := range b.Squares {
// 		for _, p := range pieces {
// 			//go func(pos move.Position, p piece.Piece) {
// 			//defer wg.Done()
// 			if b.IsValidMove(move.Move{From: p.Position, To: pos}) {
// 				if p.GetPieceType() == piece.PieceTypePawn {
// 					file := p.Position.File - pos.File
// 					rank := p.Position.Rank - pos.Rank

// 					x := math.Abs(float64(file))
// 					y := math.Abs(float64(rank))

// 					// Only diagonal moves for a pawn are an attack, therefore forward move is not a threat
// 					if x == 0 && (y == 1 || y == 2) {
// 						return
// 					}
// 				}

// 				//mu.Lock()
// 				b.ThreatMap[pos] = append(b.ThreatMap[pos], p)
// 				//mu.Unlock()
// 			}
// 			//}(pos, p)
// 		}
// 	}

// 	//wg.Wait()
// }

// // GetMoveMapForColour gets all possible moves for the position and colour specified
// func (b Board) GetMoveMapForColour(pos move.Position, c colour.Colour) []piece.Piece {
// 	var pieces []piece.Piece

// 	for _, p := range b.MoveMap[pos] {
// 		if p.Colour == c {
// 			pieces = append(pieces, p)
// 		}
// 	}

// 	return pieces
// }

// // GetAttackingPiecesForColour gets all the pieces that are threatening a square based on the colour provided
// func (b Board) GetAttackingPiecesForColour(pos move.Position, c colour.Colour) []piece.Piece {
// 	var pieces []piece.Piece

// 	for _, p := range b.ThreatMap[pos] {
// 		if p.Colour == c {
// 			pieces = append(pieces, p)
// 		}
// 	}

// 	return pieces
// }

//func (b Board) IsCheck(c colour.Colour) bool {
//	king, err := b.getKing(c)
//	if err != nil {
//		return false
//	}
//
//	return b.isThreatened(king)
//}
//
//// TODO: Try to make this concurrent
//// IsCheckMate checks to see if the colour specified is in check mate
//func (b Board) IsCheckMate(c colour.Colour) bool {
//	k, err := b.getKing(c)
//	if err != nil {
//		return false
//	}
//
//	ps := b.getRemainingPieces(c.Opposite())
//
//	validMoves := b.getValidKingMoves(k)
//
//	for pos := range validMoves {
//		for _, p := range ps {
//			attack := move.Move{
//				From: p.Position,
//				To:   pos,
//			}
//
//			if p.IsValidMove(attack) {
//				validMoves[pos] = false
//			}
//		}
//	}
//
//	for _, valid := range validMoves {
//		if valid {
//			return false
//		}
//	}
//
//	return true
//}
//
//func (b Board) isThreatened(p *piece.Piece) bool {
//	ps := b.getRemainingPieces(p.Colour.Opposite())
//
//	for _, p2 := range ps {
//		if p2.IsValidMove(move.Move{From: p2.Position, To: p.Position}) {
//			return true
//		}
//	}
//
//	return false
//}
//
//func (b Board) getValidKingMoves(k *piece.Piece) map[move.Position]bool {
//	validMoves := make(map[move.Position]bool)
//
//	for f := k.Position.File - 1; f < k.Position.File+2; f++ {
//		for r := k.Position.Rank - 1; r < k.Position.Rank+2; r++ {
//			// If it's not a move from it's current position
//			if f == 0 && r == 0 {
//				continue
//			}
//
//			m := move.Move{
//				From: k.Position,
//				To:   move.Position{File: f, Rank: r},
//			}
//
//			rs := rules.Assert(
//				rules.InBoundsOfBoard(b.Width, b.Height, m),
//				rules.IsNotFriendlyCapture(b.Pieces, m),
//			)
//
//			if err := rs(); err == nil {
//				validMoves[m.To] = true
//			}
//		}
//	}
//
//	return validMoves
//}

// isPinned checks to see if the move provided is possible based on whether it is pinned to it's king or not
//func (b Board) isPinned(m move.Move) (Board, error) {
//	p := b.Pieces[m.From]
//
//	attackedPiece, isAttacking := b.Pieces[m.To]
//
//	p.Position = m.To
//	b.Pieces[m.From] = p
//	b.Pieces[m.To] = b.Pieces[m.From]
//	delete(b.Pieces, m.From)
//
//	b.Update()
//
//	if !b.IsCheck(p.Colour) {
//		if p.GetPieceType() == piece.PieceTypePawn {
//			details := p.PieceDetails.(piece.Pawn)
//			details.HasMoved = true
//			p.PieceDetails = details
//			b.Pieces[m.To] = p
//		}
//
//		return b, nil
//	}
//
//	if isAttacking {
//		b.Pieces[m.To] = attackedPiece
//	}
//
//	p.Position = m.From
//	b.Pieces[m.From] = p
//	delete(b.Pieces, m.To)
//	return Board{}, fmt.Errorf("cannot make move: %v, as you are putting the king in check", m)
//}

// getKing gets the king piece for the colour provided
//func (b Board) getKing(c colour.Colour) (*piece.Piece, error) {
//	for _, p := range b.Pieces {
//		if p.Colour == c && p.GetPieceType() == piece.PieceTypeKing {
//			return p, nil
//		}
//	}
//
//	return &piece.Piece{}, fmt.Errorf("cannot find king")
//}

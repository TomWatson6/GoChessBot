package generation

import (
	"github.com/tomwatson6/chessbot/internal/piece"
)

var (
	pieceMap = map[int]func() piece.PieceDetails{
		0: func() piece.PieceDetails { return piece.Pawn{} },
		1: func() piece.PieceDetails { return piece.Rook{} },
		2: func() piece.PieceDetails { return piece.Knight{} },
		3: func() piece.PieceDetails { return piece.Bishop{} },
		4: func() piece.PieceDetails { return piece.Queen{} },
		5: func() piece.PieceDetails { return piece.King{} },
	}
)

//func GetRandomPiece(c colour.Colour, taken *map[move.Position]piece.PieceID) piece.Piece {
//	var p piece.Piece
//
//	dimX, dimY := config.GetBoardDimensions()
//
//	x := 1
//	y := 0
//
//	for {
//		x = rand.Intn(dimX)
//		y = rand.Intn(dimY)
//
//		if _, ok := (*taken)[move.Position{File: x, Rank: y}]; !ok {
//			break
//		}
//	}
//
//	p.ValidMoves = make(map[move.Position]bool)
//	p.History = make(map[int]move.Position)
//	p.Position = move.Position{File: x, Rank: y}
//	p.Colour = c
//
//	numTypes := config.GetNumPieceTypes()
//
//	for {
//		pieceType := rand.Intn(numTypes)
//
//		chosen := pieceMap[pieceType]()
//
//		s := set.NewFromMapValues(*taken)
//
//		isKing := func(pi piece.PieceID) bool {
//			return pi.PieceType == piece.PieceTypeKing && pi.Colour == c
//		}
//
//		if linq.Any(s.ToArray(), isKing) {
//			if chosen.GetPieceType() == piece.PieceTypeKing {
//				continue
//			}
//		}
//
//		p.PieceDetails = chosen
//
//		break
//	}
//
//	(*taken)[p.Position] = piece.PieceID{
//		Colour:    c,
//		PieceType: p.GetPieceType(),
//	}
//
//	return p
//}

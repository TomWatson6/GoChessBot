package main

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/internal/piece"
	"github.com/tomwatson6/chessbot/pkg/input"
	"github.com/tomwatson6/chessbot/pkg/output"
)

func main() {
	chess := input.Get(colour.White, make(map[move.Position]piece.Piece))
	output.PrintBoard(chess)
}

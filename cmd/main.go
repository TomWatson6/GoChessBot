package main

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/pkg/input"
	"github.com/tomwatson6/chessbot/pkg/output"
)

func main() {
	chess := input.Get(colour.White)
	output.PrintBoard(chess)
}

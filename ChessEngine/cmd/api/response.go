package api

import "github.com/tomwatson6/chessbot/internal/move"

type MoveResponse struct {
	Moves []move.Move `json:"moves"`
	Err   string      `json:"err"`
}

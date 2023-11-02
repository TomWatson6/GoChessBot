package api

import "github.com/tomwatson6/chessbot/internal/colour"

type StartGameRequest struct {
	Colour colour.Colour `json:"colour"`
}

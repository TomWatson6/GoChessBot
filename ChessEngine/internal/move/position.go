package move

import "fmt"

type Position struct {
	File int `json:"file"`
	Rank int `json:"rank"`
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.File, p.Rank)
}

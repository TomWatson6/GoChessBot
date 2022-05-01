package move

import "fmt"

type Move struct {
	From, To Position
}

type Position struct {
	File, Rank int
}

func (m Move) String() string {
	return fmt.Sprintf("%s->%s", m.From.String(), m.To.String())
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.File, p.Rank)
}

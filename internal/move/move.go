package move

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Move struct {
	From Position `json:"from"`
	To   Position `json:"to"`
}

func NewMoveFromString(s string) (Move, error) {
	// Receives a string matching the following: (0,0)->(1,1)
	var m Move

	parts := strings.Split(s, "->")

	left := strings.Trim(parts[0], "()")
	leftParts := strings.Split(left, ",")

	file, err := strconv.Atoi(leftParts[0])
	if err != nil {
		return Move{}, err
	}

	rank, err := strconv.Atoi(leftParts[1])
	if err != nil {
		return Move{}, err
	}

	m.From = Position{File: file, Rank: rank}

	right := strings.Trim(parts[1], "()")
	rightParts := strings.Split(right, ",")

	file, err = strconv.Atoi(rightParts[0])
	if err != nil {
		return Move{}, err
	}

	rank, err = strconv.Atoi(rightParts[1])
	if err != nil {
		return Move{}, err
	}

	m.To = Position{File: file, Rank: rank}

	return m, nil
}

func (m Move) Distance() int {
	dx := math.Abs(float64(m.To.File - m.From.File))
	dy := math.Abs(float64(m.To.Rank - m.From.Rank))

	if dx > dy {
		return int(dx)
	}

	return int(dy)
}

func (m Move) String() string {
	return fmt.Sprintf("%s->%s", m.From.String(), m.To.String())
}

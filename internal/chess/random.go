package chess

import (
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
	"github.com/tomwatson6/chessbot/testing/generation"
	"math/rand"
	"os"
	"time"
)

type RandomChess struct {
	fileName string
	Chess
}

func NewRandom() RandomChess {
	var randomChess RandomChess

	randomChess.fileName = "testing/output/" + time.Now().Format("2006_01_02 15_04_05") + ".txt"

	f, err := os.Create(randomChess.fileName)
	if err != nil {
		panic(err)
	}

	f.Write([]byte(randomChess.fileName + "\n\n"))

	if err := f.Close(); err != nil {
		panic(err)
	}

	randomChess.Chess = New(colour.White)

	return randomChess
}

// TODO: Refactor this to work with full codebase refactor...
func (r *RandomChess) Play() colour.Colour {
	turns := 0

	for !r.Board.IsCheckMate(colour.White) && !r.Board.IsCheckMate(colour.Black) {
		ms := r.getMoveSet()

		for _, m := range ms {
			if err := r.MakeMove(m); err != nil {
				continue
			}
		}

		r.NextTurn()
		r.Board.Update()

		turns += 1

		// TODO: remove this once tested properly
		if turns >= 100 {
			break
		}
	}

	if r.Board.IsCheckMate(colour.White) {
		r.writeToFile(colour.Black.String() + " Won!")
	} else {
		r.writeToFile(colour.White.String() + " Won!")
	}

	return colour.White
}

func (r RandomChess) getMoveSet() []move.Move {
	moves := generation.GetValidMoves(r.Board, r.Turn)
	var setOfMoves [][]move.Move

	for _, m := range moves {
		setOfMoves = append(setOfMoves, []move.Move{m})
	}

	additional := []string{"O-O", "O-O-O"}

	for _, a := range additional {
		ms, err := r.TranslateNotation(a)
		if err != nil {
			panic(err)
		}

		setOfMoves = append(setOfMoves, ms)
	}

	index := rand.Intn(len(setOfMoves))

	return setOfMoves[index]
}

func (r RandomChess) writeToFile(s string) {
	f, err := os.OpenFile(
		r.fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	f.Write([]byte(s))
}

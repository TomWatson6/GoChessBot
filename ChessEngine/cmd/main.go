package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/tomwatson6/chessbot/cmd/api"
	"github.com/tomwatson6/chessbot/generation"
	"github.com/tomwatson6/chessbot/internal/chess"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

var c chess.Chess

// Known bugs:
// - pawn promotion to queen

func getInput(r *http.Request, obj any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}

	return nil
}

func getMove(r *http.Request) (move.Move, error) {
	var move move.Move
	getInput(r, &move)

	return move, nil
}

func startGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var startGameInput api.StartGameRequest
	getInput(r, &startGameInput)

	c = chess.New(startGameInput.Colour)

	state(w, r)
}

func movePiece(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := api.MoveResponse{}
	move, err := getMove(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Err = fmt.Sprintf("%s", err)
		fmt.Fprint(w, resp)
		return
	}

	moves, err := c.MakeMove(move)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Err = fmt.Sprintf("%s", err)
	}

	resp.Moves = moves

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to marshal json with error: %s\n", err)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to write json with error: %s\n", err)
	}
}

func state(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to marshal json with error: %s\n", err)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to write json with error: %s\n", err)
	}
}

func startRandom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c = chess.New(colour.White)
	b := generation.NewBoard(10)
	c.Board = b

	jsonResponse, err := json.Marshal(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to marshal json with error: %s\n", err)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to write json with error: %s\n", err)
	}
}

func power(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()

	fileStr := queryParams.Get("file")
	file, err := strconv.Atoi(fileStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to marshal json with error: %s\n", err)
		return
	}

	rankStr := queryParams.Get("rank")
	rank, err := strconv.Atoi(rankStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to marshal json with error: %s\n", err)
		return
	}

	p := c.Board.Power(file, rank)
	jsonResponse, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to marshal json with error: %s\n", err)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to write json with error: %s\n", err)
	}
}

func main() {
	http.HandleFunc("/start", startGame)
	http.HandleFunc("/startRandom", startRandom)
	http.HandleFunc("/move", movePiece)
	http.HandleFunc("/state", state)
	http.HandleFunc("/power", power)

	fmt.Println("Listening on :8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tomwatson6/chessbot/cmd/api"
	"github.com/tomwatson6/chessbot/internal/chess"
	"github.com/tomwatson6/chessbot/internal/colour"
	"github.com/tomwatson6/chessbot/internal/move"
)

var c chess.Chess

func getUserInput(c colour.Colour) (string, error) {
	fmt.Printf("%s's move: ", c)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text, nil
}

// Known bugs:
// - pawn promotion to queen

func getInput(r *http.Request, obj any) error {
	body, err := ioutil.ReadAll(r.Body)
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
		// fmt.Fprint(w, resp)
		// return
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

// TODO: Get Valid Moves endpoint needs to be added for the AI

func main() {
	http.HandleFunc("/start", startGame)
	http.HandleFunc("/move", movePiece)
	http.HandleFunc("/state", state)

	fmt.Println("Use /start {GET}, /move {POST}, /state {GET} to use ChessBot")
	fmt.Println("Listening on :8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

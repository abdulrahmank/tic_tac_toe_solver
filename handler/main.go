package handler

import (
	"encoding/json"
	"github.com/abdulrahmank/solver/tic_tac_toe/solver"
	"github.com/abdulrahmank/solver/tic_tac_toe/ttt"
	"log"
	"net/http"
)

func Play(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("Request with post method")
		boardJson := &BoardJson{}
		if err := json.NewDecoder(r.Body).Decode(&boardJson); err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		if result, err := play(*boardJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			respond(result, w)
		}
	} else if r.Method == "OPTIONS" {
		log.Println("Request with options method")
		w.Header().Set("Access-Control-Allow-Origin", "null")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func play(boardJson BoardJson) (BoardJson, error) {
	board := boardJson.ConvertToBoard()
	var result BoardJson
	if playerWon(board) {
		board.Init(3, 3)
		result = ConvertToBoardJson(board)
		result.Status = solver.LOST
	} else if len(board.GetEmptyCells()) == 0 {
		board.Init(3, 3)
		result = ConvertToBoardJson(board)
		result.Status = solver.DRAW
	} else {
		analyserImpl := &solver.AnalyserImpl{}
		if gs, err := solver.Solve(board, analyserImpl); err != nil {
			return BoardJson{}, err
		} else {
			result = ConvertToBoardJson(board)
			result.Status = gs
		}
	}
	return result, nil
}

func respond(result BoardJson, w http.ResponseWriter) {
	if bytes, err := json.Marshal(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "null")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(bytes)
	}
}

func playerWon(board ttt.Board) bool {
	return didWinGame(board, ttt.X)
}

func didWinGame(board ttt.Board, character ttt.BoardCharacter) bool {
	for i := 0; i < board.Rows; i++ {
		if board.IsHorizontalWin(i, string(character)) {
			return true
		}
	}
	for i := 0; i < board.Cols; i++ {
		if board.IsVerticalWin(i, string(character)) {
			return true
		}
	}
	return board.IsDiagonalWin(string(character))
}

package handler

import (
	"encoding/json"
	"github.com/abdulrahmank/solver/tic_tac_toe/solver"
	"github.com/abdulrahmank/solver/tic_tac_toe/ttt"
	"net/http"
)

func Play(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cells := make([]CellJson, 0)
		if err := json.NewDecoder(r.Body).Decode(&cells); err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		boardJson := &BoardJson{Cells: cells}
		board := boardJson.ConvertToBoard()

		if playerWon(board) {
			board.Init(3, 3)
			result := ConvertToBoardJson(board)
			result.Status = solver.LOST
			respond(result, w)
			return
		}

		analyserImpl := &solver.AnalyserImpl{}
		if gs, err := solver.Solve(board, analyserImpl); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			result := ConvertToBoardJson(board)
			result.Status = gs
			respond(result, w)
		}
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func respond(result BoardJson, w http.ResponseWriter) {
	if bytes, err := json.Marshal(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		_, _ = w.Write(bytes)
		w.WriteHeader(http.StatusOK)
	}
}

func playerWon(board ttt.Board) bool {
	for i := 0; i < board.Rows; i++ {
		if board.IsHorizontalWin(i, string(ttt.X)) {
			return true
		}
	}
	for i := 0; i < board.Cols; i++ {
		if board.IsVerticalWin(i, string(ttt.X)) {
			return true
		}
	}
	return board.IsDiagonalWin(string(ttt.X))
}

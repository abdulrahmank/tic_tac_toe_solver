package handler

import (
	"encoding/json"
	"github.com/abdulrahmank/solver/tic_tac_toe/solver"
	"github.com/abdulrahmank/solver/tic_tac_toe/ttt"
	"net/http"
)

func Play(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		boardJson := &BoardJson{}
		if err := json.NewDecoder(r.Body).Decode(&boardJson); err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		board := boardJson.ConvertToBoard()

		if playerWon(board) {
			board.Init(3, 3)
			result := ConvertToBoardJson(board)
			result.Status = solver.LOST
			respond(result, w)
		} else if len(board.GetEmptyCells()) == 0 {
			board.Init(3, 3)
			result := ConvertToBoardJson(board)
			result.Status = solver.DRAW
			respond(result, w)
		} else {
			analyserImpl := &solver.AnalyserImpl{}
			if gs, err := solver.Solve(board, analyserImpl); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				result := ConvertToBoardJson(board)
				result.Status = gs
				respond(result, w)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
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

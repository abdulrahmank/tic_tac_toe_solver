package handler

import (
	"encoding/json"
	"github.com/abdulrahmank/solver/tic_tac_toe/solver"
	"net/http"
)

func Play(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		boardJson := &BoardJson{}
		if err := json.NewDecoder(r.Body).Decode(boardJson); err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		board := boardJson.ConvertToBoard()
		analyserImpl := &solver.AnalyserImpl{}
		if err := solver.Solve(board, analyserImpl); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if bytes, err := json.Marshal(ConvertToBoardJson(board)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			_, _ = w.Write(bytes)
			w.WriteHeader(http.StatusOK)
		}
	}
	w.WriteHeader(http.StatusNotImplemented)
}

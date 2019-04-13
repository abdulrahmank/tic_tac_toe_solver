package solver

import (
	"github.com/abdulrahmank/solver/tic_tac_toe/ttt"
	"testing"
)

func TestGetCellWiseWinProbability(t *testing.T) {

	t.Run("Should get 50 probability if equal chance of win", func(t *testing.T) {
		board := ttt.Board{}
		board.Init(3, 3)

		probability := GetCellWiseWinProbability(board, ttt.O)

		if probability[*board.Cells[2][1]] != 50 {
			t.Errorf("Expected 50 but was %d", probability[*board.Cells[2][0]])
		}
	})

}

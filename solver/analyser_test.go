package solver

import (
	"github.com/abdulrahmank/solver/tic_tac_toe/ttt"
	"testing"
)

func TestGetCellWiseWinProbability(t *testing.T) {

	analyserImpl := AnalyserImpl{}
	t.Run("Should get NEUTRAL status if equal chance of win", func(t *testing.T) {
		board := ttt.Board{}
		board.Init(3, 3)

		probability := analyserImpl.GetCellWiseWinProbability(board, ttt.O)

		if probability[*board.Cells[2][1]] != NEUTRAL {
			t.Errorf("Expected NEUTRAL but was %d", probability[*board.Cells[2][1]])
		}
	})

	t.Run("Should get LOSE status if opponent has chance to win", func(t *testing.T) {
		board := ttt.Board{}
		board.Init(3, 3)
		board.Cells[0][0].Val = string(ttt.X)
		board.Cells[0][1].Val = string(ttt.X)
		board.Cells[1][2].Val = string(ttt.O)

		probability := analyserImpl.GetCellWiseWinProbability(board, ttt.O)

		if probability[*board.Cells[0][2]] != LOSE {
			t.Errorf("Expected LOSE but was %d", probability[*board.Cells[0][2]])
		}
	})

	t.Run("Should get WIN status if we have chance to win", func(t *testing.T) {
		board := ttt.Board{}
		board.Init(3, 3)
		board.Cells[0][0].Val = string(ttt.O)
		board.Cells[0][1].Val = string(ttt.O)
		board.Cells[1][2].Val = string(ttt.X)
		board.Cells[2][2].Val = string(ttt.X)

		probability := analyserImpl.GetCellWiseWinProbability(board, ttt.O)

		if probability[*board.Cells[0][2]] != WIN {
			t.Errorf("Expected WIN but was %d", probability[*board.Cells[0][2]])
		}
	})

	t.Run("Should get POTENTIAL_WIN status if we have only one of our characters present in the row/col", func(t *testing.T) {
		board := ttt.Board{}
		board.Init(3, 3)
		board.Cells[0][0].Val = string(ttt.O)

		probability := analyserImpl.GetCellWiseWinProbability(board, ttt.O)

		if probability[*board.Cells[0][2]] != POTENTIAL_WIN {
			t.Errorf("Expected POTENTIAL_WIN but was %d", probability[*board.Cells[0][2]])
		}
	})

	t.Run("Should get POTENTIAL_LOSE status if characters present in the row/col are equal", func(t *testing.T) {
		board := ttt.Board{}
		board.Init(3, 3)
		board.Cells[0][0].Val = string(ttt.O)
		board.Cells[2][2].Val = string(ttt.X)

		probability := analyserImpl.GetCellWiseWinProbability(board, ttt.O)

		if probability[*board.Cells[0][2]] != POTENTIAL_LOSE {
			t.Errorf("Expected POTENTIAL_LOSE but was %d", probability[*board.Cells[0][2]])
		}
	})

	t.Run("Should get POTENTIAL_LOSE status if we have only one of our characters present in the row/col", func(t *testing.T) {
		board := ttt.Board{}
		board.Init(3, 3)
		board.Cells[0][0].Val = string(ttt.O)
		board.Cells[0][2].Val = string(ttt.X)

		probability := analyserImpl.GetCellWiseWinProbability(board, ttt.O)

		if probability[*board.Cells[0][1]] != POTENTIAL_LOSE {
			t.Errorf("Expected POTENTIAL_WIN but was %d", probability[*board.Cells[0][1]])
		}
	})


	t.Run("Should get LOSE status if opponent can win diagonally", func(t *testing.T) {
		board := ttt.Board{}
		board.Init(3, 3)
		board.Cells[0][0].Val = string(ttt.X)
		board.Cells[2][2].Val = string(ttt.X)
		board.Cells[1][0].Val = string(ttt.O)

		probability := analyserImpl.GetCellWiseWinProbability(board, ttt.O)

		if probability[*board.Cells[1][1]] != LOSE {
			t.Errorf("Expected LOSE but was %d", probability[*board.Cells[1][1]])
		}
	})

	t.Run("Should get WIN status if we can win diagonally", func(t *testing.T) {
		board := ttt.Board{}
		board.Init(3, 3)
		board.Cells[0][0].Val = string(ttt.O)
		board.Cells[2][2].Val = string(ttt.O)
		board.Cells[1][0].Val = string(ttt.X)

		probability := analyserImpl.GetCellWiseWinProbability(board, ttt.O)

		if probability[*board.Cells[1][1]] != WIN {
			t.Errorf("Expected WIN but was %d", probability[*board.Cells[1][1]])
		}
	})

}

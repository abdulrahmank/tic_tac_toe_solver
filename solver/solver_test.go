package solver_test

import (
	"github.com/abdulrahmank/solver/tic_tac_toe/solver"
	"github.com/abdulrahmank/solver/tic_tac_toe/solver/mock"
	"github.com/abdulrahmank/solver/tic_tac_toe/ttt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestSolver(t *testing.T) {
	t.Run("Should be able to make a move to prevent losing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		analyser := solver_mock.NewMockAnalyser(ctrl)
		board := ttt.Board{}
		board.Init(3, 3)

		cellStatuses := make(map[ttt.Cell]solver.GameStatus)
		cellStatuses[ttt.Cell{Row: 0, Column: 0}] = solver.LOSE
		cellStatuses[ttt.Cell{Row: 0, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 0, Column: 2}] = solver.POTENTIAL_LOSE

		cellStatuses[ttt.Cell{Row: 1, Column: 0}] = solver.POTENTIAL_WIN
		cellStatuses[ttt.Cell{Row: 1, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 1, Column: 2}] = solver.NEUTRAL

		cellStatuses[ttt.Cell{Row: 2, Column: 0}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 2, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 2, Column: 2}] = solver.NEUTRAL

		analyser.EXPECT().GetCellWiseWinProbability(board, ttt.O).Return(cellStatuses)

		err := solver.Solve(board, analyser)

		if err != nil {
			t.Errorf("Expected nil error")
		}

		if board.Cells[0][0].Val != string(ttt.O) {
			t.Errorf("Expected %s but was %s", ttt.O, board.Cells[0][0].Val)
		}
	})

	t.Run("Should be able to win", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		analyser := solver_mock.NewMockAnalyser(ctrl)
		board := ttt.Board{}
		board.Init(3, 3)

		cellStatuses := make(map[ttt.Cell]solver.GameStatus)
		cellStatuses[ttt.Cell{Row: 0, Column: 0}] = solver.LOSE
		cellStatuses[ttt.Cell{Row: 0, Column: 1}] = solver.WIN
		cellStatuses[ttt.Cell{Row: 0, Column: 2}] = solver.POTENTIAL_LOSE

		cellStatuses[ttt.Cell{Row: 1, Column: 0}] = solver.POTENTIAL_WIN
		cellStatuses[ttt.Cell{Row: 1, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 1, Column: 2}] = solver.NEUTRAL

		cellStatuses[ttt.Cell{Row: 2, Column: 0}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 2, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 2, Column: 2}] = solver.NEUTRAL

		analyser.EXPECT().GetCellWiseWinProbability(board, ttt.O).Return(cellStatuses)

		err := solver.Solve(board, analyser)

		if err != nil {
			t.Errorf("Expected nil error")
		}

		if board.Cells[0][1].Val != string(ttt.O) {
			t.Errorf("Expected %s but was %s", ttt.O, board.Cells[0][0].Val)
		}
	})

	t.Run("Should be able to make a move to prevent potential loss", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		analyser := solver_mock.NewMockAnalyser(ctrl)
		board := ttt.Board{}
		board.Init(3, 3)

		cellStatuses := make(map[ttt.Cell]solver.GameStatus)
		cellStatuses[ttt.Cell{Row: 0, Column: 0}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 0, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 0, Column: 2}] = solver.POTENTIAL_LOSE

		cellStatuses[ttt.Cell{Row: 1, Column: 0}] = solver.POTENTIAL_WIN
		cellStatuses[ttt.Cell{Row: 1, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 1, Column: 2}] = solver.NEUTRAL

		cellStatuses[ttt.Cell{Row: 2, Column: 0}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 2, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 2, Column: 2}] = solver.NEUTRAL


		analyser.EXPECT().GetCellWiseWinProbability(board, ttt.O).Return(cellStatuses)

		err := solver.Solve(board, analyser)

		if err != nil {
			t.Errorf("Expected nil error")
		}

		if board.Cells[0][2].Val != string(ttt.O) {
			t.Errorf("Expected %s but was %s", ttt.O, board.Cells[0][0].Val)
		}
	})

	t.Run("Should be able to make acquire a  potential win", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		analyser := solver_mock.NewMockAnalyser(ctrl)
		board := ttt.Board{}
		board.Init(3, 3)

		cellStatuses := make(map[ttt.Cell]solver.GameStatus)
		cellStatuses[ttt.Cell{Row: 0, Column: 0}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 0, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 0, Column: 2}] = solver.NEUTRAL

		cellStatuses[ttt.Cell{Row: 1, Column: 0}] = solver.POTENTIAL_WIN
		cellStatuses[ttt.Cell{Row: 1, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 1, Column: 2}] = solver.NEUTRAL

		cellStatuses[ttt.Cell{Row: 2, Column: 0}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 2, Column: 1}] = solver.NEUTRAL
		cellStatuses[ttt.Cell{Row: 2, Column: 2}] = solver.NEUTRAL


		analyser.EXPECT().GetCellWiseWinProbability(board, ttt.O).Return(cellStatuses)

		err := solver.Solve(board, analyser)

		if err != nil {
			t.Errorf("Expected nil error")
		}

		if board.Cells[1][0].Val != string(ttt.O) {
			t.Errorf("Expected %s but was %s", ttt.O, board.Cells[0][0].Val)
		}
	})
}

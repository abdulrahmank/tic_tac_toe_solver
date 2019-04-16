package solver

import (
	"errors"
	"github.com/abdulrahmank/solver/tic_tac_toe/ttt"
)

func Solve(board ttt.Board, analyser Analyser) error {
	cells := analyser.GetCellWiseWinProbability(board, ttt.O)
	cellBucket := make(map[CellStatus][]ttt.Cell)
	for _, gs := range []CellStatus{WIN, LOSE, POTENTIAL_LOSE, POTENTIAL_WIN, NEUTRAL} {
		cellBucket[gs] = make([]ttt.Cell, 0)
	}

	for cell, status := range cells {
		cellBucket[status] = append(cellBucket[status], cell)
	}

	row, column := -1, -1

	if len(cellBucket[WIN]) > 0 {
		cell := cellBucket[WIN][0]
		row = cell.Row
		column = cell.Column
	} else if len(cellBucket[LOSE]) > 0 {
		cell := cellBucket[LOSE][0]
		row = cell.Row
		column = cell.Column
	} else if len(cellBucket[POTENTIAL_LOSE]) > 0 {
		cell := cellBucket[POTENTIAL_LOSE][0]
		row = cell.Row
		column = cell.Column
	} else if len(cellBucket[POTENTIAL_WIN]) > 0 {
		cell := cellBucket[POTENTIAL_WIN][0]
		row = cell.Row
		column = cell.Column
	} else if len(cellBucket[NEUTRAL]) > 0 {
		cell := cellBucket[NEUTRAL][0]
		row = cell.Row
		column = cell.Column
	}

	if row == -1 || column == -1 {
		return errors.New("can't place")
	}

	board.Cells[row][column].Val = string(ttt.O)

	return nil
}

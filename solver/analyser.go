package solver

import "github.com/abdulrahmank/solver/tic_tac_toe/ttt"

type GameStatus int

const (
	LOSE           GameStatus = -1
	NEUTRAL        GameStatus = 0
	WIN            GameStatus = 1
	POTENTIAL_WIN  GameStatus = 2
	POTENTIAL_LOSE GameStatus = 3
)

func GetCellWiseWinProbability(b ttt.Board, c ttt.BoardCharacter) map[ttt.Cell]GameStatus {
	result := make(map[ttt.Cell]GameStatus)
	for _, row := range b.Cells {
		for _, cell := range row {
			result[*cell] = NEUTRAL
			result[*cell] = NEUTRAL
		}
	}

	for i := 0; i < b.Rows; i++ {
		for j := 0; j < b.Cols; j++ {
			rowStatus := make(map[string]int)
			colStatus := make(map[string]int)

			for ti := 0; ti < b.Rows; ti++ {
				rowStatus[b.Cells[ti][j].Val] += 1
			}

			for ti := 0; ti < b.Cols; ti++ {
				colStatus[b.Cells[i][ti].Val] += 1
			}

			if rowStatus[string(c)]|colStatus[string(c)] >= 2 {
				result[*b.Cells[i][j]] = WIN
			} else {
				if rowStatus[string(c)]|colStatus[string(c)] >= 1 {
					result[*b.Cells[i][j]] = POTENTIAL_WIN
				}
				chars := []ttt.BoardCharacter{ttt.X, ttt.O}
				for _, ch := range chars {
					if ch != c {
						if rowStatus[string(ch)]|colStatus[string(ch)] >= 2 {
							result[*b.Cells[i][j]] = LOSE
						} else if rowStatus[string(ch)]|colStatus[string(ch)] >= 1 {
							result[*b.Cells[i][j]] = POTENTIAL_LOSE
						}
					}
				}
			}
			//Leaving diagonal for now
		}
	}

	return result
}

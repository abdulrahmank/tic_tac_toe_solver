package solver

import "github.com/abdulrahmank/solver/tic_tac_toe/ttt"

func GetCellWiseWinProbability(b ttt.Board, c ttt.BoardCharacter) map[ttt.Cell]int {
	result := make(map[ttt.Cell]int)
	for _, row := range b.Cells {
		for _, cell := range row {
			result[*cell] = 50
		}
	}
	return result
}

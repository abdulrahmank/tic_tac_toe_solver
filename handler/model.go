package handler

import (
	"fmt"
	"github.com/abdulrahmank/solver/tic_tac_toe/solver"
	"github.com/abdulrahmank/solver/tic_tac_toe/ttt"
	"strconv"
	"strings"
)

type BoardJson struct {
	Cells  []CellJson
	Status solver.GameStatus
}

type CellJson struct {
	Position string  `json:"position"`
	Value    *string `json:"value"`
}

const BOARD_SIZE = 3

func (b *BoardJson) ConvertToBoard() ttt.Board {
	cells := make([][]*ttt.Cell, BOARD_SIZE)
	for i := 0; i < BOARD_SIZE; i++ {
		cells[i] = make([]*ttt.Cell, 0)
	}

	for _, c := range b.Cells {
		row, cell := c.ConvertToCell()
		cells[row] = append(cells[row], cell)
	}

	board := ttt.Board{}
	board.Init(BOARD_SIZE, BOARD_SIZE)
	board.Cells = cells
	return board
}

func (c *CellJson) ConvertToCell() (int, *ttt.Cell) {
	positionArr := strings.Split(c.Position, ",")
	row, _ := strconv.Atoi(positionArr[0])
	col, _ := strconv.Atoi(positionArr[1])
	var upper string
	if c.Value != nil {
		upper = strings.ToUpper(*c.Value)
	}
	return row, &ttt.Cell{Row: row, Column: col, Val: upper}
}

func ConvertToBoardJson(b ttt.Board) BoardJson {
	cJsons := make([]CellJson, 0)
	for _, row := range b.Cells {
		for _, cell := range row {
			var val *string
			if cell.Val != "" {
				val = &cell.Val
			} else {
				val = nil
			}
			cellJson := CellJson{Position: fmt.Sprintf("%d,%d", cell.Row, cell.Column), Value: val}
			cJsons = append(cJsons, cellJson)
		}
	}
	return BoardJson{Cells: cJsons}
}

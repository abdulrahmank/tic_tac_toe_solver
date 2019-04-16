package handler

import (
	"encoding/json"
	"github.com/abdulrahmank/solver/tic_tac_toe/ttt"
	"strings"
	"testing"
)

func TestCellJson_ConvertToCell(t *testing.T) {
	i := "X"
	s := &i
	cellJson := CellJson{Position: "0,1", Value: s}
	row, cell := cellJson.ConvertToCell()

	if row != 0 {
		t.Errorf("Expected %d but was %d", 0, row)
	}

	if cell.Row != 0 {
		t.Errorf("Expected %d but was %d", 0, cell.Row)
	}

	if cell.Column != 1 {
		t.Errorf("Expected %d but was %d", 1, cell.Column)
	}

	if cell.Val != "X" {
		t.Errorf("Expected %s but was %s", "X", cell.Val)
	}
}

func TestBoardJson_ConvertToBoard(t *testing.T) {

	s := "X"
	q := "O"

	x := &s
	o := &q

	boardJson := BoardJson{Cells: []CellJson{
		{Position: "0,0", Value: x},
		{Position: "0,1", Value: o},
		{Position: "0,2", Value: nil},

		{Position: "1,0", Value: x},
		{Position: "1,1", Value: o},
		{Position: "1,2", Value: o},

		{Position: "2,0", Value: x},
		{Position: "2,1", Value: x},
		{Position: "2,2", Value: o},
	}}

	board := boardJson.ConvertToBoard()

	i2 := ""
	i := &i2
	expectedVal := []*string{x, o, i, x, o, o, x, x, o}
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.Cells[i][j].Row != i {
				t.Errorf("Expected %d but was %d", i, board.Cells[i][j].Row)
			}
			if board.Cells[i][j].Column != j {
				t.Errorf("Expected %d but was %d", j, board.Cells[i][j].Column)
			}
			if board.Cells[i][j].Val != *expectedVal[count] {
				t.Errorf("Expected %s but was %s", *expectedVal[count], board.Cells[i][j].Val)
			}
			count += 1
		}
	}
}

func TestConvertToBoardJson(t *testing.T) {
	board := ttt.Board{}
	board.Init(3, 3)

	board.Cells[0][0].Val = "X"
	board.Cells[0][1].Val = "X"
	board.Cells[0][2].Val = "X"

	board.Cells[1][0].Val = "X"
	board.Cells[1][1].Val = "O"
	board.Cells[1][2].Val = "X"

	boardJson := ConvertToBoardJson(board)

	x := "X"
	o := "O"
	empty := ""

	expectedVals := []*string{&x, &x, &x, &x, &o, &x, &empty, &empty, &empty}
	expectedPos := []string{"0,0", "0,1", "0,2", "1,0", "1,1", "1,2", "2,0", "2,1", "2,2"}

	for index, c := range boardJson.Cells {
		if c.Position != expectedPos[index] {
			t.Errorf("Expected %s but was %s", expectedPos[index], c.Position)
		}

		if *c.Value != *expectedVals[index] {
			t.Errorf("Expected %s but was %s", *expectedVals[index], *c.Value)
		}
	}
}

func TestShouldBeAbleToUnmarshalJsonIntoCellJson(t *testing.T) {
	jsonStr := "[{\"position\": \"1,1\",\"value\": \"x\"}]"

	cells := make([]CellJson, 0)

	if err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(&cells); err != nil {
		t.Errorf("Expected nil but was %v", err)
	}

	if *cells[0].Value != "x" {
		t.Errorf("Expected %s but was %s", "x", *cells[0].Value)
	}
}

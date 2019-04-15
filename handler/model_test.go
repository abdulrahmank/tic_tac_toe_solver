package handler

import "testing"

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

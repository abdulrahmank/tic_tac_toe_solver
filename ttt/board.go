package ttt

import (
	"errors"
)

type Board struct {
	Rows, Cols int
	Cells      [][]*Cell
}

func (b *Board) Init(rows, cols int) {
	b.Rows = rows
	b.Cols = cols
	b.Cells = make([][]*Cell, 0)
	for i := 0; i < rows; i++ {
		cells := make([]*Cell, 0)
		for j := 0; j < cols; j++ {
			cells = append(cells, &Cell{Row: i, Column: j})
		}
		b.Cells = append(b.Cells, cells)
	}
}

func (b *Board) AddValToLeft(row, col int, val string) (bool, error) {
	if col-1 >= 0 {
		b.Cells[row][col-1].Val = val
		if b.IsHorizontalWin(row, val) || b.IsVerticalWin(col-1, val) || b.IsDiagonalWin(val) {
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("can't place")
}

func (b *Board) AddValToTopOf(row, col int, val string) (bool, error) {
	if row-1 >= 0 {
		b.Cells[row-1][col].Val = val
		if b.IsHorizontalWin(row-1, val) || b.IsVerticalWin(col, val) || b.IsDiagonalWin(val) {
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("can't place")
}

func (b *Board) AddValToBottomOf(row, col int, val string) (bool, error) {
	if row+1 >= 0 {
		b.Cells[row+1][col].Val = val
		if b.IsHorizontalWin(row+1, val) || b.IsVerticalWin(col, val) || b.IsDiagonalWin(val) {
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("can't place")
}

func (b *Board) AddValToRight(row, col int, val string) (bool, error) {
	if col+1 < len(b.Cells[0]) {
		b.Cells[row][col+1].Val = val
		if b.IsHorizontalWin(row, val) || b.IsVerticalWin(col, val) || b.IsDiagonalWin(val) {
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("can't place")
}

func (b *Board) IsHorizontalWin(row int, val string) bool {
	for i := 0; i < b.Cols; i++ {
		if b.Cells[row][i].Val != val {
			return false
		}
	}
	return true
}

func (b *Board) IsVerticalWin(col int, val string) bool {
	for i := 0; i < b.Rows; i++ {
		if b.Cells[i][col].Val != val {
			return false
		}
	}
	return true
}

func (b *Board) IsDiagonalWin(val string) bool {
	for i := 0; i < b.Rows; i++ {
		if b.Cells[i][i].Val != val {
			return false
		}
	}
	return true
}

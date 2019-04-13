package ttt_test

import (
	"github.com/abdulrahmank/solver/tic_tac_toe/ttt"
	"testing"
)

func TestBoard_Init(t *testing.T) {
	board := ttt.Board{}
	board.Init(3, 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.Cells[i][j].Row != i && board.Cells[i][j].Column != j {
				t.Errorf("Expected %d, %d, but was %d, %d", i, i, board.Cells[i][j].Row, board.Cells[i][j].Column)
			}
		}
	}
}

func TestBoard_AddValToRight(t *testing.T) {
	board := ttt.Board{}
	board.Init(3, 3)
	_, err := board.AddValToRight(0, 0, "X")
	if err != nil {
		t.Error("Expected nil")
	}

	if board.Cells[0][1].Val != "X" {
		t.Errorf("Expected X but was %v", board.Cells[0][1].Val)
	}
}

func TestBoard_AddValToLeft(t *testing.T) {
	board := ttt.Board{}
	board.Init(3, 3)
	_, err := board.AddValToLeft(0, 1, "X")
	if err != nil {
		t.Error("Expected nil")
	}

	if board.Cells[0][0].Val != "X" {
		t.Errorf("Expected X but was %v", board.Cells[0][1].Val)
	}
}

func TestBoard_HorizontalWin(t *testing.T) {
	board := ttt.Board{}
	board.Init(3, 3)

	board.Cells[0][0].Val = "X"
	board.Cells[0][1].Val = "X"
	board.Cells[0][2].Val = "X"

	board.Cells[1][0].Val = "X"
	board.Cells[1][1].Val = "Y"
	board.Cells[1][2].Val = "X"

	if !board.IsHorizontalWin(0, "X") {
		t.Error("Expected win")
	}

	if board.IsHorizontalWin(1, "X") {
		t.Error("Din't expect a win")
	}
}


func TestBoard_VerticalWin(t *testing.T) {
	board := ttt.Board{}
	board.Init(3, 3)

	board.Cells[0][0].Val = "X"
	board.Cells[1][0].Val = "X"
	board.Cells[2][0].Val = "X"

	board.Cells[0][1].Val = "X"
	board.Cells[1][1].Val = "Y"
	board.Cells[2][1].Val = "X"

	if !board.IsVerticalWin(0, "X") {
		t.Error("Expected win")
	}

	if board.IsVerticalWin(1, "X") {
		t.Error("Din't expect a win")
	}
}

func TestBoard_DiagonalWin(t *testing.T) {
	board := ttt.Board{}
	board.Init(3, 3)

	board.Cells[0][0].Val = "X"
	board.Cells[1][1].Val = "X"
	board.Cells[2][2].Val = "X"

	if !board.IsDiagonalWin("X") {
		t.Error("Expected win")
	}

	if board.IsDiagonalWin("O") {
		t.Error("Din't expect a win")
	}
}


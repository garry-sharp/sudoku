package main

import (
	"os"
	"reflect"
	"testing"
)

var data1 [][]byte = [][]byte{
	{byte("5"[0]), byte("3"[0]), byte("."[0]), byte("."[0]), byte("7"[0]), byte("."[0]), byte("."[0]), byte("."[0]), byte("."[0])},
	{byte("6"[0]), byte("."[0]), byte("."[0]), byte("1"[0]), byte("9"[0]), byte("5"[0]), byte("."[0]), byte("."[0]), byte("."[0])},
	{byte("."[0]), byte("9"[0]), byte("8"[0]), byte("."[0]), byte("."[0]), byte("."[0]), byte("."[0]), byte("6"[0]), byte("."[0])},
	{byte("8"[0]), byte("."[0]), byte("."[0]), byte("."[0]), byte("6"[0]), byte("."[0]), byte("."[0]), byte("."[0]), byte("3"[0])},
	{byte("4"[0]), byte("."[0]), byte("."[0]), byte("8"[0]), byte("."[0]), byte("3"[0]), byte("."[0]), byte("."[0]), byte("1"[0])},
	{byte("7"[0]), byte("."[0]), byte("."[0]), byte("."[0]), byte("2"[0]), byte("."[0]), byte("."[0]), byte("."[0]), byte("6"[0])},
	{byte("."[0]), byte("6"[0]), byte("."[0]), byte("."[0]), byte("."[0]), byte("."[0]), byte("2"[0]), byte("8"[0]), byte("."[0])},
	{byte("."[0]), byte("."[0]), byte("."[0]), byte("4"[0]), byte("1"[0]), byte("9"[0]), byte("."[0]), byte("."[0]), byte("5"[0])},
	{byte("."[0]), byte("."[0]), byte("."[0]), byte("."[0]), byte("8"[0]), byte("."[0]), byte("."[0]), byte("7"[0]), byte("9"[0])},
}

var data1Solved [][]byte = [][]byte{
	{byte("5"[0]), byte("3"[0]), byte("4"[0]), byte("6"[0]), byte("7"[0]), byte("8"[0]), byte("9"[0]), byte("1"[0]), byte("2"[0])},
	{byte("6"[0]), byte("7"[0]), byte("2"[0]), byte("1"[0]), byte("9"[0]), byte("5"[0]), byte("3"[0]), byte("4"[0]), byte("8"[0])},
	{byte("1"[0]), byte("9"[0]), byte("8"[0]), byte("3"[0]), byte("4"[0]), byte("2"[0]), byte("5"[0]), byte("6"[0]), byte("7"[0])},
	{byte("8"[0]), byte("5"[0]), byte("9"[0]), byte("7"[0]), byte("6"[0]), byte("1"[0]), byte("4"[0]), byte("2"[0]), byte("3"[0])},
	{byte("4"[0]), byte("2"[0]), byte("6"[0]), byte("8"[0]), byte("5"[0]), byte("3"[0]), byte("7"[0]), byte("9"[0]), byte("1"[0])},
	{byte("7"[0]), byte("1"[0]), byte("3"[0]), byte("9"[0]), byte("2"[0]), byte("4"[0]), byte("8"[0]), byte("5"[0]), byte("6"[0])},
	{byte("9"[0]), byte("6"[0]), byte("1"[0]), byte("5"[0]), byte("3"[0]), byte("7"[0]), byte("2"[0]), byte("8"[0]), byte("4"[0])},
	{byte("2"[0]), byte("8"[0]), byte("7"[0]), byte("4"[0]), byte("1"[0]), byte("9"[0]), byte("6"[0]), byte("3"[0]), byte("5"[0])},
	{byte("3"[0]), byte("4"[0]), byte("5"[0]), byte("2"[0]), byte("8"[0]), byte("6"[0]), byte("1"[0]), byte("7"[0]), byte("9"[0])},
}

var column1Unfilled = []string{"5", "3", ".", ".", "7", ".", ".", ".", "."}
var row1Unfilled = []string{"5", "6", ".", "8", "4", "7", ".", ".", "."}
var column1Filled = []string{"5", "3", "4", "6", "7", "8", "9", "1", "2"}
var row1Filled = []string{"5", "6", "1", "8", "4", "7", "9", "2", "3"}

var grid *Grid = nil

func TestMain(m *testing.M) {
	grid = ParseBoard(data1)
	code := m.Run()
	grid = nil
	os.Exit(code)
}

func TestCell_GetNext(t *testing.T) {
	results := []string{"5", "3", ".", ".", "7", ".", ".", ".", ".", "6", "."}
	grid := ParseBoard(data1)
	cell := grid.AllCells()[0]
	for i, result := range results {
		if i > 10 {
			t.Error("Failed badly")
		}
		if cell.Value != result {
			t.Error("Get next failed, actual coordinates ", cell.Column.index, cell.Row.index)
		}
		cell = cell.GetNext()
	}
}

func TestCell_GetCoordinate(t *testing.T) {
	// 4 rows in and 2 deep
	cell := grid.AllCells()[34]
	if col, row := cell.GetCoordinate(); col != 7 && row != 3 {
		t.Errorf("Expected %d, %d and received %d, %d", 7, 3, col, row)
	}
}

func TestCell_GetPossibilities(t *testing.T) {
	received := grid.AllCells()[2].GetPossibilities()
	expected := []string{"1", "2", "4"}
	if !reflect.DeepEqual(expected, received) {
		t.Errorf("Expected %s Received %s", expected, received)
	}
}

func TestParseBoard(t *testing.T) {
	for i, cell := range grid.Rows[0].Cells {
		if cell.Value != column1Unfilled[i] {
			t.Errorf("Expected %s Received %s", column1Unfilled[i], cell.Value)
		}
	}
	for i, cell := range grid.Columns[0].Cells {
		if cell.Value != row1Unfilled[i] {
			t.Errorf("Expected %s Received %s", row1Unfilled[i], cell.Value)
		}
	}
}

func Test_gridToBoard(t *testing.T) {
	board := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		board[i] = make([]byte, 9)
	}
	gridToBoard(board, grid)
	for i, _ := range data1 {
		for j, _ := range data1[i] {
			if data1[i][j] != board[i][j] {
				t.Errorf("Expected %b Received %b", data1[i][j], board[i][j])
			}
		}
	}
}

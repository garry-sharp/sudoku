package main

import (
	"fmt"
	"log"
	"strings"
)

var board [][]byte = [][]byte{}

type Grid struct {
	Rows    []*Row
	Columns []*Column
	Square  []*Square
}

func (g *Grid) AllCells() []*Cell {
	cells := []*Cell{}
	for _, row := range g.Rows {
		for _, cell := range row.Cells {
			cells = append(cells, cell)
		}
	}
	return cells
}

type Row struct {
	Cells []*Cell
	index int
}

type Column struct {
	Cells []*Cell
	index int
}

type Square struct {
	Cells []*Cell
	index int
}

type Cell struct {
	Value          string
	PossibleValues []string
	Square         *Square
	Row            *Row
	Column         *Column
	Grid           *Grid
}

func (cell *Cell) GetCoordinate() (col int, row int) {
	if cell == nil {
		return -1, -1
	}
	col = cell.Column.index
	row = cell.Row.index
	return
}

func (cell *Cell) GetNext() *Cell {
	if cell.Column.index < 8 {
		return cell.Row.Cells[cell.Column.index+1]
	} else if cell.Row.index < 8 {
		return cell.Grid.Rows[cell.Row.index+1].Cells[0]
	} else {
		return nil
	}
}

func (cell *Cell) GetPossibilities() []string {
	possibilities := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, iteratedCell := range append(append(cell.Row.Cells, cell.Column.Cells...), cell.Square.Cells...) {
		if iteratedCell != cell {
			for i, possibility := range possibilities {
				if iteratedCell.Value == possibility {
					possibilities = append(possibilities[:i], possibilities[i+1:]...)
				}
			}
		}
	}
	return possibilities
}

// Parse the bytes data to a Grid struct
func ParseBoard(board [][]byte) *Grid {
	grid := Grid{}
	for i := 0; i < 9; i++ {
		grid.Columns = append(grid.Columns, &Column{index: i})
		grid.Rows = append(grid.Rows, &Row{index: i})
		grid.Square = append(grid.Square, &Square{index: i})
	}
	for rowI, _ := range board {
		for colI, _ := range board[rowI] {
			val := fmt.Sprintf("%c", board[rowI][colI])
			squareI := 0
			if colI >= 6 {
				squareI = 2
			} else if colI >= 3 {
				squareI = 1
			}
			if rowI >= 6 {
				squareI += 6
			} else if rowI >= 3 {
				squareI += 3
			}

			col := grid.Columns[colI]
			row := grid.Rows[rowI]
			square := grid.Square[squareI]

			cell := Cell{
				Value:  val,
				Column: col,
				Row:    row,
				Square: square,
				Grid:   &grid,
			}

			col.Cells = append(col.Cells, &cell)
			row.Cells = append(row.Cells, &cell)
			square.Cells = append(square.Cells, &cell)
		}
	}
	return &grid
}

func solveSudoku(board [][]byte) {
	grid := ParseBoard(board)
	printGrid(grid)
	for _, cell := range grid.AllCells() {
		UpdateCellPossibles(cell)
	}

	backtrackSolve(grid.AllCells()[0])
	printGrid(grid)
	gridToBoard(board, grid)
}

func gridToBoard(board [][]byte, grid *Grid) {
	for _, cell := range grid.AllCells() {
		board[cell.Row.index][cell.Column.index] = []byte(cell.Value)[0]
	}
}

func UpdateCellPossibles(cell *Cell) {
	if cell.Value == "." {
		cell.PossibleValues = cell.GetPossibilities()
		if len(cell.PossibleValues) == 1 {
			cell.Value = cell.PossibleValues[0]
			cell.PossibleValues = []string{}
		}

		if cell.Value != "." {
			cells := append(append(cell.Column.Cells, cell.Row.Cells...), cell.Square.Cells...)
			for _, cell := range cells {
				UpdateCellPossibles(cell)
			}
		}
	}
}

func backtrackSolve(cell *Cell) bool {
	if cell == nil {
		return true
	}

	next := cell.GetNext()

	//If the cell is already filled then move onto the next
	if cell.Value != "." {
		return backtrackSolve(cell.GetNext())
	}

	//If there are no possibilities return false to the parent.
	possibilities := cell.GetPossibilities()
	if len(possibilities) == 0 {
		return false
	}

	for i, possValue := range possibilities {
		//Sequentially try each possibility, with all later cells being examined for each one before moving on (if at all)
		cell.Value = possValue
		if next == nil {
			return true
		} else {
			//This will return true if all cells are examined without running out of possibilities.
			//Otherwise if possibilities run out later then it will return false.
			result := backtrackSolve(next)
			if result == false {
				//If it's the last option in this cell too, then return false to the parent caller for it to iterate.
				//Otherwise it will test the next value.
				if i == len(possibilities)-1 {
					cell.Value = "."
					return false
				}
			} else {
				break
			}
		}
	}

	//This should never be called, returns of true/false should always be handled above. So if the code gets here throw and error.
	log.Fatalln("Error")
	return false
}

func printGrid(grid *Grid) {
	cells := grid.AllCells()
	fmt.Println(strings.Repeat("-", 25))
	for i, cell := range cells {
		if i == len(cells)-1 {
			fmt.Println(cell.Value, "| ")
		} else if i%27 == 0 && i != 0 {
			fmt.Println("| ")
			fmt.Println(strings.Repeat("-", 25))
			fmt.Print("| ", cell.Value, " ")
		} else if i%9 == 0 && i != 0 {
			fmt.Println("| ")
			fmt.Print("| ", cell.Value, " ")
		} else if i%3 == 0 {
			fmt.Print("| ", cell.Value, " ")
		} else {
			fmt.Print(cell.Value, " ")
		}
	}
	fmt.Println(strings.Repeat("-", 25))
}

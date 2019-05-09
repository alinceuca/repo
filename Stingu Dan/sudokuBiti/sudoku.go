package main

import "fmt"

func Solve(sudoku [9][9]uint16) bool {
	a := [3][9]uint16{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			position := sudoku[i][j]
			row := a[0][i]
			column := a[1][j]
			square := a[2][(3*(i/3))+j/3]

			if (row & (1 << position)) == (1 << position) || (column & (1 << position)) == (1 << position) || (square & (1 << position)) == (1 << position) {
				return false
			}
			a[0][i] = row | (1 << position)
			a[1][j] = column | (1 << position)
			a[2][(3*(i/3))+j/3] = square | (1 << position)
		}
	}
	return true
}

func main() {
	sudoku := [9][9]uint16{
		{3, 1, 6, 5, 7, 8, 4, 9, 2},
		{5, 2, 9, 1, 3, 4, 7, 6, 8},
		{4, 8, 7, 6, 2, 9, 5, 3, 1},
		{2, 6, 3, 4, 1, 5, 9, 8, 7},
		{9, 7, 4, 8, 6, 3, 1, 2, 5},
		{8, 5, 1, 7, 9, 2, 6, 4, 3},
		{1, 3, 8, 9, 4, 7, 2, 5, 6},
		{6, 9, 2, 3, 5, 1, 8, 7, 4},
		{7, 4, 5, 2, 8, 6, 3, 1, 9},
	}
	if(Solve(sudoku)) {
		fmt.Println("Sudoku corect")
	} else {
		fmt.Println("Sudoku incorect")
	}
}
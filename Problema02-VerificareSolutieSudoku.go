package main

import "fmt"

func main() {
	sudoku := [9][9]uint{
		{9, 6, 3, 8, 9, 4, 1, 7, 5},
		{3, 4, 5, 2, 7, 1, 6, 8, 9},
		{9, 7, 8, 6, 5, 3, 1, 4, 2},
		{1, 5, 6, 3, 4, 2, 7, 9, 8},
		{8, 3, 7, 5, 6, 9, 2, 1, 4},
		{4, 9, 2, 1, 8, 7, 5, 3, 6},
		{5, 8, 9, 7, 1, 6, 4, 2, 3},
		{7, 2, 4, 9, 3, 5, 8, 6, 1},
		{6, 1, 3, 4, 2, 8, 9, 5, 7},
	}
	if check(sudoku) {
		fmt.Println("Rezolvare corecta!")
	}
}

func check(sudoku [9][9]uint) bool {
	var a = [3][9]uint{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			row := a[0][i]
			column := a[1][j]
			square := a[2][3*(i/3)+(j/3)]
			position := sudoku[i][j]
			if (row & (1 << position)) == (1 << position) {
				fmt.Println("Greseala pe randul ", i)
				return false
			} else {
				a[0][i] = row | (1 << position)
			}
			if (column & (1 << position)) == (1 << position) {
				fmt.Println("Greseala pe coloana ", j)
				return false
			} else {
				a[1][j] = column | (1 << position)
			}
			if (square & (1 << position)) == (1 << position) {
				fmt.Println("Greseala in patrat")
				return false
			} else {
				a[2][3*(i/3)+(j/3)] = square | (1 << position)
			}
		}
	}
	return true
}

//02. Sudoku - verificare configuratie folosind operatii pe biti

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sudoku := validSudokuConfiguration()

	fmt.Println(sudoku)
	fmt.Println(isValidSudoku(sudoku))
}

func isValidSudoku(sudoku [9][9]uint) bool {
	a := [3][9]uint{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			poz := sudoku[i][j]
			line := a[0][i]
			column := a[1][j]
			square := a[2][(3*(i/3))+j/3]
			var bitPoz uint = 1 << poz
			if (line&bitPoz) == bitPoz || (column&bitPoz) == bitPoz || (square&bitPoz) == bitPoz {
				return false
			}
			a[0][i] = line | bitPoz
			a[1][j] = column | bitPoz
			a[2][(3*(i/3))+j/3] = square | bitPoz
		}
	}
	return true
}

func validSudokuConfiguration() [9][9]uint {
	return [9][9]uint{
		{6, 3, 9, 5, 7, 4, 1, 8, 2},
		{5, 4, 1, 8, 2, 9, 3, 7, 6},
		{7, 8, 2, 6, 1, 3, 9, 5, 4},
		{1, 9, 8, 4, 6, 7, 5, 2, 3},
		{3, 6, 5, 9, 8, 2, 4, 1, 7},
		{4, 2, 7, 1, 3, 5, 8, 6, 9},
		{9, 5, 6, 7, 4, 8, 2, 3, 1},
		{8, 1, 3, 2, 9, 6, 7, 4, 5},
		{2, 7, 4, 3, 5, 1, 6, 9, 8},
	}
}

func randomSudokuConfiguration() [9][9]uint {
	sudoku := [9][9]uint{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku[i][j] = uint(rand.Intn(9))
		}
	}
	return sudoku
}

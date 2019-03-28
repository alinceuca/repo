package main

import (
	"fmt"
)

func main() {
	val := validSudokuConfiguration()

	fmt.Println(val)
	fmt.Println(isValidSudoku(val...))
}

func isValidSudoku(val ...uint16) bool {

	a := [3][9]uint16{}

	for index, value := range val {
		i := index / 9
		j := index % 9

		row := a[0][i]
		column := a[1][j]
		square := a[2][(3*(i/3))+j/3]
		position := value

		if position == 0 {
			return false
		}
		if (row & (1 << position)) == (1 << position) {
			fmt.Println(index, position, "row")
			return false
		}
		if (column & (1 << position)) == (1 << position) {
			fmt.Println(index, position, "column")
			return false
		}

		if (square & (1 << position)) == (1 << position) {
			fmt.Println(index, position, "square")
			return false
		}

		a[0][i] = row | (1 << position)
		a[1][j] = column | (1 << position)
		a[2][(3*(i/3))+j/3] = square | (1 << position)
	}

	return true

}
func validSudokuConfiguration() []uint16{
	return []uint16{
		 	5, 3, 4, 6, 7, 8, 9, 1, 2,
		 	6, 7, 2, 1, 9, 5, 3, 4, 8,
		 	1, 9, 8, 3, 4, 2, 5, 6, 7,
		 	8, 5, 9, 7, 8, 1, 4, 2, 3,
		 	4, 2, 6, 8, 5, 3, 7, 9, 1,
		 	7, 1, 3, 9, 2, 4, 8, 5, 6,
		 	9, 6, 1, 5, 3, 7, 2, 8, 4,
		 	2, 8, 7, 4, 1, 9, 6, 3, 5,
		 	3, 4, 5, 2, 8, 6, 1, 7, 9}
}

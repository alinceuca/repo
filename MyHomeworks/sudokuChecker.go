package main

import "fmt"

func main() {
	numbers := []uint16{
		5, 3, 4, 6, 7, 8, 9, 1, 2,
		6, 7, 2, 1, 9, 5, 3, 4, 8,
		1, 9, 8, 3, 4, 2, 5, 6, 7,
		8, 5, 9, 7, 6, 1, 4, 2, 3,
		4, 2, 6, 8, 5, 3, 7, 9, 1,
		7, 1, 3, 9, 2, 4, 8, 5, 6,
		9, 6, 1, 5, 3, 7, 2, 8, 4,
		2, 8, 7, 4, 1, 9, 6, 3, 5,
		3, 4, 5, 2, 8, 6, 1, 7, 9}

	if Check(numbers){
		fmt.Print("Correct configuration")
	} else {
		fmt.Print("Wrong configuration")
	}
}

func Check(numbers []uint16) bool {

	x := [3][9]uint16{}

	for index, value := range numbers {
		i := index / 9
		j := index % 9

		row := x[0][i]
		column := x[1][j]
		square := x[2][(3*(i/3))+j/3]
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

		x[0][i] = row | (1 << position)
		x[1][j] = column | (1 << position)
		x[2][(3 * (i / 3)) + j / 3] = square | (1 << position)
	}

	return true
}

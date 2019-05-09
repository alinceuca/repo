package main

import (
	"fmt"
)

func main() {
	fmt.Println("sunt in program")
	var i, j, nr int
	var max, min float32;
	var v [] float32
	fmt.Println("Nr puncte :")
	fmt.Scan(&nr)
	var matrix = [100][2]float32{}
	fmt.Println("Introduceti coordonatele punctelor x si y ")
	for i = 0; i < nr; i++ {
		for j = 0; j < 2; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	fmt.Println("Matricea de coordonate x|y este: ")
	for i = 0; i < nr; i++ {
		for j = 0; j < 2; j++ {
			fmt.Print(matrix[i][j], " | ")
		}
		fmt.Println()
	}

	for i = 0; i < nr; i++ {
		max = matrix[i][0];
		for j = 0; j < 2; j++ {
			if matrix[i][j] > max {
				max = matrix[i][j]
				v = append(v, max)
			}
		}

	}

	for i = 0; i < len(v); i++ {
		min = v[0]
		if v[i] < min {
			min = v[i]
		}
	}

	fmt.Print("minimul este ", min)

}
package main

import (
	"fmt"
)

func main() {
	var i, j, nr int
	var max, min float32;
	var v [] float32
	fmt.Println("Numar puncte :")
	fmt.Scan(&nr)
	var m = [100][2]float32{}
	fmt.Println("Introduceti coordonatele punctelor x si y ")
	for i = 0; i < nr; i++ {
		for j = 0; j < 2; j++ {
			fmt.Scan(&m[i][j])
		}
	}
	fmt.Println("Matricea de coordonate x,y este: ")
	for i = 0; i < nr; i++ {
		for j = 0; j < 2; j++ {
			fmt.Print(m[i][j], "   ")
		}
		fmt.Println()
	}

	for i = 0; i < nr; i++ {
		max = m[i][0];
		for j = 0; j < 2; j++ {
			if m[i][j] > max {
				max = m[i][j]
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
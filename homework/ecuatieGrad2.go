package main

import (
	"fmt"
	"math"
)

func calculeaza(a, b, c float64) {

	var delta, x1, x2, real, imag float64
	delta = b*b - 4*a*c

	if delta > 0 {
		x1 = (-b + math.Sqrt(delta)) / (2 * a)
		x2 = (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Print("x1 = ", x1, "  x2= ", x2)
	} else if delta == 0 {
		x1 = (-b / (2 * a))
		x2 = (-b / (2 * a))
		fmt.Print("x1 = ", x1, "  x2= ", x2)
	} else {
		real = -b / (2 * a)
		imag = math.Sqrt(-delta) / (2 * a)
		fmt.Println("x1 = ", real, "-", imag, "i    x2 = ", real, "+", imag, "i")

	}
}

func main(){

	var a,b,c float64
	fmt.Println("Citire coeficienti:")
	fmt.Println("a=")
	fmt.Scan(&a)
	fmt.Println("b=")
	fmt.Scan(&b)
	fmt.Println("c=")
	fmt.Scan(&c)
	calculeaza(a,b,c)
}

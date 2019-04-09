package main

import "fmt"

func main() {
	var x int
	var p, t uint

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&t)

	print(FindNumber(x, p, t))
}

func FindNumber(x int, p uint, t uint) int {
	x = x >> t
	x = x & (1<<p - 1)

	return x
}
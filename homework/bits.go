package main

func shift(nr int, pBiti uint, shiftT uint) int {
	nr = nr >> shiftT
	nr = nr & (1<<pBiti - 1)

	return nr
}

func main() {
	println("Shift", shift(5, 2, 1))
}

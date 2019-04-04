package main

import (
	"fmt"
)

func main() {

	var numar, pBiti, shiftatCuT uint
	print("Citeste numarul")
	fmt.Scan(&numar)
	print("Citeste p-ul ")
	fmt.Scan(&pBiti)
	print("Citeste t-ul ")
	fmt.Scan(&shiftatCuT)
	fmt.Print(shiftare(numar,pBiti,shiftatCuT))

}

func shiftare(numar, pBiti, shiftCuT uint) uint {

	aux := numar >> shiftCuT
	aux = aux & (1<<pBiti - 1)
	return aux

}


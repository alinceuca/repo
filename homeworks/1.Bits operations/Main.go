package main

import (
	"math"
	"strconv"
)

func main() {
	print(FindNumber(11, 3, 2))
	// am aflat mai tarziu ca problema a fost facuta in clasa :)
}

func FindNumber(x int, p uint, t uint) string {

	if float64(x) > math.Pow(2, float64(p))-1 {
		return "X has more than p bites"
	}
	var i uint
	for i = 1; i <= t; i++ {
		x = x & int(math.Pow(2, float64(p-i))-1)
	}
	return strconv.Itoa(x)
}

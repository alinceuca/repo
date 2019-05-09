package main

import "fmt"


func shiftare(x int, p uint, t uint) int {
	x = x >> t
	x = x & (1<<p - 1)

	return x
}

func main() {
	fmt.Println(shiftare(30,8,2))

}
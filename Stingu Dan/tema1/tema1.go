package main

import "fmt"

func main() {
	var x, t, p, nr uint
	fmt.Println("x=")
	fmt.Scan(&x)
	fmt.Println("t=")
	fmt.Scan(&t)
	fmt.Println("p=")
	fmt.Scan(&p)
	x = x >> t;
	nr = x & (1<<p - 1)
	fmt.Println("Rezultatul este", nr)
}

package main

import "fmt"

func main(){
	var x uint32
	fmt.Print("Valoare x:")
	fmt.Scan(&x)
	fmt.Println("Valoare t:")
	var t uint32
	fmt.Scan(&t)
	fmt.Println("Valoare p:")
	var p uint32
	fmt.Scan(&p)
	x = x >> t
	x = x&(1<<p-1)
	fmt.Println("Valoare finala x:",  x)



}

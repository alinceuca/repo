package main

import "fmt"

func main() {
	fmt.Println("main() has started")

	c := make(chan string)
	c <- "John"

	fmt.Println("main() has stopped")
}

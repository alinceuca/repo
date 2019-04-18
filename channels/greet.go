package main

import "fmt"

func greet(c chan string){
	fmt.Println("Hello " + <- c + " !!!")
}

func main() {
	fmt.Println("main() has started")

	c := make(chan string)

	go greet(c)

	c <- "John"

	fmt.Println("main() has stopped")

}

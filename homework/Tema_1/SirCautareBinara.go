package main

import "fmt"

func binarySearch(sir []int, value int) int {
	inceput := 0
	sfarsit := len(sir) - 1
	for inceput <= sfarsit {
		mijloc := (inceput + sfarsit) / 2
		if sir[mijloc] > value {
			sfarsit = mijloc - 1
		} else if sir[mijloc] < value {
			inceput = mijloc + 1
		} else {
			return mijloc
		}
	}
	return -1
}

func main(){
	sir := []int{0, 2, 3, 6, 9, 12, 16, 19, 22, 26}
	fmt.Println(sir)

	value := 9
	pozitie := binarySearch(sir, value)
	if pozitie == -1 {
		fmt.Println("Nu exista in sir")
	} else {
		fmt.Println("Gasit la pozitia ", pozitie)
	}
}

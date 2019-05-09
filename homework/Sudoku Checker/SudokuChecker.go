package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printConfiguration(configuration *[9][9]uint) {
	for i := 0; i < 9; i++ {
		fmt.Println(configuration[i])
	}
}

func checkLine(x *int16, poz uint) bool {
	if *x & (1 << poz) == 0 {
		*x = *x | (1 << poz)
	} else {
		fmt.Errorf("Not Valid")
		return false
	}
	return true
}

func checkColumn(x *int16, poz uint) bool {
	if *x & (1 << poz) == 0 {
		*x = *x | (1 << poz)
	} else {
		fmt.Errorf("Not Valid")
		return false
	}
	return true
}

func checkSquare(x *int16, poz uint) bool {
	if *x & (1 << poz) == 0 {
		*x = *x | (1 << poz)
	} else {
		fmt.Errorf("Not Valid")
		return false
	}
	return true
}

func main() {
	
	sudokuConfiguration, err := readLines("data")
	check(err)
	fmt.Println("Initial configuration:")
	printConfiguration(sudokuConfiguration)
	
	var a [3][9]int16

	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			poz := sudokuConfiguration[i][j]
			if checkLine(&a[0][i], poz) &&
			checkColumn(&a[1][j], poz) &&
			checkSquare(&a[2][(i/3)*3+j/3], poz) {
				fmt.Println("[",i,"][",j,"]:","Valid")
			} else {
				fmt.Println("[",i,"][",j,"]:","INVALID")
			}
		}
	}
}

/* Scanner for reading file */
// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (*[9][9]uint, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var lines [9][9]uint
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			scanner.Scan()
			number, err := strconv.Atoi(scanner.Text())
			check(err)
			lines[i][j] = uint(number)
		}
	}
	return &lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
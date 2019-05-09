package main
import (
	"fmt"
)

func main(){
	configuration:=getSudokuConfiguration()
	fmt.Printf(printBoard(configuration))
	generateSudokuSolution(&configuration)
	fmt.Printf("------------------SOLUTIN-----------------\n")
	fmt.Printf(printBoard(configuration))
}

func printBoard(board [9][9]int) string {
	var sudokuSolution string
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
			}
			sudokuSolution += fmt.Sprintf("%d ", board[row][col])
		}
		if row != 2 || row != 5 || row != 8 {
			sudokuSolution += fmt.Sprintln()
		}
	}
	return sudokuSolution
}

func generateSudokuSolution(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if generateSudokuSolution(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {

	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func getSudokuConfiguration() [9][9]int {
	return [9][9]int{
		{6, 3, 0, 5, 7, 0, 1, 8, 2},
		{5, 4, 1, 8, 2, 9, 3, 7, 6},
		{7, 8, 2, 6, 1, 3, 9, 5, 4},
		{1, 9, 8, 4, 0, 7, 5, 0, 3},
		{3, 0, 5, 9, 0, 2, 4, 0, 7},
		{4, 2, 7, 1, 3, 5, 8, 6, 9},
		{9, 5, 6, 7, 4, 8, 2, 3, 1},
		{8, 0, 3, 2, 9, 6, 7, 4, 0},
		{2, 7, 4, 3, 5, 1, 6, 9, 8},
	}
}
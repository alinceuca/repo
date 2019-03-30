//02. Sudoku - verificare configuratie folosind operatii pe biti

package main

import "fmt"

func main(){
	sudoku := [9][9]uint{
		{2,6,1,8,9,4,3,7,5},
		{3,4,5,2,7,1,6,8,9},
		{9,7,8,6,5,3,1,4,2},
		{1,5,6,3,4,2,7,9,8},
		{8,3,7,5,6,9,2,1,4},
		{4,9,2,1,8,7,5,3,6},
		{5,8,9,7,1,6,4,2,3},
		{7,2,4,9,3,5,8,6,1},
		{6,1,3,4,2,8,9,5,7},
	}
	if isValidConfiguration(sudoku){
		fmt.Println("Valid configuration!")
	}
}

func isValidConfiguration(sudoku [9][9]uint) bool{
	var a = [3][9]uint{}

	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			row := a[0][i]
			column := a[1][j]
			square := a[2][3*(i/3)+(j/3)]
			position := sudoku[i][j]
			if (row & (1 << position))==(1 << position){
				fmt.Println("Invalid configuration row")
				return false
			}else{
				a[0][i] = row | (1 << position)
			}
			if (column & (1 << position))==(1 << position){
				fmt.Println("Invalid configuration column")
				return false
			}else{
				a[1][j] = column | (1 << position)
			}
			if (square & (1 << position))==(1 << position){
				fmt.Println("Invalid configuration square")
				return false
			}else{
				a[2][3*(i/3)+(j/3)] = square | (1 << position)
			}
		}
	}
	return true
}

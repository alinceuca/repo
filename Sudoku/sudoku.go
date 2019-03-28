package main

import "fmt"

func main() {
	var s = [9][9]uint{
		{9,8,5,1,3,7,6,2,4},
		{1,2,7,4,9,6,5,3,8},
		{4,3,6,8,5,2,9,7,1},
		{5,4,9,3,6,1,2,8,7},
		{3,6,2,9,7,8,1,4,5},
		{8,7,1,5,2,4,3,9,6},
		{7,9,3,6,8,5,4,1,2},
		{2,5,4,7,1,3,8,6,9},
		{6,1,8,2,4,9,4,5,3},
	}
	x:=sudoku_bun(s)
	fmt.Println(x)
}

func sudoku_bun(s[9][9]uint) int {
	var a[3][9]uint
	for i:=0; i<9; i++ {
		for j:= 0; j < 9; j++ {
			p:=s[i][j]-1
			k:=3*(i/3)+(j/3)
			if a[0][i]&(1<<p)==0 {
				a[0][i] = a[0][i]|(1<< p)
			} else {
				return 0
				}
			if a[1][j]&(1<<p)==0 {
				a[1][j]=a[1][j]|(1<<p)
			} else {
				return 0
			}
			if a[2][k]&(1<<p) == 0 {
				a[2][k]=a[2][k]|(1<<p)
				} else {
					return 0
			}
			}

		}
	return 1
	}
	


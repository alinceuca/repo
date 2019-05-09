package main

import "fmt"

func main() {
	var s=[9][9] uint {
		{2,4,8,3,9,5,7,1,6},
		{5,7,1,6,2,8,3,4,9},
		{9,3,6,7,4,1,5,8,2},
		{6,8,2,5,3,9,1,7,4},
		{3,5,9,1,7,4,6,2,8},
		{7,1,4,8,6,2,9,5,3},
		{8,6,3,4,1,7,2,9,5},
		{1,9,5,2,8,6,4,3,7},
		{4,2,7,9,5,3,8,6,1},
	}

	fmt.Print(sudokuBun(s))
}

func sudokuBun(s[9][9] uint) bool{
	var p uint
	var a [9][9] uint
	for i:=0;i<=8;i++{
		for j:=0;j<=8;j++{
			p= s[i][j] - 1
			k:=3*(i/3)+(j/3)

			if a[0][i] & (1<<p) == 0 {
				a[0][i]=a[0][i]|(1<<p)
			}else{
				return false
			}

			if a[1][j] & (1<<p) == 0 {
				a[1][j]=a[1][j]|(1<<p)
			}else{
				return false
			}

			if a[2][k] & (1<<p) == 0 {
				a[2][k]=a[2][k]|(1<<p)
			}else{
				return false
			}
		}
	}
	return true
}
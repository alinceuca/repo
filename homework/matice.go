package main

import "fmt"

func main() {
	var col,lin int
	println("Numar de linii: ")
	fmt.Scanf("%d",&lin)
	println("Numar de coloane: ")
	fmt.Scanf("%d",&col)
	var matr[][] int
	matr=make([][]int,lin)
	citire_mat(matr, lin,col)
	//fmt.Println(matr)
	/*diag_1(matr)
	diag_2(matr)*/
	afisare_mat (matr)
	mat_spirala(matr)
}

func citire_mat(matr[][] int, lin,col int){
	var elem int
	for i:= range  matr{
		matr[i]=make([] int,col)

		for j:=range matr[i]{
			fmt.Printf("mat[%d][%d]: ",i,j)
			fmt.Scanf("%d",&elem)

			matr[i][j]=elem
			j++
		}
		i++
	}


}

func diag_1(matr[][] int){

	var i,j int
	for i= range matr{
		for j= range matr[i]{
			if i==j{
				fmt.Printf("mat[%d][%d]=%d ", i, j, matr[i][j])
			}


		}
	}
}

func diag_2(matr[][] int){
    var lin=len(matr)
    var col=coloane(matr)
	for i:=0;i<lin;i++{
		if i<col {
			fmt.Printf("mat[%d][%d]=%d ", lin-1-i, i, matr[lin-1-i][i])
		}
	}
}

func coloane(matr[][] int) int{
	var i,nrelem,col int

	for i= range matr{
		for _= range matr[i] {

			nrelem++
		}

	}

	col=nrelem/ len(matr)
	return col
}


func afisare_mat (matr[][] int) {
	var i int
	println(" Matricea: ")
	for i= range matr{
		fmt.Printf(" %d ",matr[i])
		fmt.Println(" ")
	}
	fmt.Println(" ")

}

func mat_spirala(matr[][] int){

	var maxlin,maxcol,maxelem,k int
	maxlin= len(matr)
	maxcol=coloane(matr)
	maxelem=maxlin*maxcol
	var i, j,m int
	for k<maxelem {
		j=m
		for j < maxcol-m {
			fmt.Printf(" %d ",matr[i][j])
			j++
			k++
			if k>maxelem{
				break
			}
		}
		if k>maxelem{
			break
		}
		j = maxcol - 1-m
		i = m+1
		for i < maxlin-m {
			fmt.Printf(" %d ",matr[i][j])
			i++
			k++
			if k==maxelem{
				break
			}
		}
		if k==maxelem{
			break
		}
		i = maxlin - 1-m
		for j > 0 {
			j--
			fmt.Printf(" %d ",matr[i][j])
			k++
			if k==maxelem{
				break
			}

		}
		if k==maxelem{
			break
		}
		j = m
		for i > 1 {
			i--
			fmt.Printf(" %d ",matr[i][j])
			k++
			if k==maxelem{
				break
			}

		}
		m++
	}

}
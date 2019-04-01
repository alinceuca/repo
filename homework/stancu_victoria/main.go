package main

import "fmt"

func main() {
	var n int
	println("Numar de elemente vector: ")
	fmt.Scanf("%d", &n)
	var a[]int
	a=make([]int,n,n)
	citire_vector(a)
	fmt.Println(a)
	//sort_1(a)
	quickSort(0,len(a)-1,a)
	fmt.Println(a)
}

func citire_vector(a[] int){
	var i,el int
    var n=len(a)
	for i<n{
		println("Element vector: ")
		fmt.Scanf( "%d",&el)
		a[i]=el
		i++
	}
}

func sort_1(a[] int){
	i:=0
	j:=0
	var temp int
	for len(a)-1>i {
		j=i+1
		for  j<len(a){
			if a[i] > a[j] {
				temp = a[i]
				a[i]=a[j]
				a[j]=temp
			}
			j++
		}
		i++
	}
}

func pivot(st int,dr int, a[] int) int {
	i:=st
	j:=dr
	var aux int
	mod:=1
	for i<j{
		if a[i]>a[j]{
			aux=a[i]
			a[i]=a[j]
			a[j]=aux
			mod=3-mod
		}
		if mod==1{
			j--
		}else{
				i++
		}
	}
	return i
}

func quickSort(st int,dr int, a[] int){
	var poz int
	if st<dr{
		poz=pivot(st,dr,a)
		quickSort(st,poz-1,a)
		quickSort(poz+1,dr,a)
	}
}
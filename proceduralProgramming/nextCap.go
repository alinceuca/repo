package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main(){
HerFuncCall()
//fmt.Println(MinimumInt1(5, 3), MinimumInt1(7, 3, -2, 4, 0, -8, -5))
//CallFibonacci()
fmt.Print("%v",IsPalindrome("ROTOR"))
}

func CallFibonacci(){
	for n := 0; n < 20; n++ {
		fmt.Print(Fibonacci(n), " ")
	}
}
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func HerFuncCall(){
	for i := 1; i <= 4; i++ {
		a, b, c := PythagoreanTriple(i, i+1)
		t1 := Heron(a, b, c)
		t2 := Heron(PythagoreanTriple(i, i+1))
		fmt.Printf("?1 == %10f == ?2 == %10f\n", t1, t2)
	}

}


func Heron(a, b, c int) float64 {
	a, ß, ? := float64(a), float64(b), float64(c)
	s := (a + ß + ?) / 2
	return math.Sqrt(s * (s - a) * (s - ß) * (s - ?))
}
func PythagoreanTriple(m, n int) (a, b, c int) {
	if m < n {
		m, n = n, m
	}
	return (m * m) - (n * n), (2 * m * n), (m * m) + (n * n)
}


func MinimumInt1(first int, rest ...int) int {
	for _, x := range rest {
		if x < first {
			first = x
		}
	}
	return first
}


func IsPalindrome(word string) bool {
	if utf8.RuneCountInString(word) <= 1 {
		return true
	}
	first, sizeOfFirst := utf8.DecodeRuneInString(word)
	last, sizeOfLast := utf8.DecodeLastRuneInString(word)
	if first != last {
		return false
	}
	return IsPalindrome(word[sizeOfFirst : len(word)-sizeOfLast])
}
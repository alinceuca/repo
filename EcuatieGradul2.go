package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, x, delta, x1, x2 float64
	fmt.Println("a=")
	fmt.Scanln(&a)
	fmt.Println("b=")
	fmt.Scanln(&b)
	fmt.Println("c=")
	fmt.Scanln(&c)

	if (a == 0){
		if (b == 0){
			if (c == 0){
				fmt.Println("Exista o infinitate de solutii")
			}else{
				fmt.Println("Ecuatie imposibila")
			}
		}else{
			fmt.Println("Ecuatie de gradul I cu solutia:")
			x = (-c) / b
			fmt.Println(x)
		}
	}else{
		delta=math.Pow(b,2)-(4*a*c)
		if(delta<0){
			fmt.Println("Ecuatia nu are solutii reale")
		}else{
			if(delta==0){
				fmt.Println("Ecuatia are solutii egale")
				fmt.Println("x1=x2=",(-b)/(2*a))
			}else{
				x1=((-b)+math.Sqrt(delta))/2*a
				x2=((-b)-math.Sqrt(delta))/2*a
				fmt.Println("Ecuatia are solutii reale")
				fmt.Println("x1=",x1)
				fmt.Println("x2=",x2)
			}
		}
	}
}
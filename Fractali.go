package main


import (
	"fmt"
	"math"
)

func main() {

	var segment float64
	var x1,y1,x2,y2 float64
	var xA,yA,xB,yB,xC,yC float64
	var lung float64
	var L float64

	fmt.Println("Introdu x1:")
	fmt.Scanln(&x1)
	fmt.Println("Introdu y1:")
	fmt.Scanln(&y1)
	fmt.Println("Introdu x2:")
	fmt.Scanln(&x2)
	fmt.Println("Introdu y2:")
	fmt.Scanln(&y2)

	segment=math.Sqrt(math.Pow((x2-x1),2)+math.Pow((y2-y1),2))
	lung=segment/3

	L=math.Asin((y2-y1)/3*L)

	xA=x1+lung*math.Sin(L)
	yA=y1+lung*math.Sin(L)

	xB=x1+lung*math.Sqrt(3)*math.Cos(L+math.Pi/6)
	yB=y1+lung*math.Sqrt(3)*math.Sin(L+math.Pi/6)

	xC=x1+2*lung*math.Cos(L)
	yC=y1+2*lung*math.Sin(L)

	fmt.Printf("segment: %f \n",segment);
	fmt.Printf("A: (%f , %f) \n",xA,yA);
	fmt.Printf("B: (%f , %f) \n",xB,yB);
	fmt.Printf("C: (%f , %f) \n",xC,yC);

}


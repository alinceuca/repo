package main

import (
	"bufio"
	"fmt"
	drawing "github.com/StephaneBunel/bresenham"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

func floatReader() float64{
	reader := bufio.NewReader(os.Stdin)
	floatText, _ := reader.ReadString('\n')
	floatText = strings.Replace(floatText, "\n", "", -1)
	floatNr, _ := strconv.ParseFloat(floatText, 64)
	if floatNr == 0 && len(floatText) != 1{
		fmt.Println("Introduceti un numar valid!")
		return floatReader()
	}
	return floatNr
}

func intReader() int{
	reader := bufio.NewReader(os.Stdin)
	intText, _ := reader.ReadString('\n')
	intText = strings.Replace(intText, "\n", "", -1)
	intNr, _ := strconv.ParseInt(intText, 10, 64)
	if intNr == 0 && len(intText) != 1{
		fmt.Println("Introduceti un numar valid!")
		return intReader()
	}
	return int(intNr)
}

type point struct {
	x int
	y int
}

func pointReader() point{
	var p point
	fmt.Print("X: ")
	p.x = intReader()
	for p.x < 0{
		fmt.Println("Va rugam sa introduceti valori pozitive!")
		fmt.Print("X: ")
		p.x = intReader()
	}
	fmt.Print("Y: ")
	p.y = intReader()
	for p.y < 0{
		fmt.Println("Va rugam sa introduceti valori pozitive!")
		fmt.Print("Y: ")
		p.y = intReader()
	}
	return p
}

func orderReader() int{
	fmt.Print("Introduceti ordinul: ")
	n := intReader()
	for n < 0 {
		fmt.Println("Nu se accepta valori negative!")
		fmt.Print("Ordinul: ")
		n = intReader()
	}
	return n
}

var curbaKoch = image.NewRGBA(image.Rect(0, 0, 1000, 1000))
var fulg = image.NewRGBA(image.Rect(0, 0, 1000, 1000))

func main() {
	draw.Draw(curbaKoch, curbaKoch.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	draw.Draw(fulg, fulg.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	fmt.Println("Va rugam sa introduceti 2 puncte diferite!")
	fmt.Println("T1!")
	T1 := pointReader()
	fmt.Println("T2!")
	T2 := pointReader()
	for T1.x == T2.x && T1.y == T2.y {
		fmt.Println("Va rugam sa introduceti 2 puncte diferite!")
		fmt.Println("T1!")
		T1 = pointReader()
		fmt.Println("T2!")
		T2 = pointReader()
	}

	order := orderReader()

	KochCurve(order,T1.x,T1.y,T2.x,T2.y,false,curbaKoch)
	f, err := os.Create("curbakoch.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, curbaKoch)

	KochSnowflake(order,T1.x,T1.y,T2.x,T2.y,true)
	f, err = os.Create("fulg.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, fulg)
}

func KochCurve (order,x1,y1,x2,y2 int,verif bool,img draw.Image) {
	if order == 0 {
		drawing.Bresenham(img,x1,y1,x2,y2,color.RGBA{255,0,0,255})
	} else {
		if x1 > x2 {
			x1, y1, x2, y2 = x2, y2, x1, y1
			if verif == false {
				verif = true
			} else {
				verif = false
			}
		}
		T1T2 := math.Sqrt(math.Pow(float64(x2-x1),2)+math.Pow(float64(y2-y1),2))
		L := T1T2 / 3
		alpha := math.Asin(float64(y2-y1)/T1T2)
		xa := math.Round(float64(x1) + L*math.Cos(alpha))
		ya := math.Round(float64(y1) + L*math.Sin(alpha))
		var xb,yb float64
		if verif == false {
			xb = math.Round(float64(x1) + L*math.Sqrt(3)*math.Cos(alpha - (math.Pi)/6))
			yb = math.Round(float64(y1) + L*math.Sqrt(3)*math.Sin(alpha - (math.Pi)/6))
		} else {
			xb = math.Round(float64(x1) + L*math.Sqrt(3)*math.Cos(alpha + (math.Pi)/6))
			yb = math.Round(float64(y1) + L*math.Sqrt(3)*math.Sin(alpha + (math.Pi)/6))
		}
		xc := math.Round(float64(x1) + 2*L*math.Cos(alpha))
		yc := math.Round(float64(y1) + 2*L*math.Sin(alpha))
		order--
		KochCurve(order,x1,y1,int(xa),int(ya), verif, img)
		KochCurve(order,int(xa),int(ya),int(xb),int(yb), verif, img)
		KochCurve(order,int(xb),int(yb),int(xc),int(yc), verif, img)
		KochCurve(order,int(xc),int(yc),x2,y2, verif, img)
	}
}

func KochSnowflake(order,x1,y1,x2,y2 int, verif bool) {
	x3 := math.Round((float64(x1+x2)+math.Sqrt(3)*float64(y2-y1))/2)
	y3 := math.Round((float64(y1+y2)-math.Sqrt(3)*float64(x2-x1))/2)
	KochCurve(order,x1,y1,x2,y2,verif,fulg)
	KochCurve(order,x2,y2,int(x3),int(y3),verif,fulg)
	KochCurve(order,int(x3),int(y3),x1,y1,verif,fulg)
}
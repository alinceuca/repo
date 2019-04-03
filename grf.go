package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

// fractali plecand de la un pattern: curba lui Koch, triunghi echilateral
// desenarea la un pas se va face cu o culoare data



var img = image.NewRGBA(image.Rect(0, 0, 100, 100))
var col color.Color

func main() {

	var segment float64
	var x1,y1,x2,y2 float64
	var xA,yA,xB,yB,xC,yC float64
	var lung float64
	var L float64

	fmt.Println("Introdu coordonatele lui x1: ")
	fmt.Scanln(&x1)
	fmt.Println("Introdu coordonatele lui y1: ")
	fmt.Scanln(&y1)
	fmt.Println("Introdu coordonatele lui x2: ")
	fmt.Scanln(&x2)
	fmt.Println("Introdu coordonatele lui y2: ")
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

	fmt.Printf("Lungimea segmentului xy este: %f \n",segment);
	fmt.Printf("Coordonatele lui A: (%f , %f) \n",xA,yA);
	fmt.Printf("Coordonatele lui B: (%f , %f) \n",xB,yB);
	fmt.Printf("Coordonatele lui C: (%f , %f) \n",xC,yC);



	col = color.RGBA{255, 0, 0, 255} // Red
	Bresenham(img,int(x1),int(x2),int(xA),int(xB),col)

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

// Generalized with integer
func Bresenham(img draw.Image, x1, y1, x2, y2 int, col color.Color) {
	var dx, dy, e, slope int

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		img.Set(x1, y1, col)

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			img.Set(x1, y1, col)
			x1++
		}
		img.Set(x1, y1, col)

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			img.Set(x1, y1, col)
			y1++
		}
		img.Set(x1, y1, col)

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				y1--
			}
		}
		img.Set(x1, y1, col)

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			// BresenhamDxXRYD(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.Set(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		img.Set(x2, y2, col)

	// higher than wide.
	default:
		if y1 < y2 {
			// BresenhamDyXRYD(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.Set(x1, y1, col)
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.Set(x1, y1, col)
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		img.Set(x2, y2, col)
	}
}
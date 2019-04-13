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

var img = image.NewRGBA(image.Rect(0, 0, 4000, 4000))
var col = color.RGBA{0, 125, 255, 255} // Red

func main() {
	var numarDeIteratii int
	var x1, y1, x2, y2 float64
	fmt.Print("Insert the number of recursion \n")
	fmt.Scan(&numarDeIteratii)
	x1 = 500
	y1 = 3000
	x2 = 3500
	y2 = 3000

	Koch(x1, y1, x2, y2, numarDeIteratii, false)
	f, err := os.Create("KochCurve.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)

	img = image.NewRGBA(image.Rect(0, 0, 4000, 4000))
	snowFlake(x1, y1, x2, y2, numarDeIteratii, true)
	f, err = os.Create("snowflake.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)

}
func Koch(x1, y1, x2, y2 float64, n int, test bool) {
	if n == 0 {
		Bresenham(img, int(x1), int(y1), int(x2), int(y2), col)
	} else {
		if n > 0 {
			var xa, ya, xb, yb, xc, yc float64

			if x1 > x2 {
				x1, y1, x2, y2 = x2, y2, x1, y1
				if test == false {
					test = true
				} else {
					test = false
				}
			}
			segmentLenght := (math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))) / 3

			alfa := math.Asin((y2 - y1) / (3 * segmentLenght))
			xa = x1 + segmentLenght*math.Cos(alfa)
			ya = y1 + segmentLenght*math.Sin(alfa)

			if test == true {
				xb = x1 + segmentLenght*math.Sqrt(3)*math.Cos(alfa+(math.Pi/6))
				yb = y1 + segmentLenght*math.Sqrt(3)*math.Sin(alfa+(math.Pi/6))
			} else {
				xb = x1 + segmentLenght*math.Sqrt(3)*math.Cos(alfa-(math.Pi/6))
				yb = y1 + segmentLenght*math.Sqrt(3)*math.Sin(alfa-(math.Pi/6))
			}

			xc = x1 + 2*segmentLenght*math.Cos(alfa)
			yc = y1 + 2*segmentLenght*math.Sin(alfa)

			Bresenham(img, int(x1), int(y1), int(xa), int(ya), col)
			Bresenham(img, int(xa), int(ya), int(xb), int(yb), col)
			Bresenham(img, int(xb), int(yb), int(xc), int(yc), col)
			Bresenham(img, int(xc), int(yc), int(x2), int(y2), col)
			n--
			Koch(x1, y1, xa, ya, n, test)
			Koch(xa, ya, xb, yb, n, test)
			Koch(xb, yb, xc, yc, n, test)
			Koch(xc, yc, x2, y2, n, test)
		} else if n < 0 {
			fmt.Println("Illegal number")
			os.Exit(3)
		}
	}

}

func snowFlake(x1, y1, x2, y2 float64, n int, test bool) {
	x3 := (x1 + x2 + math.Sqrt(3)*(y2-y1)) / 2
	y3 := (y1 + y2 - math.Sqrt(3)*(x2-x1)) / 2
	Koch(x1, y1, x2, y2, n, test)
	Koch(x2, y2, x3, y3, n, test)
	Koch(x3, y3, x1, y1, n, test)
}

// HLine draws a horizontal line
func HLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
	HLine(x1, y1, x2)
	HLine(x1, y2, x2)
	VLine(x1, y1, y2)
	VLine(x2, y1, y2)
}

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

	// Is line a point
	case x1 == x2 && y1 == y2:
		img.Set(x1, y1, col)

		// Is line an horizontal
	case y1 == y2:
		for ; dx != 0; dx-- {
			img.Set(x1, y1, col)
			x1++
		}
		img.Set(x1, y1, col)

		// Is line a vertical
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			img.Set(x1, y1, col)
			y1++
		}
		img.Set(x1, y1, col)

		// Is line a diagonal
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

		// wider than high
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

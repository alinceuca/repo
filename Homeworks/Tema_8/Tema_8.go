package main

import (
	"github.com/fogleman/gg"
	"math"
	"math/rand"
)

type Point struct {
	x float64
	y float64
}

//func modulo(n float64) float64 {
//	if(n > 0){
//		return n
//	}
//	return -n
//}

func getPoints(a, b Point) []Point {
	var points []Point

	u := Point{b.x - a.x, b.y - a.y}
	v := Point{a.y - b.y, b.x - a.x}

	m := (b.y - a.y) / (b.x - a.x)

	points = append(points, Point{a.x + u.x/3, a.y + u.y/3})

	if m != 0 {
		points = append(points, Point{a.x + u.x/2 + math.Sqrt(3)/6*v.x,
			a.y + u.y/2 + math.Sqrt(3)/6*v.y})
	} else {
		points = append(points, Point{a.x + u.x/2 + math.Sqrt(3)/6*v.x,
			a.y + u.y/2 - math.Sqrt(3)/6*v.y})
	}

	points = append(points, Point{a.x + u.x*2/3, a.y + u.y*2/3})

	return points
}

func drawPoints(points []Point, dc *gg.Context) {
	for i := 0; i < len(points)-1; i++ {

		color := rand.Float64()
		if color < 0.33 {
			dc.SetRGB(color, 0, 0)
		} else if color < 0.66 {
			dc.SetRGB(0, color, 0)
		} else {
			dc.SetRGB(0, 0, color)
		}

		dc.DrawLine(points[i].x, points[i].y, points[i+1].x, points[i+1].y)
		dc.Stroke()
	}

	dc.SetRGBA(0, 0, 0, 0)
	dc.DrawLine(points[1].x, points[1].y, points[3].x, points[3].y)
	dc.Stroke()
}

//
//func getLength(a Point, b Point) float64{
//	return math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2))
//}
//
//func getAlpha(a, b Point, length float64) float64{
//	var sin = float64(b.y-a.y)/length
//	return math.Asin(sin)
//}

func getFractal(a, b Point, n, i int, dc *gg.Context) {
	if i >= n {
		return
	}

	var points []Point

	points = append(points, a)
	points = append(points, getPoints(a, b)...)
	points = append(points, b)

	drawPoints(points, dc)

	getFractal(points[0], points[1], n, i+1, dc)
	getFractal(points[1], points[2], n, i+1, dc)
	getFractal(points[2], points[3], n, i+1, dc)
	getFractal(points[3], points[4], n, i+1, dc)
}

func main() {
	a := Point{300, 400}
	b := Point{600, 800}
	c := Point{900, 400}

	dc := gg.NewContext(1200, 1200)
	dc.SetLineWidth(1)

	getFractal(a, b, 5, 0, dc)
	getFractal(b, c, 5, 0, dc)
	getFractal(a, c, 5, 0, dc)

	dc.SavePNG("./Homeworks/Tema_8/out.png")
}

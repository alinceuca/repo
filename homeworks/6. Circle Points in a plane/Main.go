package main

import "fmt"
import "math"
import "math/rand"
import "time"

var r1 *rand.Rand

func main() {
	r1 = rand.New(rand.NewSource(time.Now().UnixNano()))

	dotList := getListOfRandomDots(1000)
	farthestDotPair := getFarthestDotPair(dotList)
	fmt.Println("Circle center point: ", getMiddlePoint(farthestDotPair.firstDot, farthestDotPair.secondDot))
	fmt.Println("Circle radius: ", farthestDotPair.distance/2)
}

func getListOfRandomDots(n int) []Dot {
	dotList := []Dot{}
	for i := 0; i < n; i++ {
		dotList = append(dotList, getRandomDot())
	}

	return dotList
}

func getRandomDot() Dot {
	return Dot{
		x: r1.Float64() * 1000,
		y: r1.Float64() * 1000,
	}
}

func getDotPair(firstDot, secondDot Dot) DotPair {
	return DotPair{
		firstDot:  firstDot,
		secondDot: secondDot,
		distance:  getDistanceBetweenDots(firstDot, secondDot),
	}
}

func getDistanceBetweenDots(firstDot, secondDot Dot) float64 {
	xPart := math.Pow(firstDot.x-secondDot.x, 2)
	yPart := math.Pow(firstDot.y-secondDot.y, 2)
	return math.Pow(xPart+yPart, 0.5)
}

func getFarthestDotPair(dotList []Dot) DotPair {
	farthestDotPair := getDotPair(dotList[0], dotList[1])
	for _, firstDot := range dotList {
		for _, secondDot := range dotList {
			dotPair := getDotPair(firstDot, secondDot)
			if farthestDotPair.distance < dotPair.distance {
				farthestDotPair = dotPair
			}
		}
	}
	return farthestDotPair
}

func getMiddlePoint(firstDot, secondDot Dot) Dot {
	return Dot{
		x: (firstDot.x + secondDot.x) / 2,
		y: (firstDot.y + secondDot.y) / 2,
	}
}

type Dot struct {
	x float64
	y float64
}

type DotPair struct {
	firstDot  Dot
	secondDot Dot
	distance  float64
}

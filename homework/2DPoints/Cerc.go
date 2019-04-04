package main

import (
	"fmt"
	"math"
)

type distanceIndex struct {
	value float64
	index int32
}

func main(){

	var points []Point
	var nrOfPoints int32
	fmt.Print("Numar puncte: ")
	fmt.Scan(&nrOfPoints)
	var matrix [nrOfPoints][nrOfPoints]float64
	points = readPoints(nrOfPoints)

	for i = 0; i < nrOfPoints ;  i++{
		for j = 0; j < nrOfPoints ; j++  {
			matrix[i][j] = distance(points[i], points[j])
			matrix[j][i] = matrix[i][j]
		}
	}

	//var rez = findValue(matrix)
	//fmt.Println(rez)
}

func readPoints(numberOfPoints int32) []Point{
	var x, y float64
	var points []Point
	for i = 0; i < numberOfPoints; i++{
		fmt.Println("Punct %d", i)
		fmt.Println("x: ")
		fmt.Scan(&x)
		fmt.Println("y: ")
		fmt.Scan(&y)
		points = append(points, Point{X:x, Y:y})
	}
	return points
}

func distance(point1, point2 Point,) float64{
	var dist float64
	dist = math.Sqrt(math.Pow(point1.X - point2.X, 2) + math.Pow(point1.Y - point2.Y, 2))
	return dist
}

func findValue(matrix [][]float64) distanceIndex{
	var maxValues []distanceIndex
	var maxValue distanceIndex
	var minValue distanceIndex

	for i = 0; i < len(matrix) - 1; i++ {
		for j = 0; j < len(matrix); j++{
			if matrix[i][j] > maxValue.value {
				maxValue.value = matrix[i][j]
				maxValue.index = j
			}
		}
		maxValues = append(maxValues, maxValue)
	}
	for i = 0; i < len(maxValues); i++{
		if maxValues[i].value < minValue.value{
			minValue = maxValues[i]
		}
	}
	return minValue
}
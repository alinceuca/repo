package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func saveMatrix(n int, pointsVector []point) [][]float64 {

	distanceVector := make([][] float64, n)
	//var distanceVector [][]float64
	println(n)
	var name float64 = 0
	println(name)
	for i := 0; i < n; i++ {
		distanceVector[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if (i == j) {
				distanceVector[i][j] = name
			} else {
				distanceVector[i][j] = calculateDistance(pointsVector[i].x, pointsVector[i].y, pointsVector[j].x, pointsVector[j].y)
			}
		}
	}
	return distanceVector
}

func saveTheMaximValue(matriceInitiala [][]float64, n int) []float64 {

	var listaDeMaxime = make([]float64, n)
	var auxMax float64
	for i := 0; i < n; i++ {
		auxMax = matriceInitiala[i][0]
		for j := 1; j < n-1; j++ {
			if matriceInitiala[i][j] > auxMax {
				auxMax = matriceInitiala[i][j]
			}
		}
		listaDeMaxime[i] = auxMax
		println(listaDeMaxime[i])
	}
	return listaDeMaxime

}

func saveTheMinimValueFromMaxArray(n int, maxList [] float64) float64 {

	var minim float64 = 1000000
	for index := 0; index < len(maxList); index++ {
		if (maxList[index] < minim) {
			minim = maxList[index]
		}
	}
	return minim
}

func calculateDistance(distance1 float64, distance2 float64, distance3 float64, distance4 float64) float64 {
	var distanceResult float64
	distanceResult = math.Sqrt((distance4-distance2)*(distance4-distance2) + (distance3-distance1)*(distance3-distance1))
	return distanceResult
}

func main() {
	var points = []point{
		{x: 1, y: 3},
		{x: 3, y: 2},
		{x: 1, y: 4},
		{x: 1, y: 6},
		{x: 3, y: 3},
	}

	matrix := make([][] float64, len(points));
	for i := 0; i < len(points); i++ {
		matrix[i] = make([]float64, len(points));
	}
	var maxArray []float64
	var radius float64
	matrix = saveMatrix(len(points), points)
	maxArray = saveTheMaximValue(matrix, len(matrix))
	radius = saveTheMinimValueFromMaxArray(len(points), maxArray)
	println(radius)
	var convertToInt int = int(radius)
	println("Centrul cercului", radius)
	fmt.Println(points)
	fmt.Println("Centrul cercului se afla la coordonatele:",points[convertToInt])

}

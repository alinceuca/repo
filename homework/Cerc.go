package main

import (
	"fmt"
	"math"
)
type punct struct {
	x float64
	y float64
}
func main() {
	var points = []punct{{x: 3, y: 5}, {x: 1, y: 2}, {x: 6, y: 4}, {x: 1, y: 2}, {x: 4, y: 3}}
	matrix := make([][] float64, len(points));
	for i := 0; i < len(points); i++ {
		matrix[i] = make([]float64, len(points));
	}
	var maxArray []float64
	var radius float64
	matrix = saveMatrix(len(points), points)
	maxArray = saveTheMaximValue(matrix, len(matrix))
	radius = saveTheMinimValueFromMaxArray(len(points), maxArray)
	fmt.Println("Centrul cercului se afla la coordonatele:",points[int(radius)])
	fmt.Println("Cu raza:", int(radius))
}
func saveMatrix(n int, pointsVector []punct) [][]float64 {
	distanceVector := make([][] float64, n)
	var aux float64 = 0
	for i := 0; i < n; i++ {
		distanceVector[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if (i == j) {
				distanceVector[i][j] = aux
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




//06. Se dau n puncte in plan prin coordonatele lor. Sa se determine cercul de raza minima si cu centrul in unul din punctele date care contine in interior toate punctele.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point2D struct{
	x float64
	y float64
}

type minOrMax int
const (
	min minOrMax = 0
	max minOrMax = 1
)

func main(){

	reader := bufio.NewReader(os.Stdin)
	nFloat, err := readAndValidateFloat(reader,"Number of points: ")

	if err != nil{
		fmt.Println(err)
	} else {
		n := int(nFloat)
		if n>0 {

			distances, points := readPoints(reader, n)
			if (distances != nil) {
				makeDistancesMatrix(points,*distances)

				//fmt.Println(distances)
				radius := getRadius(*distances)
				fmt.Printf("Radius = %f\n", radius)
			}
		} else {
			fmt.Println("The number of points is less or equal to 0")
		}
	}
}

func readAndValidateFloat(reader *bufio.Reader, message string) (float64, error) {
	fmt.Print(message)
	numberStr, _ := reader.ReadString('\n')
	numberStr = strings.Replace(numberStr, "\n", "", -1)
	number, err := strconv.ParseFloat(numberStr, 64)

	return number, err
}

func readPoints(reader *bufio.Reader, n int) (*[][]float64,[]point2D){
	var points []point2D
	point := point2D{}
	distances := make([][]float64, n)

	for i := 0; i < n; i++ {
		distances[i] = make([]float64, n)

		fmt.Printf("Point %d :\n", (i + 1))
		x, errX := readAndValidateFloat(reader, "x= ")

		if errX != nil {
			fmt.Println(errX)
			return nil, nil
		} else {
			y, errY := readAndValidateFloat(reader, "y= ")
			if errY != nil {
				fmt.Println(errY)
				return nil, nil
			} else {
				point.x = x
				point.y = y
				points = append(points, point)

			}
		}
	}
	return &distances, points
}

func makeDistancesMatrix(points []point2D, distances [][]float64){
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			distances[i][j] = math.Sqrt(math.Pow(points[i].x-points[j].x, 2) + math.Pow(points[i].y-points[j].y, 2))
			distances[j][i] = distances[i][j]
		}
	}

}

func getRadius(distances [][]float64) float64{
	arrayMax := make([]float64, len(distances))
	for i := 0; i < len(distances); i++ {
		arrayMax[i] = getMinOrMaxFromArray(distances[i], max)
	}
	radius := getMinOrMaxFromArray(arrayMax, min)
	return radius
}

func getMinOrMaxFromArray (array []float64, operation minOrMax) float64{
	result := array[0]
	for _, value := range array {
		if operation == max {
			if value > result {
				result = value
			}
		} else {
			if value < result{
				result = value
			}
		}
	}
	return result
}
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int64
	y int64
}

func intReader() int64{
	reader := bufio.NewReader(os.Stdin)
	intText, _ := reader.ReadString('\n')
	intText = strings.Replace(intText, "\n", "", -1)
	intNr, _ := strconv.ParseInt(intText, 10, 64)
	if(intNr == 0 && len(intText) != 1){
		fmt.Println("Introduceti un numar valid!")
		return intReader()
	}
	return intNr
}

func nReader() int64{
	fmt.Println("Introduceti numarul de puncte:")
	n := intReader()
	return n
}

func pointReader(i int64) point{
	fmt.Println("Punctul ",i)
	var p point
	fmt.Print("X: ")
	p.x = intReader()
	fmt.Print("Y: ")
	p.y = intReader()
	return p
}

func pointsReader(n int64) []point{
	var points []point
	var i int64
	for i = 0; i < n ; i++{
		p := pointReader(i+1)
		points = append(points, p)
	}
	return points
}

func removePoint(points []point, index int64) []point{
	return append(points[:index], points[index+1:]...)
}

func removeDuplicates(n int64, points []point) (int64,[]point){
	var i,j int64
	for i = 0; i < n - 1; i++{
		point := points[i]
		j = i+1
		for j<n{
			if point.x == points[j].x && point.y == points[j].y{
				points = removePoint(points,j)
				n--
			} else {
				j++
			}
		}
	}
	return n,points
}

func minRadius(n int64, points []point){
	if n<2{
		fmt.Println("The number of points is lower than 2!")
		fmt.Println("Can't compute the minimum radius!")
		os.Exit(0)
	}
	var i,j int64
	var max,min float64
	var firstCheck,lastCheck []float64
	for i = 0; i < n; i++{
		for j = 0; j < n; j++{
			if i != j {
				firstCheck = append(firstCheck, math.Sqrt(math.Pow(float64(points[i].x-points[j].x),2)+math.Pow(float64(points[i].y-points[j].y),2)))
			}
		}
		max = firstCheck[0]
		for j = 1; j < int64(len(firstCheck)); j++{
			if max < firstCheck[j] {
				max = firstCheck[j]
			}
		}
		firstCheck = nil
		lastCheck = append(lastCheck,max)
	}
	min = lastCheck[0]
	for i = 1; i < int64(len(lastCheck)); i++{
		if min > lastCheck[i] {
			min = lastCheck[i]
		}
	}
	fmt.Println("Minimum radius:",min)
}

func main(){
	n := nReader()
	points := pointsReader(n)
	n,points = removeDuplicates(n,points)
	minRadius(n,points)
}

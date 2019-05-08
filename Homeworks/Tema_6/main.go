package main

import (
	"flag"
	"github.com/fogleman/gg"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Point struct {
	Xaxis          float64
	Yaxis          float64
	XaxisLabelName string
	YaxisLabelName string
}

type PageData struct {
	ErrorMessage  string
	NrOfPoints    int
	PageTitle     string
	Points        []Point
	PerfectPointX float64
	PerfectPointY float64
	Radius        float64
}

func CalculateDistanceMatrix(p []Point) [][]float64 {
	n := len(p)
	d := make([][]float64, n)

	for i := 0; i < n; i++ {
		d[i] = make([]float64, n+1)
		for j := 0; j < n; j++ {
			x1 := float64(p[i].Xaxis)
			y1 := float64(p[i].Yaxis)

			x2 := float64(p[j].Xaxis)
			y2 := float64(p[j].Yaxis)

			d[i][j] = math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
		}
	}

	return d
}

func CreateVisualisation(p []Point, perfectPoint Point, r float64) {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(perfectPoint.Xaxis, perfectPoint.Yaxis, r)
	dc.SetRGBA(120, 0, 0, 0.5)
	dc.Fill()

	for i := 0; i < len(p); i++ {
		dc.DrawCircle(p[i].Xaxis, p[i].Yaxis, 5)
	}

	dc.SetRGBA(0, 0, 0, 0.5)
	dc.Fill()
	dc.SavePNG("./statics/points.png")
}

func GetPerfectPointIndex(d [][]float64, p []Point) int {
	n := len(p)

	for i := 0; i < n; i++ {
		d[i][n] = -1
		for j := 0; j < n; j++ {
			if d[i][n] < d[i][j] {
				d[i][n] = d[i][j]
			}
		}
	}

	min := d[0][n]
	index := 0

	for i := 1; i < n; i++ {
		if min > d[i][n] {
			min = d[i][n]
			index = i
		}
	}

	return index
}

func PointsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	inputVal := r.FormValue("NrOfPoints")
	nrOfPoints, err := strconv.Atoi(inputVal)
	points := r.FormValue("Points")

	print(points)

	var data PageData

	if err == nil && nrOfPoints > 0 {
		var points []Point
		var perfectPointX float64
		var perfectPointY float64
		var radius float64

		for i := 0; i < nrOfPoints; i++ {
			xlabel := "Xaxis" + strconv.Itoa(i)
			ylabel := "Yaxis" + strconv.Itoa(i)

			xval, err1 := strconv.Atoi(r.FormValue(xlabel))
			yval, err2 := strconv.Atoi(r.FormValue(ylabel))

			if err1 == nil && err2 == nil {
				points = append(points, Point{
					Xaxis:          float64(xval),
					Yaxis:          float64(yval),
					XaxisLabelName: "Xaxis" + strconv.Itoa(i),
					YaxisLabelName: "Yaxis" + strconv.Itoa(i)})

			} else {
				points = append(points, Point{
					Xaxis: 0,

					Yaxis:          0,
					XaxisLabelName: "Xaxis" + strconv.Itoa(i),
					YaxisLabelName: "Yaxis" + strconv.Itoa(i)})
			}
		}

		if len(points) > 0 {
			dMatrix := CalculateDistanceMatrix(points)
			index := GetPerfectPointIndex(dMatrix, points)

			perfectPointX = points[index].Xaxis
			perfectPointY = points[index].Yaxis
			radius = dMatrix[index][len(points)]
		}

		data = PageData{
			PerfectPointX: perfectPointX,
			PerfectPointY: perfectPointY,
			NrOfPoints:    nrOfPoints,
			PageTitle:     "Points app",
			Points:        points,
			Radius:        radius,
		}

	} else {
		data = PageData{
			PageTitle:    "Points app",
			ErrorMessage: "Ati introdus un numar invalid!",
		}
	}

	CreateVisualisation(data.Points, Point{data.PerfectPointX, data.PerfectPointY, "", ""}, data.Radius)

	tmpl.Execute(w, data)
}

func main() {
	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", "./statics", "the directory of static file to host")
	flag.Parse()

	http.HandleFunc("/points", PointsHandler)

	http.Handle("/statics/", http.StripPrefix(strings.TrimRight("/statics/", "/"), http.FileServer(http.Dir(*directory))))

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

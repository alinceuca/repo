package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

var (
	a float64
	b float64
	c float64
)

func main() {
	tmpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		a, _ := strconv.ParseFloat(r.FormValue("aCoef"),64)
		b, _ := strconv.ParseFloat(r.FormValue("bCoef"),64)
		c, _ := strconv.ParseFloat(r.FormValue("cCoef"),64)
		solutions := make(chan float64, 2)
		fmt.Fprintf(w, "%s", solve(a, b, c, solutions))
	})
	http.ListenAndServe(":8080", nil)
}

func solve(a float64, b float64, c float64, solutions chan float64) string{
	var div float64
	var quo float64

	div = (b * -1) + math.Sqrt((math.Pow(b, 2))-(4*a*c))
	quo = div / (2 * a)

	solutions <- quo
	delta := (math.Pow(b, 2))-(4*a*c)

	if delta > 0 {
		x1 := ((-b) + math.Sqrt(delta)) / (2 * a)
		x2 := ((-b) - math.Sqrt(delta)) / (2 * a)
		return fmt.Sprintf("x1: %f  x2: %f", x1, x2)
	}
	if delta == 0 {
		x := (-b) / (2 * a)
		return fmt.Sprintf("x: %f", x)
	}
	if delta < 0 {
		r := (-b) / (2 * a)
		i := math.Sqrt(-delta) / (2 * a)
		var x1 = complex(r, i)
		var x2 = complex(r, -i)
		return fmt.Sprintf("x1: %v  x2: %v", x1, x2)
	}
	return ""
}
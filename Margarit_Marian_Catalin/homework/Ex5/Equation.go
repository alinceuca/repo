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
		go s1(a, b, c, solutions)
		go s2(a, b, c, solutions)
		fmt.Fprintf(w, "X1= %f\nX2=%f", <-solutions,<-solutions)
	})
	http.ListenAndServe(":8080", nil)
}

func s1(a float64, b float64, c float64, solutions chan float64) {
	var div float64
	var quo float64

	div = (b * -1) + math.Sqrt((math.Pow(b, 2))-(4*a*c))
	quo = div / (2 * a)

	solutions <- quo
}

func s2(a float64, b float64, c float64, solutions chan float64) {
	var div float64
	var quo float64

	div = (b * -1) - math.Sqrt((math.Pow(b, 2))-(4*a*c))
	quo = div / (2 * a)

	solutions <- quo
}
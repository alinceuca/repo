//05. Rezolvati ecuatia de gradul 2 folosind(web app)

package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Quadratic struct {
	a float64
	b float64
	c float64
	d float64
}

func main() {
	http.HandleFunc("/quadratic/", quadraticHandler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func solveQuadraticEquation(quadratic Quadratic) string {
	a := quadratic.a
	b := quadratic.b
	c := quadratic.c - quadratic.d

	delta := b*b - 4*a*c

	if delta > 0 {
		x1 := ((-b) + math.Sqrt(b*b-4*a*c)) / (2 * a)
		x2 := ((-b) - math.Sqrt(b*b-4*a*c)) / (2 * a)
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
	return "Error"
}

func quadraticHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("quadratic.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	a, _ := strconv.ParseFloat(r.FormValue("a"), 64)
	b, _ := strconv.ParseFloat(r.FormValue("b"), 64)
	c, _ := strconv.ParseFloat(r.FormValue("c"), 64)
	d, _ := strconv.ParseFloat(r.FormValue("d"), 64)

	quadratic := Quadratic{a, b, c, d}

	quadraticAnswer := solveQuadraticEquation(quadratic)

	tmpl.Execute(w, struct {
		Success bool
		Message string
	}{true, quadraticAnswer})
}

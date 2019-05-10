package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type Equation struct {
	a float64
	b float64
	c float64
}

var result string

func solveEquation(equation Equation) string {
	delta := equation.b*equation.b - 4*(equation.a*equation.c)
	if delta > 0 {
		x1 := (-equation.b + math.Sqrt(delta)) / 2 * equation.a
		x2 := (-equation.b - math.Sqrt(delta)) / 2 * equation.a
		result := fmt.Sprintf("x1:%d x2:%d", x1, x2)
		return result
	}
	if delta == 0 {
		x1 := (-equation.b + math.Sqrt(delta)) / 2 * equation.a
		result := fmt.Sprint("x: %d", x1)
		return result
	}
	if delta < 0 {
		real := (-equation.b) / (2 * equation.a)
		img := math.Sqrt(-delta) / (2 * equation.a)
		cmplex := complex(real, img)
		cmplex2 := complex(real, -img)
		result := fmt.Sprintf("x1: %v x2: %v", cmplex, cmplex2)
		return result
	}
	return ""
}

func main() {

	equation := Equation{}
	templates := template.Must(template.ParseFiles("C:\Users\Alex\Desktop\EcuatiaDeGradul2\index.html"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "equation-template.html", equation); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/solveEquation", func(w http.ResponseWriter, r *http.Request) {
		if a := r.FormValue("a"); a != "" {
			equation.a, _ = strconv.ParseFloat(a, 64)
		}
		if x_first := r.FormValue("b"); b != "" {
			equation.b, _ = strconv.ParseFloat(b, 64)
		}
		if c := r.FormValue("c"); c != "" {
			equation.c, _ = strconv.ParseFloat(c, 64)
		}
		result = solveEquation(equation)
		fmt.Println(result)
	})

	//Print any errors from starting the webserver using fmt
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":5051", nil))
}
//05. Rezolvati ecuatia de gradul 2 folosind o aplicatie web

package main

import (
	"fmt"
	"html/template"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

type solution struct {
	X1 string
	X2 string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("main/*.html"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":9000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "quadratic.html", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	var x1Str,x2Str string
	sol := solution{}

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	aStr := r.FormValue("a")
	bStr := r.FormValue("b")
	cStr := r.FormValue("c")
	a, err1 := strconv.ParseFloat(aStr, 64)
	b, err2 := strconv.ParseFloat(bStr, 64)
	c, err3 := strconv.ParseFloat(cStr, 64)

	if err1 == nil && err2 == nil && err3 == nil {
		if a == 0 {
			x1 := (-c / b)
			x2 := 0
			x1Str = fmt.Sprint(x1)
			x2Str = fmt.Sprint(x2)

		} else {
			delta := b*b - 4*a*c
			if delta < 0 {

				x1 := (-complex(b,0) + cmplx.Sqrt(complex(delta,0))) / complex(2*a,0)
				x2 := (-complex(b,0) - cmplx.Sqrt(complex(delta,0))) / complex(2*a,0)

				x1Str = fmt.Sprint(real(x1))+"+"+fmt.Sprint(imag(x1))+"i"
				x2Str = fmt.Sprint(real(x2))+"-"+fmt.Sprint(imag(x2))+"i"

				fmt.Println( "Solutiile sunt egale.")

			} else {
				if delta == 0 {
					x1 := -b / (2 * a)
					x2 := x1

					x1Str = fmt.Sprint(x1)
					x2Str = fmt.Sprint(x2)

					fmt.Println( "Solutiile sunt egale.")

				} else {
					x1 := (-b + math.Sqrt(delta)) / (2 * a)
					x2 := (-b - math.Sqrt(delta)) / (2 * a)

					x1Str = fmt.Sprint(x1)
					x2Str = fmt.Sprint(x2)
				}
			}
		}
	}

	sol.X1 = x1Str
	sol.X2 = x2Str
	fmt.Printf("X1 = %v \n",x1Str)
	fmt.Printf("X2 = %v \n",x2Str)
	tmpl.Execute(w,sol)
}

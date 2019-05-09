package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

var tmpl *template.Template



func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":9090", nil)
}

func init() {
	tmpl = template.Must(template.ParseGlob("form.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "ecGr2.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var c1, c2 string
	var a, b, c, delta, x1, x2 float64
	a, _ = strconv.ParseFloat(r.FormValue("a"), 64)
	b, _ = strconv.ParseFloat(r.FormValue("b"), 64)
	c, _ = strconv.ParseFloat(r.FormValue("c"), 64)
	fmt.Println("a=", a, "b=", b, "c=", c)
	if a == 0 {
		if b == 0 {
			if c == 0 {
				fmt.Println("Exista o infinitate de solutii!")
			}
			{
				fmt.Println("Ecuatie imposibila!")
			}
		} else {
			x1 = -c / b
		}
	} else {
		delta = b*b - 4*a*c
		if delta < 0 {

			fmt.Println("Ec are solutii complexe.")
			c1 = strconv.FormatFloat(-b/(2*a), 'g', 1, 64);
			c2 = strconv.FormatFloat(-b/(2*a), 'g', 1, 64);
			c1 += "+"
			c1 += strconv.FormatFloat(math.Sqrt(math.Abs(delta))/(2*a),'g', 1, 64)
			c1 +=  "i";

			c2 += "-"
			c2 += strconv.FormatFloat(math.Sqrt(math.Abs(delta))/(2*a),'g', 1, 64)
			c2+="i"

			fmt.Println("c1=", c1, "c2=", c2)
		} else {
			if delta == 0 {
				x1 = -b / (2 * a)
				x2 = x1
				fmt.Println("Solutii: x1=", x1, " x2=", x2)
			} else {
				x1 = (-b + math.Sqrt(delta)) / (2 * a);
				x2 = (-b - math.Sqrt(delta)) / (2 * a);
				fmt.Println("Solutii: x1=", x1, " x2=", x2)
			}
		}
	}
	d := struct {
		X1 string
		X2 string
		C1 string
		C2 string
	}{
		X1: strconv.FormatFloat(x1, 'g', 1, 64),
		X2: strconv.FormatFloat(x2, 'g', 1, 64),
		C1: c1,
		C2: c2,
	}

	tmpl.ExecuteTemplate(w, "ecGr2.html", d)
}
package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

const anError = `<p class="error">%s</p>`

func floatConvert(text string) float64{
	number := float64(0)
	i, err := strconv.ParseFloat(text, 64)
	if err == nil {
		number = number*10 + i
	}
	return number
}

func compute(a, b, c float64) (rad1, rad2 float64) {

	var delta, x1, x2, real, imag float64
	delta = b*b - 4*a*c

	if delta > 0 {
		x1 = (-b + math.Sqrt(delta)) / (2 * a)
		x2 = (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Print("x1 = ", x1, "  x2= ", x2)
	} else if delta == 0 {
		x1 = (-b / (2 * a))
		x2 = (-b / (2 * a))
		fmt.Print("x1 = ", x1, "  x2= ", x2)
	} else {
		real = -b / (2 * a)
		imag = math.Sqrt(-delta) / (2 * a)
		fmt.Println("x1 = ", real, "-", imag, "i    x2 = ", real, "+", imag, "i")

	}
	return x1, x2
}

func formatResponse(rad1, rad2 float64) string {
	return fmt.Sprintf(`<h1> rad1 = %d </h1><div/><h1> rad2 = %d</h1>`,rad1, rad2)
}

func main() {
	tmpl := template.Must(template.ParseFiles("equationForm.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		a := floatConvert(r.FormValue("coefA"))
		b := floatConvert(r.FormValue("coefB"))
		c := floatConvert(r.FormValue("coefC"))

		rad1, rad2 := compute(a,b,c)

		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, anError, err)
			fmt.Println(err)
		}else{
			fmt.Fprint(w, formatResponse(rad1, rad2))
		}

		fmt.Println(rad1)
		fmt.Println(rad2)
		fmt.Println(tmpl)

		tmpl.Execute(w, struct{Success bool; x1, x2 float64}{Success:true,	 x1: rad1, x2: rad2})
	})

	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}


}

package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func ecGrad2(a,b,c float64) (float64,float64,float64,float64){
	delta := (b*b)-(4*a*c)
	var x1,x2 float64
	imag1 := math.NaN()
	imag2 := math.NaN()

	if delta > 0 {
		x1 = (-b + math.Sqrt(delta))/(2*a)
		x2 = (-b - math.Sqrt(delta))/(2*a)
	} else if delta == 0 {
		x1, x2 = (-b + 0)/(2*a),(-b - 0)/(2*a)
	} else {
		x1, x2 = (-b)/(2*a),(-b)/(2*a)
		imag1, imag2 = math.Sqrt(-delta)/(2*a),math.Sqrt(-delta)/(2*a)
	}
	return x1,x2,imag1,imag2
}

func floatConvert(text []string) float64{
	number := float64(0)
	for _, elem := range text {
		i, err := strconv.ParseFloat(elem, 64)
		if err == nil {
			number = number*10 + i
		}
	}
	return number
}

func main(){
	http.HandleFunc("/",beforePage)
	http.HandleFunc("/ecuatie",afterPage)
	err := http.ListenAndServe(":8081",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}
}

func beforePage(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprint(w, "Hello!")
}

type AfterVariables struct {
	X1 string
	X2 string
}

func afterPage(w http.ResponseWriter, r *http.Request){
	var apVar AfterVariables
	fmt.Println("method",r.Method)
	r.ParseForm()
	x1,x2,imag1,imag2 := ecGrad2(floatConvert(r.Form["coefA"]),floatConvert(r.Form["coefB"]),floatConvert(r.Form["coefC"]))
	if math.IsNaN(x1) && math.IsNaN(x2){
		apVar = AfterVariables{
			X1 : " ",
			X2 : " ",
		}
	} else if math.IsNaN(imag1) && math.IsNaN(imag2) {
		apVar = AfterVariables{
			X1: fmt.Sprintf("%f", x1),
			X2: fmt.Sprintf("%f", x2),
		}
	} else {
		apVar = AfterVariables{
			X1: fmt.Sprintf("%f", x1) + " + " + fmt.Sprintf("%f",imag1) + "*i",
			X2: fmt.Sprintf("%f", x2) + " - " + fmt.Sprintf("%f",imag2) + "*i",
		}
	}
	t, err := template.ParseFiles("ecGr2.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, apVar)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
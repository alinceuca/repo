package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Ecuatie struct {
	coeficientA string
	coeficientB string
	coeficientC string
}

func main() {

	tmpl:= template.Must(template.ParseFiles("main/Ecuatie.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {

		if request.Method != http.MethodPost{
			tmpl.Execute(w,nil)
			return
		}

		details := Ecuatie{
			coeficientA:request.FormValue("primul"),
			coeficientB:request.FormValue("alDoilea"),
			coeficientC:request.FormValue("alTreilea"),

		}

		//fmt.Println(details.coeficientC)
		//fmt.Println(details.coeficientB)
		//fmt.Println(details.coeficientC)

		var raspuns = calculeazaEcuatia(details.coeficientA,details.coeficientB,details.coeficientC)

		trimiteRaspuns := strings.TrimPrefix(fmt.Sprint(raspuns),"/")
		trimiteRaspuns = "Solutile ecuatiei sunt:"+ trimiteRaspuns




		_ = details

		tmpl.Execute(w, struct {
			Succes bool
		}{true})

		if raspuns[1] == 0 && raspuns[0]==0{
			trimiteRaspuns := "nu se poate calcula ecuatia, delta este negativa"
			w.Write([]byte(trimiteRaspuns))
		}else {
			if raspuns[1] ==  raspuns[0]{
				var aux int =int(raspuns[0])
				trimiteRaspuns := "ecuatia are o singura solutie " + fmt.Sprint(aux)
				w.Write([]byte (trimiteRaspuns))
			}else {
				w.Write([]byte(trimiteRaspuns))
			}
		}

	})

	http.ListenAndServe(":1080", nil)
	
}

func calculeazaEcuatia(a,b,c string) [2] float64{

	i1,err := strconv.ParseFloat(a,64)
	if err == nil{
	}
	i2,err := strconv.ParseFloat(b,64)
	if err == nil{
	}
	i3,err := strconv.ParseFloat(c,64)
	if err == nil{
	}

	delta := (i2*i2-4*i1*i3)

	var rezultatEcuatie[2] float64

	if delta >0 {
		x1 := (-i2 + math.Sqrt(delta)) / 2 * i1
		x2 := (-i2 - math.Sqrt(delta)) / 2 * i1
		rezultatEcuatie[0]=x1
		rezultatEcuatie[1]=x2
	}else {
		if (delta == 0) {
			x1 := (-i2  / 2 * i1)
			rezultatEcuatie[0]=x1
			rezultatEcuatie[1]=x1
		}else {
			rezultatEcuatie[0] = 0
			rezultatEcuatie[1] = 0
		}
	}

	return rezultatEcuatie

}

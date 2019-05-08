
package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type SolutiiEcuatieGradDoi struct {
	Mesaj   string
	NumarSolutii int
	PrimaSolutie float64
	DoiSolutie float64
}

func main() {
tmpl := template.Must(template.ParseFiles("html/template/homepage.html"))


http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {


	 a, err := strconv.ParseInt(r.FormValue("x^2"), 10, 64);
	 b, err := strconv.ParseInt(r.FormValue("x"), 10, 64);
	 c, err := strconv.ParseInt(r.FormValue("x^0"), 10, 64);
	 if err ==nil {
	 	print("")
	 }
	 fmt.Println("", rezolvareEcuatieGradDoi(int(a) , int(b), int(c) ));
    HomePageVars := SolutiiEcuatieGradDoi{ //store the date and time in a struct
			Mesaj: rezolvareEcuatieGradDoi(int(a) , int(b), int(c)).Mesaj,
			NumarSolutii: rezolvareEcuatieGradDoi(int(a) , int(b), int(c)).NumarSolutii,
			PrimaSolutie: rezolvareEcuatieGradDoi(int(a) , int(b), int(c)).PrimaSolutie,
			DoiSolutie: rezolvareEcuatieGradDoi(int(a) , int(b), int(c)).DoiSolutie,
	}

	// do something with details



tmpl.Execute(w, HomePageVars)
})

http.ListenAndServe(":8080", nil)
}

func rezolvareEcuatieGradDoi( a int, b int, c int) SolutiiEcuatieGradDoi {
	var mesaj string = "" ;
	var sol1 float64 ;
	var sol2 float64 ;
	var nrSol int;
	if a == 0 {
		mesaj= "Coeficinetii ecuatiei nu sunt corecti";

	} else {
		var delta int = b*b - 4*a*c;
		if delta < 0 {
			mesaj = "Ecuatia nu are solutii reale";
			nrSol = 0;
		} else {
			if delta == 0 {
				 sol1 = float64(-b / 2 * a);
				 nrSol = 1;
			}else {
				sol1 = (float64(-b) + math.Sqrt(float64(delta)))/ (2 * float64(a));
				sol2 = (float64(-b) - math.Sqrt(float64(delta)))/ (2 * float64(a));
				nrSol = 2;
			}
		}
	}
	details := SolutiiEcuatieGradDoi{
		Mesaj:   mesaj,
		NumarSolutii: nrSol,
		PrimaSolutie: sol1,
		DoiSolutie: sol2,
	}
	return  details;
}


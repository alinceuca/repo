package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Data struct {
	An, Luna, Zi int
}

func New(dataString string) (Data, error) {

	// se imparte stringul in parti separate de "-"
	s := strings.Split(dataString, "-")
	var err  error

	// valideaza anul, ca e numar si e mai mare ca 0
	an, errAn := strconv.Atoi(s[0])
	if (an<1 || errAn != nil){
		an = 0
		if errAn != nil {
			err = errAn
		} else {
			err = errors.New("Anul trebuie sa fie mai mare decat 0")
		}
	}
	// valideaza luna, ca e numar si e intre 1 si 12
	luna, errLuna := strconv.Atoi(s[1])
	if (luna<1 || luna>12 || errLuna != nil){
		luna = 0
		if errLuna != nil {
			err = errLuna
		} else {
			err = errors.New("Luna trebuie sa fie intre 1 si 12")
		}
	}
	// valideaza zi, ca e numar si e intre 1 si maximul lunii respective
	zi, errZi := strconv.Atoi(s[2])
	if errZi != nil {
		zi = 0
		err = errZi
	} else if luna == 2 {

		if EsteAnBisect(an) {
			if (zi<1 || zi>29) {
				err = errors.New("Ziua trebuie sa fie intre 1 si 29")
			}
		} else {
			if (zi<1 || zi>28) {
				err = errors.New("Ziua trebuie sa fie intre 1 si 28")
			}
		}
	} else if (zi<1 || zi>31) {
		err = errors.New("Ziua trebuie sa fie intre 1 si 31")
	}

	d := Data{an, luna, zi}
	if err != nil {
		return d, err
	}
	return d, nil
}

func (d Data) ToString() string {

	return fmt.Sprintf("%d-%d-%d", d.An, d.Luna, d.Zi)
}

func (d1 Data) comparaDate(d2 Data) int{
	// -1 d1<d2
	// 0 d1==d2
	// 1 d1>d2

	if d1.An < d2.An{
		return -1
	}
	if d1.An > d2.An{
		return 1
	}
	if d1.Luna < d2.Luna{
		return -1
	}
	if d1.Luna > d2.Luna{
		return 1
	}
	if d1.Zi < d2.Zi{
		return -1
	}else{
		return 1
	}

	return 0
}

func (d Data) plusOZi() Data{
	nrZileLuni := [13]int{0,31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if d.Zi < nrZileLuni[d.Luna]{
		d.Zi++
	} else {
		if d.Zi == nrZileLuni[d.Luna] {
			if d.Luna == 2 && EsteAnBisect(d.An)  {
				d.Zi++
			} else {
				if d.Luna == 12 {
					d.Zi = 1
					d.Luna = 1
					d.An++

				} else {
					d.Zi = 1
					d.Luna++
				}
			}

		} else {
			if d.Luna == 2 && EsteAnBisect(d.An)  {
				d.Zi = 1
				d.Luna++
			}
		}
	}
	return d
}

func (d Data) ZiuaDinSaptamana() string {
	// calculeaza ziua din saptamana
	saptamna:=[7] string {"Duminica","Luni","Marti","Miercuri","Joi","Vineri","Sambata"}
	// se alege o data etalon de duminica
	dataEtalon:= Data{2019,4,7}
	var nrzile int

	if dataEtalon.comparaDate(d) == -1{
		nrzile=dataEtalon.NumarZile(d)
		return saptamna[nrzile%7]
	} else if dataEtalon.comparaDate(d)==1{
		nrzile=d.NumarZile(dataEtalon)
		return saptamna[(7-(nrzile%7))%7]
	} else {
		return saptamna[0]
	}
}

func (d1 Data) NumarZile(d2 Data) int{
	var totalZile int
	if d1.comparaDate(d2)== -1{
		for (d1.comparaDate(d2)== -1){
			d1=d1.plusOZi()
			totalZile++
		}
	}else if d1.comparaDate(d2)== 1{
		for (d2.comparaDate(d1)== -1){
			d2=d2.plusOZi()
			totalZile--
		}

	}else
	{
		totalZile=0
	}
	return totalZile
}

func (d Data) PlusZile(zile int) Data{
	for i:=0; i<zile; i++ {
		d = d.plusOZi()
	}

	return d
}


func (d Data) minusOZi() Data {
	nrZileLuni := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if d.Zi > 1 {
		d.Zi--
	} else {
		// la 1 martie va fi 28 sau 29 feb
		if d.Luna == 3 {
			if EsteAnBisect(d.An) {
				d.Luna--
				d.Zi = nrZileLuni[d.Luna] + 1
			} else {
				d.Luna--
				d.Zi = nrZileLuni[d.Luna]
			}
		}

		// la 1 ian va fi 31 dec si un an mai putin
		if d.Luna == 1 {
			d.Luna = 12
			d.Zi = nrZileLuni[d.Luna]
			d.An--
		} else {
			//cand nu e ian si nici feb
			d.Luna--
			d.Zi = nrZileLuni[d.Luna]
		}
	}
	return d
}

func (d Data) MinusZile(zile int) Data{
	for i:=0; i<zile; i++ {
		d = d.minusOZi()
	}
	return d
}

func EsteAnBisect(an int) bool {
	// verifica daca anul este bisect
	// dacă (an nu este divizibil cu 4) atunci (an obișnuit)
	// altfel dacă (an nu este divizibil cu 100) atunci (an bisect)
	// altfel dacă (an nu este divizibil cu 400) atunci (an obișnuit)
	// altfel (an bisect)
	if an % 4!=0 {
		return false
	} else if an % 100!=0 {
		return true
	} else if an % 400!=0{
		return false
	}
	return true
}
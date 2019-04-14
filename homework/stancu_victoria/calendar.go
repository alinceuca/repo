package main

import (
	"fmt"
)
type Data struct {
	Zi 		int
	Luna  	int
	An 		int
}
func main() {


	var data1,data2,data3 Data

	fmt.Println("Data calendaristica nr 1:\n ")
	data1 = citireData()
	fmt.Printf(" Data de %v este  %s \n",data1,ziuaDinSapatmana(data1))

	fmt.Print("Data calendaristica nr 2: \n")
	data2 = citireData()
	fmt.Printf("Data de %v este %s \n",data2,ziuaDinSapatmana(data2))

	nrZile:=numarZile(data1,data2)
	fmt.Printf("\nIntre datele %v - %v aven %d zile \n",data1,data2,nrZile)

	ani, luni, saptamini, zile := anLuniSaptamaniZile(nrZile)
	fmt.Printf(" \n Numarul de zile aproximativ = %d insemna %d ani,%d luni %d saptamini, %d zile \n",nrZile,ani,luni,saptamini,zile)

	ani1, luni1, saptamini1, zile1 :=numarAnLuniSaptZile(data1,data2)
	fmt.Printf(" \n Numarul de zile = %d insemna %d ani,%d luni %d saptamini, %d zile \n",nrZile,ani1,luni1,saptamini1,zile1)

	data3=plusZile(data2,20)
	fmt.Printf("Data %v + 20 zile = %v",data2,data3)
	fmt.Printf(" ziua din saptamana= %s \n",ziuaDinSapatmana(data3))

	data3=minusZile(data2,20)
	fmt.Printf("Data %v - 20 zile = %v",data2,data3)
	fmt.Printf(" ziua din saptamana= %s \n",ziuaDinSapatmana(data3))
}

func citireData() (Data){
	var dataCitita Data
	var zi,luna,an int

	for  {
		fmt.Print("Anul: ")
		fmt.Scanf("%d",&an)
		if (an<1 ){
			fmt.Println("Anul trebuie sa fie mai mare decat 0! ")
		}else{
			break
		}
	}

	for  {
		fmt.Print("Luna: ")
		fmt.Scanf("%d",&luna)
		if (luna<1 || luna>12){
			fmt.Println("Luna poate fi intre 1 si 12! ")
		}else{
			break
		}
	}

	for  {
		fmt.Print("Ziua: ")
		fmt.Scanf("%d", &zi)
		if luna==2{
			if anBisect(an){
				if (zi<1 || zi>29){
					fmt.Println("Ziua poate fi intre 1 si 29! ")
				}else{
					break
				}
			}else{
				if (zi<1 || zi>28){
					fmt.Println("Ziua poate fi intre 1 si 28! ")
				}else{
					break
				}
			}

		}else if (zi<1 || zi>31){
			fmt.Println("Ziua poate fi intre 1 si 31! ")
		}else{
			break
		}
	}


	dataCitita.Zi=zi
	dataCitita.Luna=luna
	dataCitita.An=an

	return dataCitita
}

func comparDate(d1, d2 Data) int{
	// -1 d1<d2
	// 0 d1==d2
	// 1 d1>d2
	if d1.An<d2.An{
		return -1
	}
	if d1.An>d2.An{
		return 1
	}
	if d1.Luna<d2.Luna{
		return -1
	}
	if d1.Luna>d2.Luna{
		return 1
	}
	if d1.Zi<d2.Zi{
		return -1
	}
	if d1.Zi>d2.Zi{
		return 1
	}


	return 0
}


func comparAn(d1,d2 Data) int{
	// 1 An1>An2; 0 A1==A2; -1 A1<A2
	if (d1.An<d2.An){
		return -1
	}else{
		if(d1.An==d2.An){
			return 0
		}
	}
	return 1
}

func plusAn(d1 Data)Data{
	d1.An++
	return d1
}

func comparLuna(d1,d2 Data) int{
	//1 nu adun, 0 am terminat,-1 adun
	if (d1.An>d2.An){
		return 1
	}
	if(d1.An==d2.An){
		if(d1.Luna<d2.Luna){
			return -1

		}else{
			if(d1.Luna==d2.Luna){
				return 0
			}
		}
     }

	return -1

}

func scadLuna(d Data) Data{
 	if (d.Luna==1){
 		d.An--
 		d.Luna=12
	}else{
		d.Luna--
	}

 	return d
 }

func plusLuna(d Data) Data{
	 if (d.Luna==12){
	 	d.An++
	 	d.Luna=1
	 }else{
	 	d.Luna++
	 }
	return d
}

func numarAnLuniSaptZile(d1,d2 Data) (int,int,int,int){
	var ani,luni,saptamani,zile int


	if comparDate(d1,d2)== -1{
       //d1<d2
		if (comparAn(d1,d2)==-1){ //a1<a2
			for (comparAn(d1, d2) == -1) {
					d1 = plusAn(d1)
					ani++
			}
			if (comparDate(d1,d2)==1){
				ani--
				d1.An--
			}

		}

		if (comparDate(d1,d2)==-1){
				for (comparLuna(d1, d2) == -1) {
					d1 = plusLuna(d1)
					luni++
				}
			if (comparDate(d1,d2)==1){
				luni--
				d1.Luna--
			}

		}
		for (comparDate(d1,d2)== -1){
			d1=plusOZi(d1)
			zile++
		}
	}else if comparDate(d1,d2)== 1{
		//d2<d1
		if (comparAn(d2,d1)==-1){ //a2<a1
			for (comparAn(d2, d1) == -1) {
				d2 = plusAn(d2)
				ani++
			}
			if (comparDate(d2,d1)==1){
				ani--
				d2.An--
			}

		}

		if (comparDate(d2,d1)==-1){

			for (comparLuna(d2, d1) == -1) {
					d2 = plusLuna(d2)
					luni++
			}
			if (comparDate(d2,d1)==1){
				luni--
				d2.Luna--
			}
		}
		for (comparDate(d2,d1)== -1){
			d2=plusOZi(d2)
			zile++
		}

	}else
	{
		zile=0
	}

	saptamani=zile/7
	zile=zile%7

	return ani,luni,saptamani,zile
}

func anBisect(an int) bool{
	/*dacă (year nu este divizibil cu 4) atunci (an obișnuit)
		altfel dacă (year nu este divizibil cu 100) atunci (an bisect)
			altfel dacă (year nu este divizibil cu 400) atunci (an obișnuit)
				altfel (an bisect)*/
	if an%4!=0 {
		return false
	}else if  an%100!=0 {
		return true
	}else if an%400!=0{
		return false
	}
	return true
}

func plusOZi(d Data) Data{
	nrZileLuni := [13]int{0,31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if d.Zi<nrZileLuni[d.Luna]{
		d.Zi++
	}else{
		if d.Zi==nrZileLuni[d.Luna] {
			if d.Luna == 2 && anBisect(d.An)  {
				d.Zi++
			}else{
				if d.Luna==12 {
					d.Zi=1
					d.Luna=1
					d.An++

				}else{
					d.Zi=1
					d.Luna++
				}
			}

		}else{
			if d.Luna == 2 && anBisect(d.An)  {
				d.Zi=1
				d.Luna++
			}
		}

	}

	return d
}

func numarZile(d1,d2 Data) int{
	var totalZile int
	if comparDate(d1,d2)== -1{
		for (comparDate(d1,d2)== -1){
			d1=plusOZi(d1)
			totalZile++
		}
	}else if comparDate(d1,d2)== 1{
		for (comparDate(d2,d1)== -1){
			d2=plusOZi(d2)
			totalZile--
		}

	}else
	{
		totalZile=0
	}
	return totalZile
}


func anLuniSaptamaniZile(zileTotal int) (int,int,int,int){
	var an,luni,saptamani,zile int
	an=zileTotal/365
	zileTotal-=an*365
	if an>4{
		zileTotal-=an/4

	}
	if zileTotal>0 {

		luni = int( float64(zileTotal) /30.5)
		zileTotal-=int(float64(luni)*30.5)

		if zileTotal>0 {
			saptamani = zileTotal  / 7
			zileTotal-=saptamani*7
			zile = zileTotal
		}
	}
	return an,luni,saptamani,zile
}

func plusZile(d Data, zile int) Data{
	for i:=0;i<zile;i++{
		d=plusOZi(d)
	}

	return d
}

func minusOZi(d Data) Data {
	nrZileLuni := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if d.Zi > 1 {
		d.Zi--
	} else {
		//daca este 1 matie atunci va fi 28 feb sau 29 feb
		if d.Luna == 3 {
			if anBisect(d.An) {
				d.Luna--
				d.Zi = nrZileLuni[d.Luna] + 1
			} else {
				d.Luna--
				d.Zi = nrZileLuni[d.Luna]
			}
		}
		// daca e 1 ian va fi 31 dec si un an mai putin
		if d.Luna == 1 {
			d.Luna = 12
			d.Zi = nrZileLuni[d.Luna]
			d.An--
		}else{
			//cand nu e ian si nici feb
			d.Luna--
			d.Zi = nrZileLuni[d.Luna]
		}

	}

	return d
}

func minusZile(d Data, zile int) Data{
	for i:=0;i<zile;i++{
		d=minusOZi(d)
	}

	return d
}

func ziuaDinSapatmana(d Data) string{
	saptamna:=[7] string {"Duminica","Luni","Marti","Miercuri","Joi","Vineri","Sambata"}
	dataEtalon:= Data{10,3,2019}
	var nrzile int
	if comparDate(dataEtalon,d)==-1{
		nrzile=numarZile(dataEtalon,d)
		return saptamna[nrzile%7]
	}else if comparDate(dataEtalon,d)==1{
		nrzile=numarZile(d,dataEtalon)
		return saptamna[(7-(nrzile%7))%7]
	}else{
		return saptamna[0]
	}

}

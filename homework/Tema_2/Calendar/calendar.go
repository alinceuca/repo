package main

import (
	"fmt"
	"./data"
)


func main(){
	var dataString string
	var data1, data2 data.Data
	var err error

	// repeta citirea unei date pana cand formatul este corect si nu sunt erori
	for  {
		fmt.Print("Data 1 (AAAA-LL-ZZ):")
		fmt.Scanf("%s", &dataString)
		data1, err = data.New(dataString)

		if err == nil {
			break
		}
		fmt.Printf("  Eroare: %+v\n", err)
	}
	fmt.Printf("Data pica intr-o zi de %s \n", data1.ZiuaDinSaptamana())


	for  {
		fmt.Print("\nData 2 (AAAA-LL-ZZ):")
		fmt.Scanf("%s", &dataString)
		data2, err = data.New(dataString)

		if err == nil {
			break
		}
		fmt.Printf("  Eroare: %+v\n", err)
	}
	fmt.Printf("Data pica intr-o zi de %s \n", data2.ZiuaDinSaptamana())

	nrZile:=data1.NumarZile(data2)
	fmt.Printf("\nIntre %v si %v sunt %d zile \n", data1.ToString(), data2.ToString(), nrZile)

	ani, luni, saptamani, zile := aniLuniSaptamaniZile(nrZile)
	fmt.Printf("Intre %v si %v sunt %d ani, %d luni, %d saptamini, %d zile\n", data1.ToString(), data2.ToString(), ani, luni, saptamani, zile)

	var deltaZile int
	fmt.Print("\nNumar de zile de adunat si scazut:")
	fmt.Scanf("%d", &deltaZile)

	data1Plus := data1.PlusZile(deltaZile)
	fmt.Printf("Data peste %d zile %v\n", deltaZile, data1Plus.ToString())
	data1Minus := data1.MinusZile(deltaZile)
	fmt.Printf("Data acum %d zile %v", deltaZile, data1Minus.ToString())

}

func aniLuniSaptamaniZile(zileTotal int) (int,int,int,int){
	var an, luni, saptamani, zile int
	an = zileTotal / 365
	zileTotal -= an * 365
	if an > 4 {
		zileTotal -= an / 4
	}
	if zileTotal > 0 {
		luni = int(float64(zileTotal) / 30.5)
		zileTotal -= int(float64(luni) * 30.5)

		if zileTotal > 0 {
			saptamani = zileTotal  / 7
			zileTotal -= saptamani * 7
			zile = zileTotal
		}
	}
	return an, luni, saptamani, zile
}


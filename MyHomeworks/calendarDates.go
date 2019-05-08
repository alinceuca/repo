package main

import (
	"fmt"
	"time"
)

func customDate(year, month, day int) time.Time{
	return time.Date(year, time.Month(month), day, 0, 0,0,0,time.UTC)
}

func difference(date1, date2 time.Time){

	if date2.After(date1){
		aux := date1
		date1 = date2
		date2 = aux
	}

	diff := date1.Sub(date2)
	fmt.Println("Diferenta in ore: ", diff.Hours())
	fmt.Println("Diferenta in minute: ", diff.Minutes())
	fmt.Println("Diferenta in secunde: ", diff.Seconds())
	fmt.Println("Diferenta in zile: ", diff.Hours()/24)

	years := int(diff.Seconds() / 31557600)
	months := int(diff.Seconds()/2629743) % 12
	weeks := (int(diff.Seconds()/604800) % 12) % 4
	days := ((int(diff.Seconds()/86400) % 12) % 4) % 7

	fmt.Println("Ani: ", years)
	fmt.Println("Luni: ", months)
	fmt.Println("Saptamani: ", weeks)
	fmt.Println("Zile: ", days)
}

func weekDay(date time.Time){
	fmt.Println(date.Weekday())
}

func addDays(date time.Time, days int) time.Time{
	//fmt.Println("Rezultat: ", date.AddDate(0,0,days))
	return date.AddDate(0,0, days)
}
func substractDays(date time.Time, days int) time.Time{
	//fmt.Println(" Rezultat: ", date.AddDate(0, 0, -days))
	return date.AddDate(0, 0, -days)
}

func main(){
	var date1 = customDate(2019, 7, 10)
	var date2 = customDate(2019, 12,11)
	difference(date1, date2)
	date1 = substractDays(date1, 1)
	fmt.Println(date1)
}
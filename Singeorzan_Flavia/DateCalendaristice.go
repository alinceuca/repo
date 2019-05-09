package main

import (
	"fmt"
	"time"
)


func main(){
	date := time.Now()

	secondDate := time.Date(1998, time.June, 30, 23, 30, 0, 0, time.UTC)

	DateDifferenceInDays(date, secondDate)
	DateDifferenceFull(date, secondDate)

	fmt.Println("Azi: ", date.String())
	fmt.Println("Azi + 10 zile: ", AddDays(date, 10).String())
	fmt.Println("Azi - 10 zile: ", RemoveDays(date, 10).String())

	PrintWeekDay(date)

}


//DateDifferenceInDays rreturneaza diferenta in zile dintre 2 date
func DateDifferenceInDays(a, b time.Time) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	days := int(b.Sub(a).Hours() / 24)
	fmt.Println("Diferenta in zile este: ", days)
}

//DateDifferenceFull va afisa ani,luni, zile, minute, secunde
func DateDifferenceFull(a, b time.Time)(year, month, weeks, day int) {

	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()


	year = int(y2 - y1)
	month = int(M2 - M1)
	weeks = 0
	day = int(d2 - d1)


	if day < 0 {
		// zilele in luni:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
		//zilele in saptamani:
		for day>7 {
			weeks = weeks+1
			day = day - 7
		}
	}
	if month < 0 {
		month += 12
		year--
	}

	fmt.Println("Ani : ", year)
	fmt.Println("Luni : ", month)
	fmt.Println("Saptamani: ", weeks)
	fmt.Println("Zile : ", day)
	return
}

//AddDays adauga zile la o data
func AddDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

func RemoveDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, -days)
}

func PrintWeekDay(date time.Time) {
	fmt.Println("Today is: ", date.Weekday().String())
}
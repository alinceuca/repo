package main

import (
	"fmt"
	"time"
)


func main(){
	date := time.Now()

	secondDate := time.Date(1997, time.October, 8, 20, 30, 0, 0, time.UTC)

	DateDifferenceInDays(date, secondDate)
	DateDifferenceFull(date, secondDate)

	fmt.Println("Today: ", date.String())
	fmt.Println("Today plus 2 weeks: ", AddDays(date, 14).String())
	fmt.Println("Today minus 2 weeks: ", RemoveDays(date, 14).String())

	PrintWeekDay(date)

}


//DateDifferenceInDays returns the difference in days between 2 dates
func DateDifferenceInDays(a, b time.Time) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	days := int(b.Sub(a).Hours() / 24)
	fmt.Println("The difference in days is", days)
}

//DateDifferenceFull will print years, months, days, minutes, seconds
func DateDifferenceFull(a, b time.Time)(year, month, weeks, day int) {
	//duration := otherDate.Sub(date)

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


	// Normalize negative values

	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
		//days in week:
		for day>7 {
			weeks = weeks+1
			day = day - 7
		}
	}
	if month < 0 {
		month += 12
		year--
	}

	fmt.Println("Years : ", year)
	fmt.Println("Months : ", month)
	fmt.Println("Weeks : ", weeks)
	fmt.Println("Days : ", day)
	return
}

//AddDays adds days to a date
func AddDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

//RemoveDays substracts days to a date
func RemoveDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, -days)
}

//PrintWeekDay substracts days to a date
func PrintWeekDay(date time.Time) {
	fmt.Println("Today is: ", date.Weekday().String())
}

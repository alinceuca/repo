package main

import (
	"strconv"
	"time"
)

func main() {

}

func addDaysToDate(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

func dateToDayOfTheWeek(date time.Time) string {
	return date.Weekday().String()
}

func dateDifferencesInDays(firstDate, secondDate time.Time) string {
	days := firstDate.Sub(secondDate).Hours() / 24
	return strconv.FormatFloat(days, 'f', 6, 64) + " days"
}
func dateDifference(firstDate, secondDate time.Time) string {
	days := firstDate.Sub(secondDate).Hours() / 24
	weeks := days / 7
	months := weeks / 30.436875
	years := months / 12

	return strconv.FormatFloat(days, 'f', 6, 64) + " days or" +
		strconv.FormatFloat(weeks, 'f', 6, 64) + " weeks or" +
		strconv.FormatFloat(months, 'f', 6, 64) + " months or" +
		strconv.FormatFloat(years, 'f', 6, 64) + " years or"
}

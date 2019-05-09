package main
import (
	"fmt"
	"time"
)
func main() {
	date := time.Now()
	date2 := time.Date(2090, time.September, 10, 23, 0, 0, 0, time.UTC)
	getDifferenceBetween2DatesInDays(date, date2)
	getDifferenceBetween2Dates(date, date2)
	fmt.Println(date.String())
	fmt.Println(sumDays(date, 10).String())
	fmt.Println(removeDays(date, -10).String())
	 PrintWeekDay(date)
}
func getCurrentDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}


func getDifferenceBetween2DatesInDays(date time.Time, otherDate time.Time) {
	days := int(otherDate.Sub(date).Hours() / 24)
	fmt.Println("The difference in days is", days)
}

func getDifferenceBetween2Dates(date time.Time, otherDate time.Time) {
	duration := otherDate.Sub(date)

	years := int(duration.Seconds() / 31207680)

	months := int(duration.Seconds()/2600640) % 12
	weeks := (int(duration.Seconds()/604800) % 12) % 4
	days := ((int(duration.Seconds()/86400) % 12) % 4) % 7

	fmt.Println("Years : ", years)
	fmt.Println("Months : ", months)
	fmt.Println("Weeks : ", weeks)
	fmt.Println("Days : ", days)
}

func sumDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

func removeDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, -days)
}

func PrintWeekDay(date time.Time) {
	fmt.Println(date.Weekday().String())
}
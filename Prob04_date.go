//04. Operatii cu date calendaristice - diferenta dintre doua date in zile;
//				    - diferenta dintre doua date calendaristice in ani, luni, saptamani, zile
//				    - data + nr zile
//				    - data - nr zile
//				    - in ce zi a saptamanii se gaseste o data calendaristica

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const shortDate = "2019-01-01"

func main() {
	reader := bufio.NewReader(os.Stdin)

	var optionNumber = 0
	for optionNumber != 5 {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~ Meniu ~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("1. Add two dates and find the time between the two")
		fmt.Println("2. Add a date and add a number of days")
		fmt.Println("3. Add a date and substract a number of days")
		fmt.Println("4. Get the day of the week from a date")
		fmt.Println("5. Exit")

		optionNumber, err := readStringAndConvertToNumber(reader,"Optiunea Dvs. : ")
		fmt.Println()

		if err != nil {
			fmt.Println(err)
		} else {
			if optionNumber < 1 || optionNumber > 5 {
				fmt.Println("Optiune eronata")
				optionNumber = 0
			} else {
				switch optionNumber {
				case 1:
					addTwoDatesAndFindTheDifferenceBetweenThem(reader)
				case 2:
					addOneDateAndDoOperation(reader, "add")
				case 3:
					addOneDateAndDoOperation(reader, "substract")
				case 4:
					addOneDateAndDoOperation(reader, "getWeekday")
				case 5:
					os.Exit(0)
				}
			}
		}
		fmt.Println()
	}
}

func readAndValidateDate(reader *bufio.Reader, message string) (time.Time, error) {
	fmt.Print(message)
	date, _ := reader.ReadString('\n')
	date = strings.Replace(date, "\n", "", -1)
	tdate, err := validateDate(date)

	return tdate, err
}

func addTwoDatesAndFindTheDifferenceBetweenThem(reader *bufio.Reader) {
	fmt.Println("Date format: yyyy-mm-dd")
	startdate, err1 := readAndValidateDate(reader, "Start Date: ")
	enddate, err2 := readAndValidateDate(reader, "End Date: ")

	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	if err1 == nil && err2 == nil {
		res, err := checkDataBoundariesStr(startdate, enddate)
		if res == true {
			var diff = enddate.Sub(startdate)

			var days = int(diff.Hours() / 24)
			fmt.Printf("Diffrence in days  : %d days\n", days)

			var weeks = int(diff.Hours() / (24 * 7))
			fmt.Printf("Diffrence in weeks : %d weeks\n", weeks)

			fmt.Printf("Diffrence in monts : %d monts\n", diffMonths(startdate, enddate))
			fmt.Printf("Diffrence in years : %d years\n", enddate.Year()-startdate.Year())
		} else {
			fmt.Println(err)
		}
	}
}

func readStringAndConvertToNumber(reader *bufio.Reader, message string) (int, error) {
	fmt.Print(message)
	stringNumber, _ := reader.ReadString('\n')
	stringNumber = strings.Replace(stringNumber, "\n", "", -1)

	number, err := strconv.Atoi(stringNumber)
	return number, err
}

func addOneDateAndDoOperation(reader *bufio.Reader, operation string) {
	fmt.Println("Date format: yyyy-mm-dd")
	var dateNew time.Time

	date, err := readAndValidateDate(reader, "Date: ")
	if err == nil {
		if operation == "getWeekday" {
			fmt.Println("The weekday is: ", date.Weekday())
		} else {
			numberDays, err := readStringAndConvertToNumber(reader, "Number of days to "+operation+": ")
			if err != nil {
				fmt.Println(err)
			} else {
				switch operation {
				case "add":
					dateNew = date.AddDate(0, 0, numberDays)
				case "substract":
					dateNew = date.AddDate(0, 0, -numberDays)
				}
				fmt.Println("Result date: ", dateNew)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func diffMonths(startdate, enddate time.Time) int {
	var months = 0
	for startdate.Before(enddate) {
		startdate = startdate.AddDate(0, 1, 0)
		if startdate.Year() != enddate.Year() {
			months++
		} else {
			if startdate.Month() != enddate.Month() {
				months++
			}
		}
	}
	return months
}

func shortDateFromString(ds string) (time.Time, error) {
	t, err := time.Parse(shortDate, ds)
	if err != nil {
		return t, err
	}
	return t, nil
}

func validateDate(date string) (time.Time, error) {
	tdate, err := shortDateFromString(date)
	if err != nil {
		return tdate, err
	}
	return tdate, nil
}

func checkDataBoundariesStr(startdate, enddate time.Time) (bool, error) {
	if startdate.After(enddate) {
		return false, fmt.Errorf("startdate > enddate - please set proper data boundaries")
	}
	return true, nil
}

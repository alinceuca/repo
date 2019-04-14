package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func citireValidareData(reader *bufio.Reader, mesaj string) (time.Time, error) {
	fmt.Print(mesaj)
	data, _ := reader.ReadString('\n')
	data = strings.Replace(data, "\n", "", -1)
	data2, err := validareData(data)
	return data2, err
}

func diferentaIntreDouaDate(reader *bufio.Reader) {
	fmt.Println("format data: aaaa-ll-zz")
	dataInceput, err1 := citireValidareData(reader, "introduceti data de inceput: ")
	dataSfarsit, err2 := citireValidareData(reader, "introduceti data de sfarsit: ")
	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	if err1 == nil && err2 == nil {
		res, err := verificareData(dataInceput, dataSfarsit)
		if res == true {
			var diferenta = dataSfarsit.Sub(dataInceput)
			var zile = int(diferenta.Hours() / 24)
			fmt.Printf("Diferenta in zile este de : %d zile\n", zile)
			var saptamani = int(diferenta.Hours() / (24 * 7))
			fmt.Printf("Diferenta in saptamani este de : %d saptamani\n", saptamani)
			fmt.Printf("Diffrence in luni este de : %d luni\n", diferentaLuni(dataInceput, dataSfarsit))
			fmt.Printf("Diffrence in ani este de : %d ani\n", dataSfarsit.Year()-dataInceput.Year())
		} else {
			fmt.Println(err)
		}
	}
}

func convertireTextInNumar(reader *bufio.Reader, mesaj string) (int, error) {
	fmt.Print(mesaj)
	textNumar, _ := reader.ReadString('\n')
	textNumar = strings.Replace(textNumar, "\n", "", -1)
	numar, err := strconv.Atoi(textNumar)
	return numar, err
}

func ziuaSaptamanii(reader *bufio.Reader) {
	fmt.Println("format data: aaaa-ll-zz")
	data, err := citireValidareData(reader, "Data: ")
	if err == nil {
		fmt.Println("Ziua saptamanii este: ", data.Weekday())
	} else {
		fmt.Println(err)
	}
}

func diferentaLuni(dataInceput, dataSfarsit time.Time) int {
	var luni = 0
	for dataInceput.Before(dataSfarsit) {
		dataInceput = dataInceput.AddDate(0, 1, 0)
		if dataInceput.Year() != dataInceput.Year() {
			luni++
		} else {
			if dataInceput.Month() != dataSfarsit.Month() {
				luni++
			}
		}
	}
	return luni
}

func convertireTextInData(ds string) (time.Time, error) {
	const dataReferinta = "2006-01-02"
	data, err := time.Parse(dataReferinta, ds)
	if err != nil {
		return data, err
	}
	return data, nil
}

func validareData(date string) (time.Time, error) {
	data2, err := convertireTextInData(date)
	if err != nil {
		return data2, err
	}
	return data2, nil
}

func verificareData(dataInceput, dataSfarsit time.Time) (bool, error) {
	if dataInceput.After(dataSfarsit) {
		return false, fmt.Errorf(" data de inceput este mai mare decat data de sfarsit")
	}
	return true, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var optiuneMeniu = 0
	for optiuneMeniu != 3 {
		fmt.Println("1. Diferenta intre doua date")
		fmt.Println("2. ziua saptmanii la o data anume")
		fmt.Println("3. Iesire")

		optiuneMeniu, err := convertireTextInNumar(reader,"Optiunea Dvs. : ")
		fmt.Println()

		if err != nil {
			fmt.Println(err)
		} else {
			if optiuneMeniu < 1 || optiuneMeniu > 3 {
				fmt.Println("Optiune eronata")
				optiuneMeniu = 0
			} else {
				switch optiuneMeniu {
				case 1:
					diferentaIntreDouaDate(reader)
				case 2:
					ziuaSaptamanii(reader)
				case 3:
					os.Exit(0)
				}
			}
		}
		fmt.Println()
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clearConsole(){
	cmd := exec.Command("cmd","/c","cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func verifAnBisect(an int) bool{
	if an % 4 != 0 {
		return false
	} else if an % 400 == 0 {
		return true
	} else if an % 100 == 0 {
		return false
	} else {
		return true
	}
}

type dataCalendar struct {
	an int
	luna int
	zi int
}

func dateReader() dataCalendar{
	var an,luna,zi int
	fmt.Print("An:")
	fmt.Scan(&an)
	for an < 0 {
		fmt.Scan(&an)
	}
	fmt.Print("Luna:")
	fmt.Scan(&luna)
	for luna < 0 || luna > 12 {
		fmt.Scan(&luna)
	}
	fmt.Print("Zi:")
	fmt.Scan(&zi)
	if luna == 2 {
		if verifAnBisect(an){
			for zi < 0 || zi > 29{
				fmt.Scan(&zi)
			}
		} else {
			for zi < 0 || zi > 28{
				fmt.Scan(&zi)
			}
		}
	} else if luna == 1 || luna == 3 || luna == 5 || luna == 7 || luna == 8 || luna == 10 || luna == 12{
		for zi < 0 || zi > 31{
			fmt.Scan(&zi)
		}
	} else {
		for zi < 0 || zi > 30{
			fmt.Scan(&zi)
		}
	}
	return dataCalendar{an,luna,zi}
}

func difDate(data1,data2 dataCalendar) int{
	data1zile := int(float64(data1.an)*365.242199 + float64(data1.luna)*30.4368499 + float64(data1.zi))
	data2zile := int(float64(data2.an)*365.242199 + float64(data2.luna)*30.4368499 + float64(data2.zi))
	return data1zile - data2zile
}

func difDateInZile(){
	fmt.Println("Data 1:")
	data1 := dateReader()
	fmt.Println("Data 2:")
	data2 := dateReader()
	if data1.an<data2.an{
		data1,data2 = data2,data1
	} else if data1.an == data2.an && data1.luna < data2.an{
		data1,data2 = data2,data1
	} else if data1.an == data2.an && data1.luna == data2.luna && data1.zi<data2.zi{
		data1,data2 = data2,data1
	}
	difDate := difDate(data1,data2)
	if difDate < 0{
		difDate = -difDate
	}
	fmt.Println("Diferenta dintre cele doua dati este de",difDate,"zile.")
	fmt.Println("Apasati tasta Enter pentru a continua...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clearConsole()
	meniu()
}

func difDateInALSZ(){
	fmt.Println("Data 1:")
	data1 := dateReader()
	fmt.Println("Data 2:")
	data2 := dateReader()
	if data1.an<data2.an{
		data1,data2 = data2,data1
	} else if data1.an == data2.an && data1.luna < data2.an{
		data1,data2 = data2,data1
	} else if data1.an == data2.an && data1.luna == data2.luna && data1.zi<data2.zi{
		data1,data2 = data2,data1
	}
	difAni := data1.an - data2.an
	difLuni := data1.luna - data2.luna
	if difLuni < 0 {
		difLuni += 12
		difAni--
	}
	difZile := data1.zi - data2.zi
	if difZile < 0 {
		difLuni--
		difDate := difDate(data1,data2)
		if difDate < 0{
			difDate = -difDate
		}
		difZile = difDate - int(float64(difAni)*365.242199) - int(float64(difLuni)*30.4368499)
	}
	difSapt := 0
	for difZile >= 7{
		difZile -= 7
		difSapt++
	}
	fmt.Println("Diferenta dintre cele doua dati este de",difAni,"ani,",difLuni,"luni,",difSapt,"saptamani si",difZile,"zile.")
	fmt.Println("Apasati tasta Enter pentru a continua...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clearConsole()
	meniu()
}

func dataPlusMinusNrZile(semn bool){
	fmt.Println("Data:")
	data := dateReader()
	timeDate := time.Date(data.an,time.Month(data.luna),data.zi,0,0,0,0,time.UTC)
	if semn{
		//data PLUS un numar de zile
		fmt.Print("Zile de adaugat:")
		var zilePlus int
		fmt.Scan(&zilePlus)
		for zilePlus<0{
			fmt.Scan(&zilePlus)
		}
		timeDate = timeDate.AddDate(0,0,zilePlus)
	} else {
		//data MINUS un numar de zile
		fmt.Print("Zile de scazut:")
		var zileMinus int
		fmt.Scan(&zileMinus)
		for zileMinus<0{
			fmt.Scan(&zileMinus)
		}
		timeDate = timeDate.AddDate(0,0,-zileMinus)
	}
	fmt.Println("Noua data:")
	fmt.Println("An:",timeDate.Year())
	fmt.Println("Luna:",int(timeDate.Month()))
	fmt.Println("Zi:",timeDate.Day())
	fmt.Println("Apasati tasta Enter pentru a continua...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clearConsole()
	meniu()
}

func ziSaptDataCalendar(){
	fmt.Println("Data:")
	data := dateReader()
	timeDate := time.Date(data.an,time.Month(data.luna),data.zi,0,0,0,0,time.UTC)
	ziSapt := timeDate.Weekday().String()
	switch ziSapt {
	case "Monday":
		ziSapt = "Luni"
		break
	case "Tuesday":
		ziSapt = "Marti"
		break
	case "Wednesday":
		ziSapt = "Miercuri"
		break
	case "Thursday":
		ziSapt = "Joi"
		break
	case "Friday":
		ziSapt = "Vineri"
		break
	case "Saturday":
		ziSapt = "Sambata"
		break
	case "Sunday":
		ziSapt = "Duminica"
		break
	default:
		ziSapt = "Nu exista in calendar!"
		break
	}
	fmt.Println("Data introdusa se regaseste in ziua de:",ziSapt)
	fmt.Println("Apasati tasta Enter pentru a continua...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clearConsole()
	meniu()
}

func meniu(){
	fmt.Println("Operatii cu date calendaristice:")
	fmt.Println("1) Diferenta dintre doua date in zile")
	fmt.Println("2) Diferenta dintre doua date in ani, luni, saptamani si zile")
	fmt.Println("3) Data + un numar de zile")
	fmt.Println("4) Data - un numar de zile")
	fmt.Println("5) In ce zi a saptamanii se gaseste o anumita data")
	fmt.Println("0) Iesire din aplicatie")
	fmt.Println("Alegeti o optiune:")
	alegereOpt()
}

func alegereOpt(){
	var opt int
	fmt.Scan(&opt)
	switch opt {
	case 0:
		os.Exit(0)
	case 1:
		difDateInZile()
		break
	case 2:
		difDateInALSZ()
		break
	case 3:
		dataPlusMinusNrZile(true)
		break
	case 4:
		dataPlusMinusNrZile(false)
		break
	case 5:
		ziSaptDataCalendar()
		break
	default:
		fmt.Println("Optiune inexistenta!")
		fmt.Println("Apasati tasta Enter pentru a continua...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		clearConsole()
		meniu()
		break
	}
}

func main(){
	meniu()
}

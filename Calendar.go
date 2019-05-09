package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func AnBisect() bool{
	var an int
	fmt.Print("An:")
	fmt.Scan(&an)
	if an % 4 != 0 {
		return false
	} else if an % 400 == 0 {
		println("Da,este an bisec")
		meniu()
		return true
	} else if an % 100 == 0 {
		println("Nu, nu este an bisec")
		meniu()
		return false
	} else {
		println("Da,este an bisec")
		meniu()
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
		if AnBisect(){
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

func adaugareZile(){
	fmt.Println("Data:")
	data := dateReader()
	timeDate := time.Date(data.an,time.Month(data.luna),data.zi,0,0,0,0,time.UTC)
	fmt.Print("Zile de adaugat:")
	var zilePlus int
	fmt.Scan(&zilePlus)
	for zilePlus<0{
		fmt.Scan(&zilePlus)
	}
	timeDate = timeDate.AddDate(0,0,+zilePlus)
	fmt.Println("Noua data:")
	fmt.Println("An:",timeDate.Year())
	fmt.Println("Luna:",int(timeDate.Month()))
	fmt.Println("Zi:",timeDate.Day())
	meniu()
}

func scadereZile(){
	fmt.Println("Data:")
	data := dateReader()
	timeDate := time.Date(data.an,time.Month(data.luna),data.zi,0,0,0,0,time.UTC)

		//data MINUS un numar de zile
		fmt.Print("Zile de scazut:")
		var zileMinus int
		fmt.Scan(&zileMinus)
		for zileMinus<0{
			fmt.Scan(&zileMinus)
		}
		timeDate = timeDate.AddDate(0,0,-zileMinus)

	fmt.Println("Noua data:")
	fmt.Println("An:",timeDate.Year())
	fmt.Println("Luna:",int(timeDate.Month()))
	fmt.Println("Zi:",timeDate.Day())
	meniu()
}

func ziSaptamana(){
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
	meniu()
}

func meniu(){
	fmt.Println("1) E an bisec?")
	fmt.Println("2) Adaugare numar de zile")
	fmt.Println("3) Scadere numar de zile")
	fmt.Println("4) Ziua saptamanii in care se gaseste o data")
	fmt.Println("Alegeti o optiune:")

	var opt int
	fmt.Scan(&opt)
	switch opt {
	case 1:
		AnBisect()
		break
	case 2:
		adaugareZile()
		break
	case 3:
		scadereZile()
		break
	case 4:
		ziSaptamana()
		break
	default:
		fmt.Println("Aceasta optiune nu exista!")
		meniu()
		break
	}
}

func main(){
	meniu()
}
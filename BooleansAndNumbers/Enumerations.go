package main

import (
	"fmt"
	"strings"
)

func main() {

	flag := Active | Send

	fmt.Println("Which day it is? \n",Monday)

	fmt.Printf("Which day it is? \n", Weekday(7))

	fmt.Println("\n I want Cyan \n" ,Cyan,Yellow)

	fmt.Println(BitFlag(0), Active,Send,flag ,Receive, flag|Receive)

}

// Declare a new type named Weekday which will unify our enum values
// It has an underlying type of unsigned integer (uint).
type Weekday int

type Color int

// Declare typed constants each with type of Weekday
const (
	Sunday    Weekday = 0
	Monday    Weekday =  1//iota
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

const (
	Cyan Color = iota
	Magenta
	Yellow
	Black
)

type BitFlag int
const (
	Active BitFlag = 1 << iota // 1 << 0 == 1
	Send // Implicitly BitFlag = 1 << iota // 1 << 1 == 2
	Receive // Implicitly BitFlag = 1 << iota // 1 << 2 == 4
)



// String returns the name of the day
func (day Weekday) String() string {
	names := []string{"Sunday", "Monday", "Tuesday", "Wednesday",
		"Thursday", "Friday", "Saturday"}

	// prevent panicking in case of Weekday is out-of-range
	if day < Sunday || day > Saturday {
		return "Unknown"
	}

	return names[day]
}

//func (aux Color ) int() int{
//
//	colors := [...]int { 1,3,4,5,6}
//	if(aux != 3){
//		return 2;
//	}
//	return colors[aux];
//}

func (flag BitFlag) String() string {
	var flags []string
	if flag&Active == Active {
		flags = append(flags, "Active" )
		println("am intrat in 1")
	}
	if flag&Send == Send {
		flags = append(flags, "Send" )
		println("am intrat in 2")
	}
	if flag&Receive == Receive {
		flags = append(flags, "Receive" )
		println("am intrat in 3")
	}
	if len(flags) > 0 { // int(flag) is vital to avoid infinite recursion!
		return fmt.Sprintf( "%d(%s)" , int(flag), strings.Join(flags, "|" ))
	}
	return "0()"
}

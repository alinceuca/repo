//01. Sa se determine numarul format din x cu p biti incepand cu al t-lea numarat de la dreapta la stanga

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	number, err := readAndValidateUInt(reader,"x= ")
	if err != nil{
		fmt.Println(err)
	} else {
		p, errP := readAndValidateUInt(reader,"p= ")
		if errP != nil{
			fmt.Println(errP)
		} else {
			t, errT := readAndValidateUInt(reader,"t= ")
			if errT != nil{
				fmt.Println(errT)
			} else {
				x := number >> t
				x = x & (1<<p - 1)
				fmt.Println(x)
			}
		}
	}
}

func readAndValidateUInt(reader *bufio.Reader, message string) (uint64, error) {
	fmt.Print(message)
	numberStr, _ := reader.ReadString('\n')
	numberStr = strings.Replace(numberStr, "\n", "", -1)
	number, err := strconv.ParseUint(numberStr, 10, 64)

	return number, err
}
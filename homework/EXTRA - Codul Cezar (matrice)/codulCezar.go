package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func citireText() string {
	fmt.Print("Introduceti un text: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func CezarMatrice(text string) string {
	cheieCezar := [10][10] byte{
		{48, 49, 50, 51, 52, 53, 54, 55, 56, 57},
		{97, 65, 98, 66, 99, 67, 100, 68, 101, 69},
		{102, 70, 103, 71, 104, 72, 105, 73, 106, 74},
		{107, 75, 108, 76, 109, 77, 110, 78, 111, 79},
		{112, 80, 113, 81, 114, 82, 115, 83, 116, 84},
		{117, 85, 118, 86, 119, 87, 120, 88, 121, 89},
		{122, 90, 46, 58, 44, 59, 63, 33, 39, 34},
		{43, 45, 42, 47, 92, 124, 38, 36, 35, 37},
		{40, 41, 91, 93, 123, 125, 32, 95, 126, 96},
		{60, 61, 62, 94, 64, 174, 175, 236, 251, 227}}
	var textCriptat string
	iter := 0
	for iter < len(text) {
		var iPozCh1, jPozCh1, iPozCh2, jPozCh2 int
		ch1 := text[iter]
		if ch1 == 32 || iter == len(text)-1 {
			textCriptat += string(ch1)
			iter++
		} else {
			ch2 := text[iter+1]
			if ch2 == 32 {
				textCriptat += string(ch1)
				textCriptat += string(ch2)
				iter += 2
			} else {
				ok1, ok2 := false, false
				for i := 0; i < len(cheieCezar); i++ {
					for j := 0; j < len(cheieCezar[0]); j++ {
						if cheieCezar[i][j] == ch1 && !ok1 {
							ok1 = true
							iPozCh1 = i
							jPozCh1 = j
						}
						if cheieCezar[i][j] == ch2 && !ok2 {
							ok2 = true
							iPozCh2 = i
							jPozCh2 = j
						}
						if ok1 && ok2 {
							break
						}
					}
					if ok1 && ok2 {
						break
					}
				}
				if iPozCh1 == iPozCh2 || jPozCh1 == jPozCh2 {
					ch1, ch2 = ch2, ch1
				} else {
					ch1 = cheieCezar[iPozCh2][jPozCh1]
					ch2 = cheieCezar[iPozCh1][jPozCh2]
				}
				textCriptat += string(ch1)
				textCriptat += string(ch2)
				iter += 2
			}
		}
	}
	return textCriptat
}

func main(){
	textCriptat := CezarMatrice(citireText())
	textDecriptat := CezarMatrice(textCriptat)
	fmt.Println("Textul criptat:",textCriptat)
	fmt.Println("Textul decriptat:",textDecriptat)
}

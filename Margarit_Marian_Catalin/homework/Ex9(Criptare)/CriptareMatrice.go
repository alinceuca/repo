package main

import (
	"bytes"
	"fmt"
	"strings"
)

var matrice = [6][6] rune{
	{'1', '2', '3', '4', '5', '6'},
	{'A', 'B', 'C', 'D', 'E', '7'},
	{'F', 'G', 'H', 'I', 'J', '8'},
	{'K', 'L', 'M', 'N', 'O', '9'},
	{'P', 'R', 'S', 'T', 'U', '0'},
	{'V', 'W', 'X', 'Y', 'Z', 'Q'},
}

func findPosition(character rune) (int,int){

	for i:=0; i< 6; i++   {
		for j:=0; j<6 ; j++  {
			if matrice[i][j] == character {
				return i,j
			}
		}
	}
	return 0, 0
}

func criptare(message string) string {
	var encodedMessage bytes.Buffer
	for i:=0 ; i+1< len(message); i+=2  {
		i1, j1:= findPosition(rune(message[i]))
		i2, j2 := findPosition(rune(message[i+1]))
		if i1 == i2 {
			encodedMessage.WriteString(string(matrice[i1][j2]))
			encodedMessage.WriteString(string(matrice[i2][j1]))
		} else {
			encodedMessage.WriteString(string(matrice[i2][j1]))
			encodedMessage.WriteString(string(matrice[i1][j2]))
		}

	}
	if len(message) %2 != 0 {
		encodedMessage.WriteString(string(message[len(message) -1]))
	}
	fmt.Println(string(encodedMessage.String()))
	return string(encodedMessage.String())
}
func main(){

	var message string
	fmt.Println("Mesaj: ")
	fmt.Scan(&message)
	criptare(criptare(strings.ToUpper(message)))


}
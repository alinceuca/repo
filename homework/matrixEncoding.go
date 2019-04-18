package main

import (
	"bytes"
	"fmt"
	"strings"
)

var matrix = [6][6] rune{
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
			if matrix[i][j] == character {
				return i,j
			}
		}
	}
	return 0, 0
}

func encode(message string) string {

	var encodedMessage bytes.Buffer
	for i:=0 ; i+1< len(message); i+=2  {
		i1, j1:= findPosition(rune(message[i]))
		i2, j2 := findPosition(rune(message[i+1]))
		if i1 == i2 {
			encodedMessage.WriteString(string(matrix[i1][j2]))
			encodedMessage.WriteString(string(matrix[i2][j1]))
		} else {
			encodedMessage.WriteString(string(matrix[i2][j1]))
			encodedMessage.WriteString(string(matrix[i1][j2]))
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
	fmt.Println("Introduceti mesaj: ")
	fmt.Scan(&message)


	encode(encode(strings.ToUpper(message)))


}

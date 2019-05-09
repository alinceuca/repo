package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var encodeMatrix = [6][6] rune{
	{'1', '2', '3', '4', '5', '6'},
	{'A', 'B', 'C', 'D', 'E', '7'},
	{'F', 'G', 'H', 'I', 'J', '8'},
	{'K', 'L', 'M', 'N', 'O', '9'},
	{'P', 'R', 'S', 'T', 'U', '0'},
	{'V', 'W', 'X', 'Y', 'Z', ' '},
}


func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("The text you want to encode is: ")
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(matrixEncode(strings.ToUpper(text)))

}

func matrixEncode(text string) string {

	var encodedText []rune
	for i := 0; i < len(text)-1; i += 2 {
		first := text[i]
		second := text[i+1]

		firstI, firstJ := findPosition(rune(first))
		secondI, secondJ := findPosition(rune(second))
		if firstJ == secondJ {
			encodedText = append(encodedText, rune(second), rune(first))
		} else {
			encodedText = append(encodedText, encodeMatrix[firstI][secondJ])
			encodedText = append(encodedText, encodeMatrix[secondI][firstJ])
		}

	}
	if len(text)%2 != 0 {
		encodedText = append(encodedText, rune(text[len(text)-1]))
	}

	return string(encodedText)
}

func findPosition(character rune) (int,int) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if encodeMatrix[i][j] == character {
				return i, j
			}
		}
	}
	return 0, 0
}
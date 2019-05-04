package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
	"unicode"
)

const howManyNumbers = 100
const shiftsFileName = "prob9.bin"
const encodeFileName = "encode.txt"
const decodeFileName = "decode.txt"

func main() {

	writeBinaryFile()
	shifts := readBinaryFile()
	text := readTextToEncode()
	fmt.Println(text)
	fmt.Println(shifts)
	writeEncodedAndDecodedText(text, shifts)

}

func writeBinaryFile() {
	file, err := os.Create(shiftsFileName)
	defer file.Close()
	check(err)

	rand.Seed(time.Now().UnixNano())
	bin_buf := new(bytes.Buffer)

	for i := 0; i < howManyNumbers; i++ {
		s := randomByte(0,255)
		binary.Write(bin_buf, binary.BigEndian, s)

		//fmt.Printf(" s = %v\n", s)
	}
	writeBytes(file, bin_buf.Bytes())
}

func writeBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	check(err)
}

func randomByte(min, max int) byte {
	return byte(min + rand.Intn(max-min))
}

func readBinaryFile() []byte{
	var shifts []byte
	file, err := os.Open(shiftsFileName)
	defer file.Close()
	check(err)

	var m byte
	data := readBytes(file)
	buffer := bytes.NewBuffer(data)

	for i := 0; i < howManyNumbers; i++ {
		err = binary.Read(buffer, binary.BigEndian, &m)
		check(err)
		shifts = append(shifts, m)
		fmt.Println(m)
	}

	return shifts
}

func readBytes(file *os.File) []byte {
	bytes := make([]byte, howManyNumbers)

	_, err := file.Read(bytes)
	check(err)

	return bytes
}

func readTextToEncode() string{
	text, err := ioutil.ReadFile(encodeFileName)
	check(err)

	return string(text)
}

func writeEncodedAndDecodedText(text string, shifts []byte) {
	file, err := os.Create(decodeFileName)
	defer file.Close()
	check(err)

	fmt.Fprintln(file, "Initial text:")
	fmt.Fprintln(file, text)

	fmt.Fprintln(file, "Encoded text:")
	encodedText := getEncodeText(text, shifts)
	fmt.Fprintln(file, encodedText)

	fmt.Fprintln(file, "Decoded text:")
	decodedText := getDecodeText(encodedText, shifts)
	fmt.Fprintln(file, decodedText)

	fmt.Println("file appended successfully")
}

func getEncodeText(text string, shifts []byte) string {
	var encodedText bytes.Buffer

	i := 0
	for _, character := range text {
		if !unicode.IsLetter(character) {
			encodedText.WriteString(string(character))
		} else {

			shift := shifts[i]%26
			characterAscii := int(character)+int(shift)
			if unicode.IsUpper(character) {
				if characterAscii > 90 {
					encodedText.WriteString(string(65-90+characterAscii))
				} else {
					encodedText.WriteString(string(characterAscii))
				}
			} else {
				if characterAscii > 122{
					encodedText.WriteString(string(97-122+characterAscii))
				} else {
					encodedText.WriteString(string(characterAscii))
				}
			}

			i++
		}
	}

	return encodedText.String()
}

func getDecodeText(text string, shifts []byte) string {
	var decodedText bytes.Buffer

	i := 0
	for _, character := range text {
		if !unicode.IsLetter(character) {
			decodedText.WriteString(string(character))
		} else {

			shift := shifts[i] % 26
			characterAscii := int(character)-int(shift)
			if unicode.IsUpper(character) {
				if characterAscii < 65 {
					decodedText.WriteString(string(90-65+characterAscii))
				} else {
					decodedText.WriteString(string(characterAscii))
				}
			} else {
				if characterAscii < 97{
					decodedText.WriteString(string(122-97+characterAscii))
				} else {
					decodedText.WriteString(string(characterAscii))
				}
			}

			i++
		}
	}

	return decodedText.String()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
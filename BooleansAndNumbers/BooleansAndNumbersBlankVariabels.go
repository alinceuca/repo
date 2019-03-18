package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {


	a := "A"
	// in acest caz nu as fi putut rula intrucat aveam aroare la compilare
	//_ := "this is b"

	// aceasta ar fi utilitatea blank-ului, dar trebuie sa mentionez ca programul
	// nu va rula, dar vom primi un mesaj de eroare in care ni se va spune
	// variabila din partea stanga nu este definita


	fmt.Println("am ales", a)

	res, err := http.Get("https://www.google.com/")

	//res, _ := http.Get("https://www.google.com/")
	page,_ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println( page)



	//if err != nil{
	//	log.Fatal(err)
	//}
	//page, err := ioutil.ReadAll(res.Body)
	//res.Body.Close()
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println("%s",page)
	
}


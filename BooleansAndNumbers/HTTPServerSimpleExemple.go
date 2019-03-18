package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Println(w," This is my super dupper web page! 2+2=4 Quick maths")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Println(w," This is aboute page")
}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8880",nil)

}

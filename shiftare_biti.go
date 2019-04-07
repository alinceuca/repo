package main

import "fmt"

func main() {
   var numar,pbiti,mutareT uint
   fmt.Scan(&numar)
   fmt.Scan(&pBiti)
   fmt.Scan(&mutareT)
   fmt.Print(mutare(numar,pbiti,mutareT))

}

func mutare(numar,pBiti,mutareT uint) uint  {
   q := numar >> mutareT
   q = q & (1<<pBiti-1)
   return q
}
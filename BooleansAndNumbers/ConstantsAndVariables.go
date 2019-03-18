package main

func main() {

	const limit = 512 // constant; type-compatible with any number
	const top uint16 = 1421 // constant; type: uint16
	start := -19 // variable; inferred type: int
	end := int64( 9876543210) // variable; type: int64
	var i int // variable; value 0; type: int
	var debug = false // variable; inferred type: bool
	checkResults := true // variable; inferred type: bool
	stepSize := 1.5 // variable; inferred type: float64
	acronym := "FOSS" // variable; inferred type: string
	var emptyString string // variable; value null; type: string

	if len(emptyString) > 0 {
		println("it isn't a empty string")
	}

	if emptyString == ""{
		println("it is a empty string")
	}

	println(start,end,i,debug,checkResults,stepSize,acronym,emptyString)
	
}

package main

import (
	"bytes"
	"fmt"
	"math"
)

func main() {

/*	//Correct form
	// Compiler will conveniently insert semicolons automatically at the end of nonblank lines t
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}*/


/*	// WRONG! (This won't compile.) ✗
	//Compiler will add semicolon after i++
	for i := 0; i < 5; i++
	{
		fmt.Println(i)
	}*/

/*	//This compiles
	var i int = 0
	i++
	fmt.Println("My variable is", i)

	//This will compile in C/C++, but not in Golang
	i=i++
	fmt.Println("My new variable is", i)*/

	// := will declare new variable and assign a value,
	// if there is a variable that already exists it will be assign to
	// unless := is inside a new scope such as IF or FOR
	// shadow variables
/*	a, b, c := 2, 3, 5
	for a := 7; a < 8; a++ { // a is unintentionally shadowing the outer a
		b := 11 // b is unintentionally shadowing the outer b
		c = 13 // c is the intended outer c ✓
		fmt.Printf("inner: a→%d b→%d c→%d\n", a, b, c)
	}
	fmt.Printf("outer: a→%d b→%d c→%d\n", a, b, c)*/


/*	type StringSlice []string
	fancy := StringSlice{"Lithium", "Sodium", "Potassium", "Rubidium"}
	fmt.Println(fancy)
	plain := []string(fancy)
	//plain := fancy
	plain[0]="Test"
	fmt.Println("test")
	fmt.Println(plain)
	fmt.Println(fancy)*/

/*	var i interface{} = 99.5
	var s interface{} = []string{"left", "right"}
	j := i.(float64) // j is of type int (or a panic() has occurred)
	fmt.Printf("%T→%f\n", j, j)
	if i, ok := i.(int); ok {
		fmt.Printf("%T→%f\n", i, i) // i is a shadow variable of type int
	}
	if s, ok := s.([]string); ok {
		fmt.Printf("%T→%q\n", s, s) // s is a shadow variable of type []string
	}*/


	//classifier(5, -17.9, "ZIP", nil, true, complex(1, 1))

/*	MA := []byte(`{"name": "Massachusetts", "area": 27336, "water": 25.7,
"senators": ["John Kerry", "Scott Brown"]}`)
	var object interface{}
	if err := json.Unmarshal(MA, &object); err != nil {
		fmt.Println(err)
	} else {
		jsonObject := object.(map[string]interface{})
		fmt.Println(jsonObjectAsString(jsonObject))
	}*/


	/*counterA := createCounter(2) // counterA is of type chan int
	counterB := createCounter(102) // counterB is of type chan int
	for i := 0; i < 5; i++ {
		a := <-counterA
		fmt.Printf("(A→%d, B→%d) ", a, <-counterB)
	}
	fmt.Println()*/


	fmt.Printf("Convert int64 to int value is %v ",ConvertInt64ToInt(1000000000000))
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is an int\n", i)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is an unsigned int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d's type is unknown\n", i)
		}
	}
}

func jsonObjectAsString(jsonObject map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	comma := ""
	for key, value := range jsonObject {
		buffer.WriteString(comma)
		switch value := value.(type) { // shadow variable
		case nil:
			fmt.Fprintf(&buffer, "%q: null", key)
		case bool:
			fmt.Fprintf(&buffer, "%q: %t", key, value)
		case float64:
			fmt.Fprintf(&buffer, "%q: %f", key, value)
		case string:
			fmt.Fprintf(&buffer, "%q: %q", key, value)
		case []interface{}:
			fmt.Fprintf(&buffer, "%q: [", key)
			innerComma := ""
			for _, s := range value {
				if s, ok := s.(string); ok { // shadow variable
					fmt.Fprintf(&buffer, "%s%q", innerComma, s)
					innerComma = ", "
				}
			}
			buffer.WriteString("]")
		}
		comma = ", "
	}
	buffer.WriteString("}")
	return buffer.String()
}

func createCounter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)
	return next
}

func ConvertInt64ToInt(x int64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		return int(x)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", x))
}

func IntFromInt64(x int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	i = ConvertInt64ToInt(x)
	return i, nil
}
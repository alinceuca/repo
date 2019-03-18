package main

import (
	"fmt"
	"math"
)

func main() {

	const factor= 3		  // factor is compatible with any numeric type
	i := 20000			  // i is of type int by inference
	i *= factor
	j := int16(20)       // j is of type int16; same as: var j int16 = 20
	i += int(j)          // Types must match so conversion is required
	k := uint8(0)        // Same as: var k uint8
	k = uint8(i)         // Succeeds, but k's value is truncated to 8 bits ✗
	fmt.Println(i, j, k) // Prints: 60020 20 116
	fmt.Println(Uint8FromInt(i))
}

func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf( "%d is out of the uint8 range" , x)
}
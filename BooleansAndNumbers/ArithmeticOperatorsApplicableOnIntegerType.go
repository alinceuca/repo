package main

func main() {

	var x, y, u uint
	x = 300
	y = 10
	u = 2
	//x := 3
	//y := 10
	//
	//u :=1

	println(^x)
	x %= y
	println(x)
	x &= y
	println(x)
	x |= y
	println(x)
	x ^= y
	println(x)
	x = 100
	x <<= u
	println(x)
	x = 300
	println(x << u)

	println(x |y)

	//int8(200) // eroare

	println(int8(u))

	println(float64(x))

}

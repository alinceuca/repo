package main
func searchNumber(x int, p uint, t uint) int {
	x = x >> t
	x = x & (1<<p - 1)

	return x
}

func main(){
	searchNumber(3,8,4)
}

package syntax

import "fmt"

func main() {

	//Staic Arrays
	var a [2]int
	a[0] = 1
	a[1] = 2

	//Dynamic Arrays
	l := make([]int, 0)
	l = append(l, 1)
	l = append(l, 2)

	//Maps
	m := make(map[int]int)
	m[0] = 1
	m[1] = 2
	fmt.Println(m[1])
}

package main

import (
	"fmt"
	"goproj/newpackage"
)

func main() {
	var n int
	fmt.Println("Enter the value of n")
	fmt.Scan(&n)

	v := make([]newpackage.Car, n)

	for i := 0; i < n; i++ {
		var rno int
		var ms string
		var c newpackage.Car
		fmt.Println("For the car", i+1, "Enter the value of rno and ms")
		fmt.Scan(&rno)
		fmt.Scan(&ms)
		//	fmt.Println("ms is ", ms)
		c.SetRno(rno)
		c.SetMs(ms)
		v[i] = c
	}

	for i := 0; i < n; i++ {
		fmt.Println("For the car ", i+1, " the rno is ", v[i].GetRno(), " and ms is ", v[i].GetMs())
	}
}

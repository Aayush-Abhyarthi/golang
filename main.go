package main

import (
	"fmt"
	"golang/newpackage"
)

func main() {

	var n int

	fmt.Println("Enter the value of n")
	fmt.Scan(&n)

	v := make([]newpackage.Person, n)
	for i := 0; i < n; i++ {

		var x newpackage.Person
		var id int
		var height float32
		var name string

		fmt.Println("Enter the value of ", i+1, " person")
		fmt.Scan(&id, &height, &name)
		x.SetID(id)
		x.SetHeight(height)
		x.SetName(name)

		v[i] = x
	}

	var m int
	fmt.Println("Enter the value of m")
	fmt.Scan(&m)

	w := make([]newpackage.Car, m)
	for i:=0;i<m;i++{

		var x newpackage.Car
		var y int
		fmt.Println("Enter the value of ", i+1, " car")
		fmt.Scan(&y)
		x.SetId(y)
		w[i]=x
	}

	for i:=0;i<n;i++{
		fmt.Println("The value of ", i+1, " person is")
		fmt.Println(v[i].GetID(), v[i].GetHeight(), v[i].GetName())
	}

	for i:=0;i<m;i++{
		fmt.Println("The value of the ", i+1, "car is", w[i].GetId() )
	}

}

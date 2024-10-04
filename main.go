package main

import "fmt"

type Person struct { // Struct definition
	Name     string `json:"name"`
	PersonId int    `json:"personid"`
}

func (p Person) printPersonValue() {
	fmt.Println("The person name is : ", p.Name, " and the person id is : ", p.PersonId)
}

func basicVariableDefinition() {

	var x = 2 // Normal variable declaration
	fmt.Println("The value of the variable x is ", x)

	var n int
	fmt.Scan(&n)

	var a = make([]int, n) // Slices
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}

}

func main() {

	basicVariableDefinition()

}

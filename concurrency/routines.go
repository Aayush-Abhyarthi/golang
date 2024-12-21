package main

import "fmt"

func fun1(a int, ch chan bool) {
	for i := 1; i <= a; i++ {
		fmt.Println("fun1 : ", i)
	}
	ch <- true
}

func fun2(a int, ch chan bool) {
	for i := 1; i <= a; i++ {
		fmt.Println("fun2 : ", i)
	}
	ch <- true
}

func main() {

	ch := make(chan bool, 2)
	go fun1(100, ch)
	go fun2(100, ch)
	<-ch
	<-ch

}

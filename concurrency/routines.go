package main

import "fmt"

func countDigits(ch chan bool){
	for i:=1;i<=26;i++{
		fmt.Println(i)
	}
	ch <- true
}

func countAlpha(ch chan bool){
	for i:='a';i<='z';i++{
		fmt.Printf("%c\n", i)
	}
	ch <- true
}

func main(){

	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go countDigits(ch1)
	go countAlpha(ch2)
	<- ch1
	<- ch2

}
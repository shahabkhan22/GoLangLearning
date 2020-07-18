package main

import "fmt"

//receives a channel and prints
func returns(c chan string) {
	fmt.Println("Hello " + <-c)
}

//receives a channel and writes a channel to it
func re_returns(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

func main() {

	//fmt.Println("Main Started")
	cc := make(chan chan string) //making channel of type channel of string

	go re_returns(cc)

	c := <-cc //receives a channel from re_returns goroutine

	go returns(c)

	c <- "ABC"

	//fmt.Println("Main Stopped")
}

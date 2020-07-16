package main

import "fmt"

func sendValues(myIntChannel chan int) {
	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
		//fmt.Println("Sent value ", i)
	}
	close(myIntChannel) //solution for DEADLOCK due to i < 6 in main()
}

func main() {
	myIntChannel := make(chan int)
	go sendValues(myIntChannel) //goroutine created that starts sending values

	for i := 0; i < 6; i++ { //change to i < 6 and see a DEADLOCK
		fmt.Println(<-myIntChannel) //receiving value
		//fmt.Println("Received value ", i)
	}
}

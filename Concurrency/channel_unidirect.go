package main

import "fmt"

func result(rec <-chan string) {
	fmt.Println("Hello " + <-rec)
	//rec <- "Done"
	//this is invalid and gives following error-
	//invalid operation: rec <- "Done" (send to receive-only type <-chan string)
}

func main() {
	rchan := make(<-chan int)
	schan := make(chan<- int)
	mychan := make(chan string)

	fmt.Printf("Type of a Channel : `%T`\n", mychan)
	fmt.Printf("Type of Receive Channel : `%T`\n", rchan)
	fmt.Printf("Type of Send Channel : `%T`\n", schan)

	//Any Bi-directional channel can be easily converted into unidirectional channel

	go result(mychan)
	mychan <- "ABC"

}

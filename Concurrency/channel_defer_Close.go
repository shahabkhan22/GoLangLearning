package main

import "fmt"

func sendValues(myIntChannel chan int) {

	for i := 0; i < 5; i++ {
		myIntChannel <- i
	}
}

func main() {
	myIntChannel := make(chan int)
	defer close(myIntChannel)
	go sendValues(myIntChannel)

	for i := 0; i < 5; i++ {
		fmt.Println(<-myIntChannel)
	}

}

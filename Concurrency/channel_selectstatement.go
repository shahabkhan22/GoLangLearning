package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(mychan chan string) {
	time.Sleep(3 * time.Second)
	mychan <- "Hello from Service 1"
}

func service2(mychan chan string) {
	time.Sleep(5 * time.Second)
	mychan <- "Hello from Service 2"
}

func main() {
	/*
		chan1 := make(chan string)
		chan2 := make(chan string)

		go service1(chan1)
		go service2(chan2)

		select {
		case res := <-chan1:
			fmt.Println("Response from Service 1 :", res, time.Since(start))
		case res := <-chan2:
			fmt.Println("Response from Service 2 :", res, time.Since(start))
		}

		fmt.Println("Main Ended")

		//stimulating a scenario where all cases are non-blocking and
		//response is availabe at the same time using buffered channel
	*/
	fmt.Println("Main Started", time.Since(start))

	chan3 := make(chan string, 2)
	chan4 := make(chan string, 2)

	chan3 <- "Value 1"
	chan3 <- "Value 2"
	chan4 <- "Value 1"
	chan4 <- "Value 2"

	select {
	case res := <-chan3:
		fmt.Println("Response from Chan3 ", res, time.Since(start))
	case res := <-chan4:
		fmt.Println("Response from Chan4 ", res, time.Since(start))
	}

	fmt.Println("Main Ended", time.Since(start))

}

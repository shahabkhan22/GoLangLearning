package main

import "fmt"

func main() {
	mych1 := make(chan int)
	mych2 := make(chan int)
	mych3 := make(chan int)

	go func() { //in deadlock as no send operation that sends to the channels
		<-mych1
	}()

	go func() {
		mych2 <- 20
	}()

	go func() { //in deadlock as no send operation that sends to the channels
		<-mych3
	}()

	fmt.Println(<-mych2)
}

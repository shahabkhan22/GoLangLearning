package main

import "fmt"

func main() {
	mychan := make(chan int, 2)
	mychan <- 10
	mychan <- 20
	mychan <- 30 //this will cause DEADLOCK as buffer/capacity of mychan is 2
	fmt.Println(<-mychan)
}

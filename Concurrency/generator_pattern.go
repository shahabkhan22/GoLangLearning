package main

import "fmt"

func foo() <-chan string { //returns channel of type string
	mychan := make(chan string)

	go func() {
		for i := 0; ; i++ {
			mychan <- fmt.Sprintf("%s %d", "Counter at : ", i)
		}
	}()

	return mychan
}

func main() {
	mychan := foo()

	for i := 0; i < 5; i++ {
		fmt.Printf("%q\n", <-mychan)
	}

	fmt.Println("Main Ended")
}

package main

import "fmt"

func updatePosition(name string) <-chan string {
	positionChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			positionChannel <- fmt.Sprintf("%s %d", name, i)
		}
	}()

	return positionChannel
}

func main() {

	//launching a go routine from inside a function
	//returning a channel, function enables us to communicate with the service it provides
	positionChannel1 := updatePosition("ABC: ")
	positionChannel2 := updatePosition("PQR : ")

	//we can have more than one instance of the function
	//still, both the reading from channel statments is line 28,29 are blocking each other

	for i := 0; i < 5; i++ {
		fmt.Println(<-positionChannel1)
		fmt.Println(<-positionChannel2)
	}

	fmt.Println("MAIN ENDED")
}

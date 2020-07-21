//improving upon the code in generator_pattern2.go as it was waiting for each other
//to complete before proceedin to get next value for themselves
//we will now combine the imputs from both channels and send them through a single channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func updatePosition(name string) <-chan string {
	positionChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			positionChannel <- fmt.Sprintf("%s %d", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return positionChannel
}

func fanIn(mychan1, mychan2 <-chan string) <-chan string {
	mychan := make(chan string)

	go func() {
		for {
			mychan <- <-mychan1
		}
	}()

	go func() {
		for {
			mychan <- <-mychan2
		}
	}()

	return mychan
}

func main() {
	positionChannel := fanIn(updatePosition("ABC : "), updatePosition("PQR : "))

	for i := 0; i < 10; i++ {
		fmt.Println(<-positionChannel)
	}

	fmt.Println("Now we get positions for both ABC and PQR")
}

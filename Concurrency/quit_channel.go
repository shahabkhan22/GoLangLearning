package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Race(channel, quit chan string, i int) {
	channel <- fmt.Sprintf("Car %d started!", i)

	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		quit <- fmt.Sprintf("Car %d Wins the Race!", i)
		fmt.Println(<-quit)
	}
}

func main() {
	channel := make(chan string)
	quit := make(chan string)

	for i := 1; i <= 3; i++ {
		go Race(channel, quit, i)
	}

	for {
		select {
		case raceupdates := <-channel:
			fmt.Println(raceupdates)
		case winner := <-quit:
			fmt.Println(winner)
			quit <- "You Win"
			return
		}
	}
	time.Sleep(time.Second * 10)
}

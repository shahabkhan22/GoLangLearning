package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeBomb := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		timeBomb <- "Bomb Diffused"
	}()

	timeout := time.After(time.Duration(rand.Intn(500)) * time.Millisecond)
	for {
		select {
		case s := <-timeBomb:
			fmt.Println(s)
			return
		case <-timeout:
			fmt.Println("Bomb Exploded")
			return
		}
	}

}

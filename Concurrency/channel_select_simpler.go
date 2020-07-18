package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			chan1 <- "I will print every 100 ms"
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			chan1 <- "I will print every 1 s"
			time.Sleep(time.Second * 1)
		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case rec := <-chan1:
			fmt.Println(rec)
		case rec := <-chan2:
			fmt.Println(rec)
		default:
			fmt.Println("Default Case: No Channel is Ready")
		}
	}
}

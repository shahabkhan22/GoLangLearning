package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// .Add(int) - adds number of goroutines which waitgroup has to wait for
// .Done() - called within goroutine to signal it has executed
// .Wait() - blocks program until all goroutines specified in Add() have invoked Done()

func WelcomeMessage() {
	fmt.Println("HEllo WelcomeMessage")
}

func WelcomeMessage2() {
	fmt.Println("Hello from WelcomeMessage2")
	wg.Done()
}

func main() {
	wg.Add(3)

	go func() {
		WelcomeMessage()
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from MAin")
		wg.Done()
	}()

	go WelcomeMessage2()

	wg.Wait()
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("Right Now CPU : \t", runtime.NumCPU)
	fmt.Println("Right Now Routines: \t", runtime.NumCPU)
	go func() {
		fmt.Println("First newly created Go Routine")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second newly created Go Routine")
		wg.Done()
	}()

	fmt.Println("After launching 2 Go Routines")
	fmt.Println("Right Now CPU : \t", runtime.NumCPU)
	fmt.Println("Right Now Routines: \t", runtime.NumCPU)

	wg.Wait()

	fmt.Println("Exiting main")
	fmt.Println("Right Now CPU : \t", runtime.NumCPU)
	fmt.Println("Right Now Routines: \t", runtime.NumCPU)
}

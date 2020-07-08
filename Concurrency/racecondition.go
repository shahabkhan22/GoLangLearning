package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Number of CPU Cores :\t", runtime.NumCPU())
	fmt.Println("Number of GO Routines :\t", runtime.NumGoroutine())
	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second * 2)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Number of GO Routines :\t", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Counter Value: ", counter)
}

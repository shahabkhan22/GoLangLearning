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

	var mu sync.Mutex //using a mutex to prevent race condition

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second * 2)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Print("i = ", i)
		fmt.Println("\tNumber of GO Routines :\t", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Counter Value: ", counter)
}

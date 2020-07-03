package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

//stack trace example
func main() {
	//fmt.Println("OS :\t", runtime.GOOS)
	//fmt.Println("Architecture :\t", runtime.GOARCH)
	//fmt.Println("Lock OS Thread :\t", runtime.LockOSThread)
	fmt.Println("Number of CPU Cores :\t", runtime.NumCPU())
	fmt.Println("Number of GO Routines :\t", runtime.NumGoroutine())

	wg.Add(1)
	go func1()

	go func2()

	fmt.Println("Number of CPU Cores :\t", runtime.NumCPU())
	fmt.Println("Number of GO Routines :\t", runtime.NumGoroutine())
	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func1:", i)
	}
	wg.Done()

}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func2:", i)
	}

}

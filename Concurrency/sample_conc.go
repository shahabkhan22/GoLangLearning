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

/*
A type may have a method set associated with it.
Method set of an interface type is its interface.
The method set of any other type T consists of all the methods
declared with the receiver type T.

In a nustshell,
METHOD SETS determine WHAT METHODS attach to a TYPE.

Receivers		|       Values
---------------------------------
(t T)			|	T and *T
(t *T)			|	*T


IMPORTANT :-
The method set of a type determines the interfaces that the type
implements and the methods that can be called using a receiver of
that type
*/

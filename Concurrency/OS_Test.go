package main

import (
	"fmt"
	"runtime"
)

//stack trace example
func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	example(make([]string, 2, 4), "hello", 10)
}

func example(slice []string, str string, i int) {
	panic("Want Stack Trace")
}

/* OUTPUT of the PROGRAM

panic: Want Stack Trace

goroutine 1 [running]:
main.example(...)
        D:/Go/GoLangLearning/Concurrency/OS_Test.go:9
main.main()
        D:/Go/GoLangLearning/Concurrency/OS_Test.go:5 +0x40
exit status 2

*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOMAXPROCS : %d", runtime.GOMAXPROCS(3))
	fmt.Println("GOMAXPROCS : %d", runtime.GOMAXPROCS(0))
}

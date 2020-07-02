package main

import (
	"fmt"
	"runtime"
)

func main() {

	//NumCPU return the number of logical CPUs usage by current process
	fmt.Println(runtime.NumCPU())

}

package main

import (
	"fmt"
)

func main() {
	defer function1()
	function2()
	fmt.Println("In Function Main")

}

func function1() {
	fmt.Println("In Function 1")
}

func function2() {
	fmt.Println("In Function 2")
	defer function3()
}

func function3() {
	fmt.Println("In Function 3")
}

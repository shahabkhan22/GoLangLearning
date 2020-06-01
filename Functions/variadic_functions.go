package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9, 10} // unfurling slice
	function(xi...)
}

func function(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0 //short decleration operator
	// var sum int - this would do the same as above short decleration
	for v := range x {
		sum += v
	}
	fmt.Println("Sum = ", sum)
}

package main

import "fmt"

func main() {
	n := fact(10)
	fmt.Println(n)

}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

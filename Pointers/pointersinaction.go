package main

import "fmt"

//everything in Go is pass by value only

func func1(x int) {
	x = 22
}

func func2(x *int) {
	*x = 27
}

func main() {
	x := 0
	func1(x)
	fmt.Println("X Before : ", &x)
	fmt.Println("X Before : ", x)
	func2(&x)
	fmt.Println("X After : ", &x)
	fmt.Println("X After : ", x)
}

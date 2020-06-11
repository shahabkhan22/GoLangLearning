package main

import "fmt"

var (
	area = func(a int, b int) int {
		return a * b
	}
)

func main() {
	fmt.Println()

	func() {
		fmt.Println("Hello from Anonymous")
	}()

	func(i int) {
		fmt.Println("Hello from ", i)
	}(42)

	fmt.Println(area(20, 20))

	//func is a type in go like any other type
	//they can be
	f := func() {
		fmt.Println("from func expressions")
	}
	f()
}

package main

import "fmt"

var t int

type myType int

var x myType

func main() {
	fmt.Println("x = ", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println("x = ", x)
}

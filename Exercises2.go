package main

import "fmt"

var x int
var y string
var z bool

func main() {

	x = 20
	y = "This is a string"
	z = true
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}

package main

import "fmt"

/*
var y = 42
var z = "This is a string"
var x float32 = 3123
var p string = `I said




"this is a string"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(p)
	fmt.Printf("%T\n", p)
}
*/
// Creating own types and type conversion(NOT SAME AS TYPE CASTING)
var a int

type hotdog int

var b hotdog
var c hotdog = 44

func main() {
	a = 42
	b = 43
	//c = 44
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}

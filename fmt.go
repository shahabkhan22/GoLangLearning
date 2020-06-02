package main

import "fmt"

var x = 10

//var y = true

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)
	fmt.Printf("%#v\n", x)
	fmt.Printf("%%\n", x)
	s := fmt.Sprintf("%#x\t%b\t%x", x, x, x)
	fmt.Println(s)

}

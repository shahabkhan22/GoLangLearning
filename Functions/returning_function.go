package main

import "fmt"

func myfunction() string {
	return "hello"
}

//syntax is just like any other function, just that return tyoe is of a function type
//func <function name>() <return type> {code}
func myfunction2() func() int {
	return func() int {
		return 22
	}
}

func main() {
	s := myfunction()
	fmt.Println(s)
	x := myfunction2()
	fmt.Println(x)
	fmt.Printf("%T", x) //this statement prints "func() int" as that is type of this x

}

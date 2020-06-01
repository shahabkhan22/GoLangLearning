package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello from Main")
	foo()
	vara("vara", 10)
	ret1 := ret(1234)
	fmt.Println(ret1)
	x, y := ret2("abc", "xyz")
	fmt.Println(x, " ", y)

}

func foo() {
	fmt.Println("Hello from foo")
}

func vara(s string, i int) {
	fmt.Println("Hello ", s, i)
}

func ret(i int) string {
	return fmt.Sprint("Hello from ret", i)
}

func ret2(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln)
	b := false
	return a, b
}

package main

import "fmt"

func main() {
	fmt.Println("Hello Playground")
	foo()
	bat("main")
}

func foo() {
	fmt.Println("Hello from called function")
}

func bat(s string) {
	fmt.Println("Hello ", s)
}

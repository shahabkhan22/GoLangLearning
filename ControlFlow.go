package main

import "fmt"

//For loop in GoLang unifies For and While loop found in many languages, and GoLang has no Do-While Loop.
//Syntax - for init; condition; post { }
//GoLang has a break and continue statements
//Doing a while loop in Go
func main() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	//Using for without any parameters
	x = 0
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++

	}
}

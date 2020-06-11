package main

import "fmt"

//callback is a concept of functional programming
//functional programming is not recommended in go
//In go callback are implemented by sending a function as an argument
func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := sum(i...)
	fmt.Println("Sum Total: ", s)

	e := even(sum, i...)
	fmt.Println("Sum Evan: ", e)

	o := odd(sum, i...)
	fmt.Println("Sum Odd: ", o)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

//taking a function as an argument
func even(f func(x ...int) int, vi ...int) int {
	var s []int
	for _, v := range vi {
		if v%2 == 0 {
			s = append(s, v)
		}
	}
	return f(s...)
}

//similarly sum of odd numbers
func odd(f func(x ...int) int, vi ...int) int {
	var s []int
	for _, v := range vi {
		if v%2 != 0 {
			s = append(s, v)
		}
	}
	return f(s...)
}

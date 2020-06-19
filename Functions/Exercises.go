package main

import (
	"fmt"
	"math"
)

func defer1() {
	fmt.Println("In Defer1")
}

func defer2() {
	fmt.Println("In Defer2")
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

type person struct {
	fname string
	lname string
	age   int
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func (p person) speak() {
	fmt.Println(p.fname, " ", p.lname, " ", p.age, " years")
}

func func1() func() int {
	return func() int {
		return 27
	}
}

//this is to understand "callback" for functions
func func2(f func(x []int) int, xi []int) int {
	n := f(xi)
	n++
	return n
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t1 := foo(x...)
	fmt.Println("Total from Foo = ", t1)

	y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t2 := bar(y)
	fmt.Println("Total from Bar = ", t2)

	defer defer1()
	defer2()
	fmt.Println("Exiting main()")

	p1 := person{
		fname: "ABC",
		lname: "XYZ",
		age:   26,
	}
	p1.speak()

	circ := circle{
		radius: 27.4,
	}

	squa := square{
		side: 22.5,
	}
	info(circ)
	info(squa)

	func() { //anonymous function
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}() //these are used to create a go sub-routine which can run something on itself

	//assigning a function to a variable
	f := func() {
		fmt.Println("This function is assigned to a variable")
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}
	f() // calling that function assigned to a varibale
	fmt.Printf("type of function assigned to variable = %T\n", f)
	fmt.Println() //adding for line change to create a cleaner output

	//calling a function that return another functiomn
	m := func1()     //assigning the returned function to a variable
	fmt.Println(m()) //printing the value returned

	g := func(x []int) int {
		if len(x) == 0 {
			return 0
		}
		if len(x) == 1 {
			return x[0]
		}
		return x[0] + x[len(x)-1]

	}
	val := func2(g, []int{1, 2, 3, 4, 5})
	fmt.Println(val)
}

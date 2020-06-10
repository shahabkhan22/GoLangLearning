package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type users struct {
	person
	exists bool
}

func (u users) method() {
	fmt.Println("I am", u.fname, u.lname)
}

func main() {
	p1 := users{
		person: person{
			"ABC",
			"XYZ",
		},
		exists: true,
	}

	p2 := users{
		person: person{
			"PQR",
			"XYZ",
		},
		exists: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	p1.method()
	p2.method()

}

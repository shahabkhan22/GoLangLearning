package main

import "fmt"

type person struct {
	fname string
	lname string
}

type users struct {
	person
	exists bool
}

func (p person) Speak() {
	fmt.Println(p.fname, p.lname, "- Person")
}

type human interface {
	Speak()
}

func humantest(h human) {
	fmt.Println("Human", h)
}

//a value can be of more than one type
//p1 value is of type users but it also has Speak method attached to it,
//it is also type human

func main() {
	p1 := users{
		person: person{
			"ABC",
			"XYZ",
		},
		exists: true,
	}

	p2 := person{
		fname: "PQR",
		lname: "XYZ",
	}

	fmt.Println(p1)
	p1.Speak()
	p2.Speak()
	fmt.Println(p2)
}

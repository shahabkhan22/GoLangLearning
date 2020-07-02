package main

import "fmt"

type Employee struct {
	//person
	fname  string
	status bool
}

func main() {

	p1 := struct {
		fname string
		lname string
		age   int
	}{
		fname: "ABC",
		lname: "XYZ",
		age:   10,
	}

	p2 := person{
		fname: "PQR",
		lname: "STU",
		age:   20,
	}
	emp1 := Employee{
		person: person{
			fname: "shahab",
			lname: "khan",
			age:   25,
		},
		fname:  "mohd shahab",
		status: true,
	}

	//fmt.Println(p1)
	//fmt.Println(p2)

	fmt.Println(p1.fname, p1.lname, p1.age)
	fmt.Println(p2.fname, p2.lname, p2.age)
	fmt.Println(emp1.fname, emp1.person.fname, emp1.lname, emp1.age, emp1.status)

	fmt.Println(p1)
}

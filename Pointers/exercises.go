package main

import "fmt"

type People struct {
	fname string
	lname string
	age   int
}

func changePeople(p *People) {
	fmt.Printf("%T\n", *p)
	(*p).fname = "HAHA"
	p.lname = "JOKE"
	(*p).age = 79
}

func main() {
	a := 27
	fmt.Println(a)
	fmt.Println(&a)

	p1 := People{
		fname: "ABC",
		lname: "XYZ",
		age:   26,
	}

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", &p1)
	fmt.Println(p1)
	changePeople(&p1)
	fmt.Println(p1)
}

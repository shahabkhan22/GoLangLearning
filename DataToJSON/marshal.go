package main

import (
	"encoding/json"
	"fmt"
)

/*
Marshal Function- converts a GoLang struct into JSON object using json.Marshal
Its GoLang's way of saying "encode/convert to JSON Object" because GoLang is a strictly
typed language and JSON is a dynamically typed language.

*/
type People struct {
	Fname string
	Lname string
	Age   int
}

func main() {
	p1 := People{
		Fname: "ABC",
		Lname: "XYZ",
		Age:   27,
	}

	p2 := People{
		Fname: "PQR",
		Lname: "XYZ",
		Age:   26,
	}

	PAll := []People{
		p1,
		p2,
	}
	fmt.Println("This is a struct and its type")
	fmt.Printf("%T\n", PAll)
	fmt.Println(PAll)

	bs, err := json.Marshal(PAll)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("This is a JSON of our above Struct and its type")
	fmt.Printf("%T\n", bs) // this will give us uint8
	fmt.Println(string(bs))
	fmt.Println(bs)

}

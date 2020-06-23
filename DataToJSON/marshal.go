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
	fname string
	lname string
	age   int
}

func main() {
	p1 := People{
		fname: "ABC",
		lname: "XYZ",
		age:   27,
	}

	p2 := People{
		fname: "PQR",
		lname: "XYZ",
		age:   26,
	}

	PAll := []People{
		p1,
		p2,
	}

	fmt.Println(PAll)

	bs, err := json.Marshal(PAll)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	fmt.Println(bs)

}

package main

import (
	"encoding/json"
	"fmt"
)

//if you need a data structure in Go for a JSON object, go tho
//https://mholt.github.io/json-to-go/
//enter your JSON and it gives you an appropriate Go data structure

type People struct {
	Fname string `json:Fname`
	Lname string `json:Lname`
	Age   int    `json:Age`
}

func main() {
	s := `[{"Fname":"ABC","Lname":"PQR","Age":27},{"Fname":"XYZ","Lname":"PQR","Age":22}]`
	bs := []byte(s) //byte slice
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var person []People

	err := json.Unmarshal(bs, &person)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nAll of the Data :", person)

	for i, v := range person {
		fmt.Println("\nPERSON NUMBER: ", i)
		fmt.Println(v.Fname, v.Lname, v.Age)
	}
}

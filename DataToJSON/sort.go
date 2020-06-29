package main

import (
	"fmt"
	"sort"
)

//sort package comes with all the type's sorting methods implemented as well as built in checks

type People struct {
	Fname string
	Age   int
}

//ByAge implements sort.Interface for the []People based on Age Field
//If sort has these 3 methods implemented, then it can go through these methods
//and it will perform the sorting for your used defined type based on the fields
//of your user defined type
//You first have to convert and pass the type you are implementing to the already
//type and pass the slice you want to sort
type ByAge []People

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []People

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Fname < a[j].Fname }

func main() {
	/*
		i := []int{2, 3245, 456, 2345, 4, 45, 68, 889, 798787}
		fmt.Println(i)
		sort.Ints(i)
		fmt.Println(i)

		s := []string{"ABC", "PQR", "XYZ", "ZZZ", "A", "BC", "M"}
		fmt.Println(s)
		sort.Strings(s)
		fmt.Println(s)

		if sort.IntsAreSorted(i) == true {
			fmt.Println("Ints are Sorted")
			if sort.StringsAreSorted(s) == true {
				fmt.Println("Strings are Sorted")
			}

		}
	*/
	p1 := People{"ABC", 22}
	p2 := People{"PQR", 27}
	p3 := People{"XYZ", 4}

	Persons := []People{p3, p2, p1}

	fmt.Println(Persons)

	//sorting all these People w.r.t their Age/Name
	sort.Sort(ByAge(Persons))
	fmt.Println(Persons)

	sort.Sort(ByName(Persons))
	fmt.Println(Persons)
}

package main

import "fmt"

//basically stores an address

func main() {
	a := 27
	fmt.Println(a, &a)     //& is an operator that gives the address in memory where value of a is stored
	fmt.Printf("%T\n", &a) // this gives the type of the value that &a stores, i.e. *int
	fmt.Println()
	//cannot assing the pointer type to a value type
	// b := &a doesn't work

	b := &a                //storing the address of a in b
	fmt.Println(b)         //this prints the address of a because b stores the adress
	fmt.Printf("%T\n", &b) //this returns type of b i.e. **int  because it is a pointer that holds a pointer
	fmt.Println(*b)        //this * is dereferencing the address stored in b, i.e. the value that a holds as the address b has is of a
	fmt.Println(*&a)       //this lines is same as above
}

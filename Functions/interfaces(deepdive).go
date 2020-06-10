package main

import "fmt"

/*Interfaces in GO is two things -
1. Set of Methods
2. A type

Go's type system has a core concept "Design abstractions in terms of"-
	What actions our types can execute
			RATHER THAN
	What kind of data our types can hold

For example- Animal is defines as anything that can Speak

*Whether or not a type satisfies an interface is determined automatically
*/

type Animals interface {
	Speak() string //Any type that defines this method is said to satisfy Animal interface
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "woof woof"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow Meow"
}

/*interface{} type

*This has no methods, hence a bit confusing concept to grasp
 As it has zero method to implement then any method that takes
 interface{} values as a parameter can be supplied with any value.
 So the below funtion will accept any parameter

	func anyfunction(v interface{}){
	}

Now, what is type of 'v' ?

v is not of any type, it is of interface{} type
Conversion is required of the desired type to an iterface{} type
*/

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	fmt.Println("Interfaces : Set of Methods implementation")
	animals := []Animals{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	//calling interface type
	/*
		names := []string{"ABC", "PQR", "XYZ"}
		PrintAll(names)

		Above code will generate error as:
		cannot use names (type []string) as type []interface in funtion argument

		Hence it requires type conversion as follows-
	*/
	names := []string{"ABC", "PQR", "XYZ"}
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)

}

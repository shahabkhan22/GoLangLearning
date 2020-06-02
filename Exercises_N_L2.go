package main

import (
	"fmt"
)

/* E1
func main() {
	a := 22
	fmt.Printf("%d\t\t%b\t\t%#x", a, a, a)
	fmt.Println()

}
*/

/* E2
func main() {
	a := (42 == 42)
	b := 12 <= 24
	c := 13 >= 22
	d := 12 > 22
	e := 12 < 12
	f := 12 != 21
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
*/

/* E3
const (
	a     = 22
	b int = 27
)

func main() {
	fmt.Println(a, b)

}
*/
/* E4

var a int = 42

func main() {
	fmt.Printf("%d\t%b\t%#x", a, a, a)
	b := a << 1
	fmt.Println()
	fmt.Printf("%d\t%b\t%#x", b, b, b)
	fmt.Println()

}
*/
/* E5
func main() {
	s := `"This is a string
	created using a string literal
	which allows mulitple lines into the 				string
	as well as tabs.



	Also skipped lines in the code"`

	fmt.Println(s)
}
*/
/* E6 */
const (
	a = iota + 2016
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	var a int = 42
	fmt.Printf("%#x\t", a)
}

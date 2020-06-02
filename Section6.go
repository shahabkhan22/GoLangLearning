package main

import "fmt"

/*//Bool Type
const a = 42
const b = 3.14
const c = "This is a string"
func main() {
	//fmt.Println(runtime.GOOS)
	//fmt.Println(runtime.GOARCH)
	s := "This is a string"
	fmt.Println(s)
	s = "hello worodl"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}
}

//IOTA
const (
	a = iota + 2
	b
	c
	d
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
*/

//BIT SHIFTING

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}

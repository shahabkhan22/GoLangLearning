package main

import "fmt"

const s string = "constant string"

func main() {
	/*
		fmt.Println("hello world")

		var i = "string variable"
		//int a,b,c = 10,230,40
		var x, y int = 21, 22
		var z = true
		var e int
		f := "apple"

		fmt.Println(i)
		//fmt.Println(a, b, c)
		fmt.Println(x, y)
		fmt.Println(z)
		fmt.Println(e)
		fmt.Println(f)

		fmt.Println("-----------Constants-----------")
		fmt.Println(s)
		const n = 500000000000
		const d = 3e20 / n
		fmt.Println(d)
		fmt.Println(int64(d))
		fmt.Println(math.Sin(n))
	*/

	/*
		fmt.Println("------Variables------")
		f := 1
		a := 3.14
		b := "a"
		c := "this is a string"
		var x float32
		var p float64
		var q string
		var y, z = 21, 22
		fmt.Println(f, a, b, c, y, z)
		fmt.Println(x, p)
		fmt.Println("Q = ", q)
	*/

	fmt.Println("-------For Loop----------------")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}

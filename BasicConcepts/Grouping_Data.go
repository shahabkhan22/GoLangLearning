package main

import "fmt"

func main() {
	//var x [5]int

	x := []int{4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println("------------------------------------------")
	//fmt.Println(cap(x))
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("------------------------------------------")
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[0:2])
	fmt.Println("------------------------------------------")
	x = append(x, 10, 12, 13, 14, 15)
	fmt.Println(x)
	y := []int{20, 30, 40, 50, 999}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("------------------------------------------")
	y = append(x[:2], x[4:]...)
	fmt.Println(x)
	fmt.Println("------------MAKE COMMAND------------------------------")

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 11)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//x = append(x, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x := make([]int, 50, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	y := new([100]int)[0:50]
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	fmt.Println("------------Slice of Slices------------------------------")
	s1 := []string{"This", "is", "for", "slice", "one"}
	s2 := []string{"This", "is", "for", "slice", "two"}

	mslice := [][]string{s1, s2}

	fmt.Println(mslice)

	fmt.Println("------------MAP------------------------------")
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m1)

	fmt.Println(m1["a"])

	fmt.Println(m1["d"])

	if value, ok := m1["c"]; ok {
		fmt.Println(value)
	}

	m1["d"] = 4
	for k, v := range m1 {
		fmt.Println(k, v)

	}

	delete(m1, "a")
	for k, v := range m1 {
		fmt.Println(k, v)

	}
}

package main

import "fmt"

func main() {
	/*
							x := [5]int{1, 2, 3, 4, 5}
							for i, v := range x {
								fmt.Println(i, v)
							}
							fmt.Printf("%T\n", x)

							y := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
							fmt.Println(y[:5])
							fmt.Println(y[5:])
							fmt.Println(y[3:8])
							fmt.Println(y[4:9])
							fmt.Println(y[1:6])


						x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50}
						x = append(x, 52)
						for i, v := range x {
							fmt.Println(i, v)
						}
						x = append(x, 53, 54, 55)
						for i, v := range x {
							fmt.Println(i, v)
						}

						y := []int{56, 57, 58, 59, 60}
						x = append(x, y...)
						for i, v := range x {
							fmt.Println(i, v)
						}

					x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
					x = append(x[:3], x[6:]...) //[42, 43, 44, 48, 49, 50]
					fmt.Println(x)

				x := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
				fmt.Println(len(x))
				fmt.Println(cap(x))

				for i := 0; i < 50; i++ {
					x[i] = fmt.Sprintf("Position no %d : ", i)
				}

				fmt.Println(x)
				fmt.Println(len(x))
				fmt.Println(cap(x))

				x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
				//fmt.Println(x)
				//fmt.Println(len(x))
				//fmt.Println(cap(x))

				for i := 0; i < len(x); i++ {
					fmt.Println(i, x[i])
				}

		s1 := []string{"Shahab", "Khan", "Arsenal FC"}
		s2 := []string{"Khan", "Mohd", "FC Barcelona"}
		fmt.Println(s1)
		fmt.Println(s2)
		ss := [][]string{s1, s2}
		fmt.Println(ss)

		for i, s := range ss {
			fmt.Println("record", i)
			for j, v := range s {
				fmt.Printf("\tindex position:%v\t value : %v\t\n", j, v)
			}

		}
	*/

	m := map[string][]string{
		"a_b": []string{"football", "Arsenal FC", "FC Barcelona"},
		"c_d": []string{"volleyball", "Brazil", "USA"},
		"e_f": []string{"basketball", "LA Lakers", "Golden State Warriors"},
	}

	fmt.Println(m["a_b"])
	fmt.Println(m["c_d"])
	fmt.Println(m["e_f"])
	/*
		for k, v := range m {
			fmt.Println("Record of : ", k)
			for i, v2 := range v {
				fmt.Println("\t", i, v2)
			}
		}
	*/
	m["g_h"] = []string{"Tennis", "Roger Federer", "Rafael Nadal"}
	/*	for k, v := range m {
			fmt.Println("Record of : ", k)
			for i, v2 := range v {
				fmt.Println("\t", i, v2)
			}
		}
	*/
	delete(m, "a_b")
	for k, v := range m {
		fmt.Println("Record of : ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

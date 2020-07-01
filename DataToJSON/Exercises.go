package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type User struct {
	Fname string
	Age   int
}

type People struct {
	Fname string `json:Fname`
	Lname string `json:Lname`
	Age   int    `json:Age`
}

func main() {
	user1 := User{
		Fname: "ABC",
		Age:   26,
	}
	user2 := User{
		Fname: "MAN",
		Age:   22,
	}
	user3 := User{
		Fname: "SHA",
		Age:   26,
	}

	Users := []User{user1, user2, user3}

	fmt.Println(Users)

	bs, err := json.Marshal(Users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	fmt.Println(bs)

	myjson := `[{"Fname":"ABC","Lname":"PQR","Age":27},{"Fname":"XYZ","Lname":"PQR","Age":22}]`
	var people []People

	err2 := json.Unmarshal([]byte(myjson), &people)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("All People: ", people)

	for i, v := range people {
		fmt.Println("People  :", i)
		fmt.Println(v.Fname, v.Lname, v.Age)
	}

	fmt.Println(Users)
	err3 := json.NewEncoder(os.Stdout).Encode(Users)
	if err3 != nil {
		fmt.Println(err)
	}

	i := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	s := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(i)
	sort.Ints(i)
	fmt.Println(i)

	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//Brcypt is a password hasing function
//it is in the x library - s stands for experimental
//this means it has not yet been included in a standard library

//this code takes a password and converts it into a hash and stores it
//then at the time of login, it comapres the hash with the hash of
//password entered by the user
//system never stores the actual password, just their hashes

func main() {
	s := "password"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPass := "password"
	err1 := bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err1 != nil {
		fmt.Println("Password Wrong")
		return
	}
	fmt.Println("Login Successful")

}

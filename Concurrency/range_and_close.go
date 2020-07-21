package main

import "fmt"

type Money struct {
	amount int
	year   int
}

func sendMoney(parent chan Money) {
	for i := 1; i <= 18; i++ {
		parent <- Money{5000, i}
	}
	close(parent)
}

func main() {
	money := make(chan Money)

	go sendMoney(money)

	for childMoney := range money {
		fmt.Printf("Money receive : %d \t Year : %d\n", childMoney.amount, childMoney.year)
	}
}

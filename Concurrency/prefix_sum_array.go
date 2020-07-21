package main

import "fmt"

func PrefixSum(myarray, myoutput []int, parent chan int) {
	if len(myarray) < 2 {
		parent <- myarray[0]
		myoutput[0] = myarray[0] + <-parent

	} else if len(myarray) < 1 {
		parent <- 0
		<-parent

	} else {

		mid := len(myarray) / 2
		left := make(chan int)
		right := make(chan int)
		go PrefixSum(myarray[:mid], myoutput[:mid], left)
		go PrefixSum(myarray[mid:], myoutput[mid:], right)

		leftsum := <-left
		parent <- leftsum + <-right
		fromleft := <-parent

		left <- fromleft
		right <- fromleft + leftsum

		<-left
		<-right
	}

	parent <- 0
}

func main() {

	myarray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := make([]int, len(myarray))
	parent := make(chan int)

	go PrefixSum(myarray, output, parent)

	sum := <-parent

	fromleft := 0
	parent <- fromleft

	donezero := <-parent
	fmt.Println(myarray, output, sum, donezero)
}

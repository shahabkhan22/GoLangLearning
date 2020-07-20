package main

import "fmt"

func fibonacci(n int) chan int {
	mychan := make(chan int)

	go func() {
		k := 0
		for i, j := 0, 1; k < n; k++ {
			mychan <- i
			i, j = i+j, i

		}
		close(mychan)
	}()

	return mychan
}

func main() {

	for i := range fibonacci(10) {
		fmt.Println(i)
	}
}

//as the channel is returned from fibonacci() function, it allows the function to
//communicate each and every term to the main routine as soon as it is comouted
//in the function body.
//On printing the n-th term in sequence, we can do anything with it as the /
// next term is being computed

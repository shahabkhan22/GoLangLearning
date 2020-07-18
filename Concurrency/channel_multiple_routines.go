package main

import "fmt"

func square(myIntChannel chan int) {
	fmt.Println("[Inside Square]")
	num := <-myIntChannel
	myIntChannel <- num * num
}

func cube(myIntChannel chan int) {
	fmt.Println("[Inside Cube]")
	num := <-myIntChannel
	myIntChannel <- num * num * num
}

func main() {
	squareChannel := make(chan int)
	cubeChannel := make(chan int)

	go square(squareChannel)
	go cube(cubeChannel)

	num := 3

	fmt.Println("[Inside Main] - sending number to square")
	squareChannel <- num
	fmt.Println("[Inside Main] - Resuming Main")

	fmt.Println("[Inside Main] - sending number to cube")
	cubeChannel <- num
	fmt.Println("[Inside Main] - Resuming Main")

	fmt.Println("[Inside Main] - Reading from Channels")

	squareval, cubeval := <-squareChannel, <-cubeChannel
	sum := squareval + cubeval

	fmt.Println("[Inside Main] - Sum of Square and CUbe of ", num, ": ", sum)

}

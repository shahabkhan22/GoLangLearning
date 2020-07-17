package main

import "fmt"

func sendValues(myStringChannel chan string) {
	for i := 0; i < 5; i++ {
		myStringChannel <- string(i + 65) //sending value

	}
	close(myStringChannel)
}

func main() {
	myStringChannel := make(chan string)
	go sendValues(myStringChannel)
	/*
		for i := 0; i < 10; i++ {
			value, open := <-myStringChannel
			//this checks if the channel is still open so receiver doesn't stays in receive mode while channel has been closed
			if !open {

				break
			}
			fmt.Println(value)
		}
	*/
	/* another way to check channel status on receiver end

	 */
	for value := range myStringChannel {
		fmt.Println(value)
	}

}

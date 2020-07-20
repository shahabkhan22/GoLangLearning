package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var mynums [10]int

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		mynums[i] = rand.Intn(50)
	}

	mychanout := chanGenerator(mynums) //common channel

	//fan-out common channel mychanout to two goroutines, which are created inside double
	//two instances of double() as two goroutines are concurrently receiving data from a single channel
	//this is basically fan-out where data feom one channel is distributed to two different channels concurrently
	//which divides the computation of doubling the integers and returns them as strins among two goroutines
	mychan1 := double(mychanout)
	mychan2 := double(mychanout)

	mychanIn := fanIn(mychan1, mychan2) //fan-in results from the channels returned by the double function

	for i := 0; i < len(mynums); i++ {
		fmt.Println(<-mychanIn)
	}
}

//this function returns a channel which will receive data from goroutines create in the function itself
//converting an integer value from array to string and sending on the channel
func chanGenerator(nums [10]int) <-chan string {
	channel := make(chan string)

	go func() {
		for _, i := range nums {
			channel <- strconv.Itoa(i)
		}
		close(channel)
	}()
	return channel
}

//takes common channel as input and sends data on to another channel
//
func double(inputchannel <-chan string) <-chan string {
	channel := make(chan string)

	go func() {
		for i := range inputchannel {
			num, err := strconv.Atoi(i)
			if err != nil {

			}
			channel <- fmt.Sprintf("%d * 2 = %d", num, num*2)
		}
		close(channel)
	}()
	return channel
}

func fanIn(inputchannel1, inputchannel2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			select {
			case message1 := <-inputchannel1:
				channel <- message1
			case message2 := <-inputchannel2:
				channel <- message2
			}
		}
	}()
	return channel
}

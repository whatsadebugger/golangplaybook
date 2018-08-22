package main

import "fmt"

func main() {
	channelForInts := make(chan int)
	go printIntFromChannel(channelForInts)
	channelForInts <- 5

}

// someimes this method will have enough time to print a value.
func printIntFromChannel(channel chan int) {
	number := <-channel
	fmt.Println("The number is:", number)
}

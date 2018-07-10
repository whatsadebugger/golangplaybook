package main

import "fmt"

func main() {
	// https://blog.mergermarket.it/now-youre-thinking-with-channels/ amazing read
	intChan := make(chan int) // unbuffered channel
	go sendOnThisChannel(intChan) // if you take out the go here there will be a deadlock
	number := <-intChan
	fmt.Println("Got int ", number)
}

func sendOnThisChannel(c chan int) {
	c <- 5
}

package main

import "fmt"

func main() {
	// https://blog.mergermarket.it/now-youre-thinking-with-channels/ amazing read
	intChan := make(chan int) // unbuffered channel blocks causes block
	go sendOnThisChannel(intChan)
	number := <-intChan
	fmt.Println("Got int ", number)
}

func sendOnThisChannel(c chan int) {
	c <- 5
}

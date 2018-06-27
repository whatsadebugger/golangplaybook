package main

import "fmt"

func main() {
	intChan := make(chan int) // unbuffered channel blocks causes block
	go sendOnThisChannel(intChan)
	number := <-intChan
	fmt.Println("Got int ", number)
}

func sendOnThisChannel(c chan int) {
	c <- 5
}

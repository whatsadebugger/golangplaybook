package main

import "fmt"

// buffered channels dont block
func main() {
    channelForInts := make(chan int, 3)
    channelForInts <- 5
    number := <- channelForInts
    fmt.Println("The number is", number)
}
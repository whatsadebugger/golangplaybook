package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://pragmacoders.com/multithreading-go-tutorial/ example taken from here. Article states that the code would
// not work with a unbuffered channel however that is not true. The first routine to respond will win the race and
// the other one will be stuck waiting for someone to recieve its value. It will just get removed because the program ended.
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	query := "Our Query"
	results := make(chan string)

	go googleIt(results, query)
	go bingIt(results, query)

	resp := <-results

	fmt.Printf("Sent query:\t\t %s\n", query)
	fmt.Printf("Got Response:\t\t %s\n", resp)
}

func googleIt(results chan<- string, query string) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	results <- "Google Response"
}

func bingIt(results chan<- string, query string) {
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)

	results <- "Bing Response"
}

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
  respond := make(chan string)

  go googleIt(respond, query)
  go bingIt(respond, query)

  queryResp := <-respond

  fmt.Printf("Sent query:\t\t %s\n", query)
  fmt.Printf("Got Response:\t\t %s\n", queryResp)
}

func googleIt(respond chan<- string, query string) {
  time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

  respond <- "Google Response"
}

func bingIt(respond chan<- string, query string) {
  time.Sleep(time.Duration(rand.Intn(2)) * time.Second)

  respond <- "Bing Response"
}
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()
	<-time.After(time.Microsecond * 3) // give the goroutine enoough time to occasionally win
	return i
}

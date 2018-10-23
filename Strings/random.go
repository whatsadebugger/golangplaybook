package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

func main() {
	fmt.Println(RandString(10))
	fmt.Println(RandString(11))
	fmt.Println(RandString(12))
	fmt.Println(RandString(13))
	fmt.Println(source.Int63() % 1000000000000000)
	fmt.Println(source.Int63() % 1000000000000000)
	fmt.Println(source.Int63() % 1000000000000000)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

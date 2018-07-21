package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 10)
	for i := range slice {
		slice[i] = i
	}
	fmt.Println(slice)

	for i := range slice[:3] {
		fmt.Println(i)
	}
}

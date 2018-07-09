package main

import (
	"fmt"
	"math"
)

func main() {
	// overflow wraps silently in go
	max := math.MaxInt64
	fmt.Println(max)
	fmt.Println(max + 1)
	fmt.Println(max + math.MaxInt64)
}

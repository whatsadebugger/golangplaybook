package main

import (
	"fmt"
)

func main() {
	spreadCreamCheese("2", "3", "4")

}

func spreadCreamCheese(args ...interface{}) {
	fmt.Println(args...)
	// Prints 2 3 4
}

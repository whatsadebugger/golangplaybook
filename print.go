package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("wow an error help")
	fmt.Printf("A bad error %v", err)
	// this example shows that error can be printed easily in printf using %v
	// Output: A bad error wow an error help
}

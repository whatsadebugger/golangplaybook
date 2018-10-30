// go run short-circuit.go
package main

import "fmt"

func val1() bool {
	fmt.Println("val1 gets called")
	return true
}

func val2() bool {
	fmt.Println("val2 gets called")
	return false
}

func main() {
	if false && val1() {
		fmt.Println("Shouldn't print")
	}
	if true || val1() {
		fmt.Println("Should print")
	}
}

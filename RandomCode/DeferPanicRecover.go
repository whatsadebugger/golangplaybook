package main

import (
	"fmt"
)

func main() {
	recoverer()
	fmt.Println("returned after panic")
}

func recoverer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in recoverer", r)
		}
	}()

	fmt.Println("Calling panicker.")
	panicker(0)
	fmt.Println("Returned normally from panicker.") // wont print
}

func panicker(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in panicker", i)
	fmt.Println("Printing in g", i)
	panicker(i + 1)
}

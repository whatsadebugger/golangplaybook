package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	delete := 1

	fmt.Println(numbers[5:])
	copy(numbers[delete:], numbers[delete+1:])
	numbers[len(numbers)-1] = 0
	numbers = numbers[:len(numbers)-1]
	fmt.Println("After deleting index", delete, numbers)

}

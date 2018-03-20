package main

import "fmt"

func main() {

	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	if attended["Ann"] { // will be false if person is not in the map
		fmt.Println("Ann", "was at the meeting")
	}
	values := []int{1, 4, 5, 6, 7, 8, 1, 6, 2, 6, 7, 1, 11111}
	fmt.Println(min(values...))

	fmt.Println(^uint(0))
	fmt.Println(^uint(0) >> 1)
}

func min(a ...int) int {
	min := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

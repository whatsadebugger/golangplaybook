package main

import (
	"fmt"
	"reflect"
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

	// divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2})

	numbers := "123456789"
	fmt.Println("type of an index in a string", reflect.TypeOf(numbers[0]))

	print("numbers: ")
	for _, v := range numbers {
		print(v, " ")
	}
	println()

	runes := []rune(numbers)
	fmt.Println("runes intitial", string(runes))
	size := len(runes)

	for i, j := 0, size-1; i < size/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	fmt.Printf("runes rev: %+v\n", runes)
	fmt.Println("string(runes):", string(runes))
	fmt.Println("string(numbers):", string(numbers))
}

// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	sumPairs := int32(0)
	for i := range ar {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				sumPairs++
			}
		}
	}
	return sumPairs
}

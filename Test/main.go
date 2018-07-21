package main

import "fmt"

func main() {
	divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2})
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
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

func migratoryBirds(ar []int32) int32 {
	count := make(map[int32]int32)

	for _, v := range ar {
		count[v]++
	}

	maxKey := int32(-1)
	max := int32(-1)

	for k, v := range count {
		if v > max {
			maxKey = k
			max = v
		}
	}
	return maxKey
}

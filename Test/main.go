package main

func main() {
	divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2})
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

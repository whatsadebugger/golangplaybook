package main

import "fmt"

func main() {
	divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2})
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
	fmt.Println(solve(2000))

	input := []int{3, 2, 4}
	// input = []int{2, 4, 11, 3}
	fmt.Println("Division", 0/10)
	fmt.Println(twoSum(input, 6))

	l1 := &ListNode{
		Val: 2, Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Println(addTwoNumbers(l1, l2))

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(longestDrome("babad"))
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
		if v > max || (k < maxKey && v >= max) {
			maxKey = k
			max = v
		}
	}
	return maxKey
}

// Complete the solve function below.
func solve(year int32) string {
	y := int(year)
	progDay := ""
	if y == 1918 {
		progDay = fmt.Sprintf("%02d.%02d.%d", 26, 9, y)
	} else if (y < 1918 && y%4 == 0) || (y > 1918 && ((y%400 == 0) || (y%4 == 0 && y%100 != 0))) {
		progDay = fmt.Sprintf("%02d.%02d.%d", 12, 9, y)
	} else {
		progDay = fmt.Sprintf("%02d.%02d.%d", 13, 9, y)
	}

	return progDay
}

func twoSum(nums []int, target int) []int {
	visited := make(map[int]int)
	sum := make([]int, 2)

	// solve for y
	for i, v := range nums {

		y := target - v

		// if we have a value in our map that adds to target we can return. Otherwise keep looking.
		if j, ok := visited[y]; ok {
			// add the values to our slice
			sum[0] = i
			sum[1] = j
			break
		} else {
			visited[v] = i
		}
	}
	return sum
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for l1 != nil || l2 != nil {
		var operand1 int
		var operand2 int

		if l1 != nil {
			fmt.Println(l1.Val)
			operand1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			fmt.Println(l2.Val)
			operand2 = l2.Val
			l2 = l2.Next
		}

		// add all the existing values
		current.Val += operand1 + operand2

		// find out how much carries over
		carry := current.Val / 10

		// mod 10 removes the carry and leaves the remainder
		current.Val = current.Val % 10

		// we only need a new next node if there is a carry or l1/l2 still has more nodes
		// to go through
		if l1 != nil || l2 != nil || carry != 0 {
			current.Next = &ListNode{}
		}

		// if our sum requires a carry add it to the new node
		if carry != 0 {
			current.Next.Val = carry
		}

		current = current.Next
	}

	return result
}

func lengthOfLongestSubstring(s string) int {
	max := 0

	// we need the index so we can start the search over
	// from that location
	seen := make(map[byte]int)

	// use this to keep track of our substring
	// when a duplicate value is found we will advance this forward
	left := 0

	for i := 0; i < len(s); i++ {
		// have we seen it?
		// slide the window over until we find the duplicate
		// delete the keys until the duplicate is removed
		if idx, ok := seen[s[i]]; ok {

			for ; left <= idx; left++ {
				// remove the value no longer in our window
				delete(seen, s[left])
			}
		}

		// mark current value as seen
		seen[s[i]] = i

		// check if our sequence exceeded that max
		length := (i - left) + 1
		if length > max {
			max = length
		}

	}

	return max
}

func longestDrome(s string) string {

	res := ""
	max := 0

	for i, _ := range s {
		l, r := i, i
		// slide out starting from the center expanding as far as we can
		// if l or r go out of bounds we stop
		// if the two characters we want to slide to are not the same we stop
		for l >= 0 && r < len(s) && s[l] == s[r] {

			// check if the existing palindrome is greater than the max
			if r-l+1 > max {
				max = r - l + 1
				res = s[l : r+1]
			}
			l--
			r++
		}

		// we need to do the same thing for even palindromes
		l, r = i, i+1

		for l >= 0 && r < len(s) && s[l] == s[r] {

			// check if the existing palindrome is greater than the max
			if r-l+1 > max {
				max = r - l + 1
				res = s[l : r+1]
			}
			l--
			r++
		}

	}

	return res
}

func mergeInterval(intervals [][]int) bool {

}

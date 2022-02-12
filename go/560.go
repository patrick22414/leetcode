package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	k := 3
	fmt.Println(subarraySum(nums, k))

	nums = []int{1, 1, 1}
	k = 2
	fmt.Println(subarraySum(nums, k))

	nums = []int{1, -1, 2, 0}
	k = 0
	fmt.Println(subarraySum(nums, k))

	nums = []int{1}
	k = 1
	fmt.Println(subarraySum(nums, k))
}

// The correct solution I found online, paraphrased and commented by myself.
func subarraySum(nums []int, k int) (count int) {
	// for brevity, I'll denote `sum of nums[0:i]` as `sum(0:i)` (not including nums[i])

	// store sum(0:i), for i from 0 to len(nums); aka `prefix sum`
	s := 0

	// store the freq of appearance of sum(0:j) for all j<i
	// `0` has already appeared once as sum(0:0) (empty array)
	freq := map[int]int{0: 1}

	for _, n := range nums {
		// we do 3 things in this loop

		// 1st, compute the prefix sum `s = sum(0:i+1)`
		s += n

		// 2nd, try to find any previous sums that satisfy k
		// that is, we want to know how many times `r = sum(0:j), j<=i` has appeared, so that `s-r = k`
		// there can be multiple `j`s because negative nums can appear
		// |---------- s ----------|
		// ============j===j=======i===========...
		// |------ r ------|-- k --|
		// |---- r ----|---- k ----|
		r := s - k
		count += freq[r] // Go returns 0 if key not in map

		// 3rd, update the freq table
		freq[s] += 1
	}

	return
}

// This is my solution.
// It works, but not very well.
func subarraySumMine(nums []int, k int) (count int) {
	sums := make([]int, len(nums))

	// populate sums with [0, i]
	sums[0] = nums[0]
	for i, n := range nums[1:] {
		sum := sums[i] + n
		if sum == k {
			count++
		}
		sums[i+1] = sum
	}

	// subtract by nums[j] each time
	for j, n := range nums {
		//fmt.Println(sums)
		for i, sum := range sums[j+1:] {
			sum -= n
			if sum == k {
				count++
			}
			sums[i+j+1] = sum
		}
	}

	return
}

package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:len(x)-1]
	z := append(y, -1)
	fmt.Println(x, y, z)

	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))

	nums = []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	sets := make([][]int, 1<<len(nums))
	head, tail := nums[0], nums[1:]

	for i, s := range subsets(tail) {
		sets[i*2] = s

		// ensure copy into a slice with +1 cap
		sCopy := make([]int, len(s), len(s)+1)
		copy(sCopy, s)
		sets[i*2+1] = append(sCopy, head)
	}

	return sets
}

func subsets2(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	head, tail := nums[0], nums[1:]

	headSubsets := make([][]int, 0)
	tailSubsets := subsets(tail)
	for _, ss := range tailSubsets {
		headSubsets = append(headSubsets, append(ss, head))
	}

	return append(tailSubsets, headSubsets...)
}

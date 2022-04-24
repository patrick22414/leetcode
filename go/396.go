package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}))
	fmt.Println(maxRotateFunction([]int{100}))
}

// Naive solution
func maxRotateFunction(nums []int) int {
	max := math.MinInt32

	for k := 0; k < len(nums); k++ {
		sum := 0
		for i := 0; i < len(nums); i++ {
			if i >= k {
				sum += i * nums[i-k]
			} else {
				sum += i * nums[len(nums)+i-k]
			}
		}

		fmt.Println(k, sum)

		if sum > max {
			max = sum
		}
	}

	return max
}

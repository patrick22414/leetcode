package main

import (
	"math"
	"strconv"
)

func main() {
}

func sequentialDigits(low int, high int) []int {
	results := make([]int, 0)

	n_digits_low := int(math.Log10(float64(low))) + 1
	n_digits_high := int(math.Log10(float64(high))) + 1

	for n := n_digits_low; n <= n_digits_high; n++ {
		for start := 0; start < 10-n; start++ {
			integer, _ := strconv.Atoi("123456789"[start : start+n])

			if integer > high {
				return results
			} else if integer >= low {
				results = append(results, integer)
			}
		}
	}

	return results
}

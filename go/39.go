package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}

func combinationSum(candidates []int, target int) [][]int {
	if target < 0 {
		return [][]int{}
	}
	if target == 0 {
		return [][]int{{}}
	}

	results := make([][]int, 0)
	for i, candidate := range candidates {
		subResults := combinationSum(candidates[i:], target-candidate)
		for j := 0; j < len(subResults); j++ {
			subResults[j] = append(subResults[j], candidate)
		}
		results = append(results, subResults...)
	}

	return results
}

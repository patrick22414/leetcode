package main

import (
	"fmt"
)

func main() {
	fmt.Println(consecutiveNumbersSum(1))
	fmt.Println(consecutiveNumbersSum(2))
	fmt.Println(consecutiveNumbersSum(3))
	fmt.Println(consecutiveNumbersSum(9))
	fmt.Println(consecutiveNumbersSum(15))
	fmt.Println(consecutiveNumbersSum(313653678))
	fmt.Println(consecutiveNumbersSum(933320757))
}

func consecutiveNumbersSum(n int) (count int) {
	count = 1

	// the idea is to divide `n` into `m` numbers
	// and see if we can redistribute them into `m` consecutive numbers
	for m := 2; ; m++ {
		i, j := n/m-(m-1)/2, n/m+m/2
		if i <= 0 {
			break
		}
		if (j+i)*(j-i+1)/2 == n {
			count++
		}
	}

	return
}

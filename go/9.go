package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 121
	fmt.Println(isPalindrome(x))

	x = 123
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

	s := []rune(strconv.Itoa(x))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

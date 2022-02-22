package main

import "fmt"

func main() {
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("100", 1))
	fmt.Println(removeKdigits("112", 1))
}

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	stack := make([]rune, 0)
	for _, char := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > char {
			stack = stack[:len(stack)-1]
			k--
		}

		if !(len(stack) == 0 && char == '0') {
			stack = append(stack, char)
		}
	}

	if len(stack) == 0 || len(stack) <= k {
		return "0"
	}
	if k > 0 {
		stack = stack[:len(stack)-k]
	}
	return string(stack)
}

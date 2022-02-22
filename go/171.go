package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("Z"))
	fmt.Println(titleToNumber("AA"))
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(columnTitle string) int {
	var col float64 = 0
	chars := []rune(columnTitle)

	for i, c := range chars {
		col += float64(c-'A'+1) * math.Pow(26, float64(len(columnTitle)-i-1))
	}

	return int(col)
}

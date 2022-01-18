package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	if wordPattern("abba", "dog cat cat dog") != true ||
		wordPattern("abba", "dog dog dog dog") != false ||
		wordPattern("aaa", "aa aa aa aa") != false {
		panic(errors.New("Test case failed"))
	} else {
		fmt.Println("Done")
	}
}

func wordPattern(pattern string, s string) bool {
	runes := []rune(pattern)
	words := strings.Fields(s)

	if len(runes) != len(words) {
		return false
	}

	book1 := make(map[rune]string)
	book2 := make(map[string]rune)

	for i := 0; i < len(runes); i++ {
		r := runes[i]
		w := words[i]
		if book1[r] == "" {
			book1[r] = w
		}
		if book2[w] == 0 {
			book2[w] = r
		}

		if book1[r] != w || book2[w] != r {
			return false
		}
	}

	return true
}

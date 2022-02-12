package main

import "fmt"

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}

// Breadth-first Search
func ladderLength(beginWord string, endWord string, wordList []string) (length int) {
	type Node struct {
		val string
		len int
	}

	queue := []Node{{beginWord, 1}}
	used := make([]bool, len(wordList))

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		fmt.Printf("%v -> ", curr)

		if curr.val == endWord {
			return curr.len
		}

		for i := 0; i < len(wordList); i++ {
			if !used[i] && differByALetter(curr.val, wordList[i]) {
				queue = append(queue, Node{
					wordList[i],
					curr.len + 1,
				})
				used[i] = true
			}
		}
	}

	return
}

func differByALetter(word1 string, word2 string) bool {
	count := 0

	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			count++
		}
	}

	return count == 1
}

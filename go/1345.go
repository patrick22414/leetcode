package main

import "fmt"

func main() {
	input := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
	output := minJumps(input)

	fmt.Println(output)
}

func minJumps(arr []int) int {
	if len(arr) == 1 {
		return 0
	}

	indices := make(map[int][]int)
	for i, n := range arr {
		if indices[n] == nil {
			indices[n] = make([]int, 0)
		}

		indices[n] = append(indices[n], i)
	}

	q := make([]int, 0)
	seen := make([]bool, len(arr))
	dist := make([]int, len(arr))

	q = append(q, 0)
	seen[0] = true

	for len(q) > 0 {
		index := q[0]
		value := arr[index]
		newDist := dist[index] + 1

		for _, x := range []int{index - 1, index + 1} {
			if x >= 0 && x < len(arr) && !seen[x] {
				if x == len(arr)-1 {
					return newDist
				}

				q = append(q, x)
				seen[x] = true
				dist[x] = newDist
			}
		}

		if indices[value] != nil {
			for _, x := range indices[value] {
				if seen[x] {
					continue
				}
				if x == len(arr)-1 {
					return newDist
				}
				q = append(q, x)
				seen[x] = true
				dist[x] = newDist
			}
			delete(indices, value)
		}

		q = q[1:]
	}

	return -1
}

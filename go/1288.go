package main

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	a, b, removed := -1, -1, 0
	for _, interval := range intervals {
		if interval[1] <= b {
			removed++
		} else if interval[0] == a {
			removed++
			a, b = interval[0], interval[1]
		} else {
			a, b = interval[0], interval[1]
		}
	}

	return len(intervals) - removed
}

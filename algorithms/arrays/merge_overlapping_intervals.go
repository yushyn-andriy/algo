package arrays

import (
	"sort"
)

const (
	START = 0
	END   = 1
)

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	merged := [][]int{}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][START] < intervals[j][START]
	})
	cidx := 0
	merged = append(merged, intervals[cidx])
	for i := 1; i < len(intervals); i++ {
		curr := merged[cidx]
		next := intervals[i]
		if curr[END] >= next[START] {
			merged[cidx][END] = max(curr[END], next[END])
		} else {
			merged = append(merged, next)
			cidx++
		}
	}
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

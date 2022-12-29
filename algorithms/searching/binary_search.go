package searching

import (
	"fmt"
	"sort"
)

func BinarySearch(v int, arr []int) (int, bool) {
	if len(arr) == 0 {
		return 0, false
	}
	getMiddle := func(left, right int) int {
		return (left + right) / 2
	}
	sort.Ints(arr)

	l := 0
	r := len(arr) - 1
	m := getMiddle(l, r)

	for {
		fmt.Println(l, r, m)
		if arr[m] == v {
			return m, true
		}
		if m >= len(arr)-1 || m == 0 {
			return 0, false
		}

		if arr[m] < v {
			l = m + 1
			m = getMiddle(l, r)
		} else {
			r = m
			m = getMiddle(l, r)
		}
	}
	return 0, false
}

func BinarySearchAlgoSolution1(array []int, target int) int {
	if len(array) == 0 {
		return -1
	}
	getMiddle := func(left, right int) int {
		return (left + right) / 2
	}
	sort.Ints(array)

	l := 0
	r := len(array) - 1
	m := getMiddle(l, r)

	for {
		currValue := array[m]
		if currValue == target {
			return m
		}
		if m >= len(array)-1 || m == 0 || l == r {
			return -1
		}

		if currValue < target {
			l = m + 1
			m = getMiddle(l, r)
		} else {
			r = m
			m = getMiddle(l, r)
		}
	}
}

package arrays

import (
	"fmt"
	"sort"
)

func ThreeNumbers(arr []int, targetSum int) ([][]int, error) {
	sort.Ints(arr)

	sum := [][]int{}
	if len(arr) < 3 {
		return sum, fmt.Errorf("Len of array must be greater than 2")
	}

	var curNum, l, r int
	for i := 0; i <= len(arr)-2; i++ {
		curNum = arr[i]
		l = i + 1
		r = len(arr) - 1

		for l < r {
			leftNum, rightNum := arr[l], arr[r]
			curSum := curNum + leftNum + rightNum
			if curSum == targetSum {
				ss := []int{curNum, leftNum, rightNum}
				sort.Ints(ss)
				sum = append(sum, ss)
				l++
				r--
			} else if curSum < targetSum {
				l++
			} else if curSum > targetSum {
				r--
			}
		}
	}

	return sum, nil
}

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)

	sum := [][]int{}
	if len(array) < 3 {
		return sum
	}

	var curNum, l, r int
	for i := 0; i <= len(array)-2; i++ {
		curNum = array[i]
		l = i + 1
		r = len(array) - 1

		for l < r {
			leftNum, rightNum := array[l], array[r]
			curSum := curNum + leftNum + rightNum
			if curSum == target {
				ss := []int{curNum, leftNum, rightNum}
				sort.Ints(ss)
				sum = append(sum, ss)
				l++
				r--
			} else if curSum < target {
				l++
			} else if curSum > target {
				r--
			}
		}
	}

	return sum
}

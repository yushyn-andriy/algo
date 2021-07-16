package arrays

import "math"

const (
	FIRST = iota
	SECOND
	THIRD
)

type ElMax struct {
	IsSet bool
	Value int
}

type ThreeMax []ElMax

func ThreeNumbersMax(arr []int) ThreeMax {
	threeMaxArr := ThreeMax{
		{false, 0},
		{false, 0},
		{false, 0},
	}
	for _, v := range arr {
		if !threeMaxArr[THIRD].IsSet {
			threeMaxArr[THIRD].IsSet = true
			threeMaxArr[THIRD].Value = v
		} else {
			shiftUpdate(threeMaxArr, v)
		}
	}
	return threeMaxArr
}

func shiftUpdate(arr ThreeMax, value int) {
	if value > arr[THIRD].Value {
		second := arr[SECOND]
		arr[SECOND] = arr[THIRD]
		arr[FIRST] = second
		arr[THIRD] = ElMax{true, value}
	} else if value > arr[SECOND].Value {
		if !arr[SECOND].IsSet {
			arr[SECOND].IsSet = true
			arr[SECOND].Value = value
		} else {
			arr[FIRST] = arr[SECOND]
			arr[SECOND] = ElMax{true, value}
		}
	} else if value > arr[FIRST].Value || !arr[FIRST].IsSet {
		arr[FIRST].Value = value
		arr[FIRST].IsSet = true
	}
}

func FindThreeLargestNumbersAlgoSoulution1(array []int) []int {
	max := []int{
		math.MinInt32,
		math.MinInt32,
		math.MinInt32,
	}
	for _, v := range array {
		if max[THIRD] == math.MinInt32 {
			max[THIRD] = v
		} else {
			shiftUpdateAlgoSoulution1(max, v)
		}
	}
	return max
}

func shiftUpdateAlgoSoulution1(max []int, value int) []int {

	if value > max[THIRD] {
		second := max[SECOND]
		max[SECOND] = max[THIRD]
		max[FIRST] = second
		max[THIRD] = value
	} else if value > max[SECOND] {
		if max[SECOND] == math.MinInt32 {
			max[SECOND] = value
		} else {
			max[FIRST] = max[SECOND]
			max[SECOND] = value
		}
	} else if value > max[FIRST] || max[FIRST] == math.MinInt32 {
		max[FIRST] = value
	}

	return max
}

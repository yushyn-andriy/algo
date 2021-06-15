package recursion

import (
	"math"
	"sort"
)

func Factors(fn []int) [][][]int {
	r := [][][]int{}
	factors(fn, len(fn), [][]int{}, &r)
	return r
}

func factors(fn []int, fnLen int, currCombination [][]int, r *[][][]int) {
	if len(fn) == 0 {
		toAppend := make([][]int, len(currCombination))
		copy(toAppend, currCombination)
		*r = append(*r, toAppend)
		return
	}

	for i := 1; i <= fnLen; i++ {
		tmp := []int{}
		if len(fn) < i {
			break
		}
		for j := 0; j < i; j++ {
			if len(fn) == 0 {
				break
			}
			el := fn[len(fn)-1]
			fn = fn[:len(fn)-1]
			tmp = append(tmp, el)
		}
		appTmp := make([]int, len(tmp))
		copy(appTmp, tmp)

		currCombination = append(currCombination, appTmp)
		factors(fn, fnLen, currCombination, r)
		currCombination = currCombination[:len(currCombination)-1]
		for j := 0; j < i; j++ {
			if len(tmp) == 0 {
				break
			}
			el := tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
			fn = append(fn, el)
		}
	}
}

func FlattenFactorCombinations(x [][][]int) [][]int {
	flatten := [][]int{}
	for _, factors := range x {
		r := []int{}
		for _, combinations := range factors {
			r = append(r, mul(combinations))
		}
		flatten = append(flatten, r)
	}
	return flatten
}

func mul(x []int) int {
	s := 1
	for _, v := range x {
		s *= v
	}
	return s
}

func FindNumbers(x [][]int) []int {
	y := []int{}
	var primeNumbers []int = []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 71,
		73, 79, 83, 90, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151,
	}

	r := 1
	for _, combination := range x {
		sort.Sort(sort.Reverse(sort.IntSlice(combination)))
		for i, p := range combination {
			r *= int(math.Pow(float64(primeNumbers[i]), float64(p-1)))
		}
		if r > 0 {
			y = append(y, r)
		}
		r = 1
	}

	return y
}

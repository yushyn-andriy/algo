package recursion

func Powerset(array []int) [][]int {
	sets := [][]int{}
	powerset(array, []int{}, 0, &sets, &map[int]bool{})
	return sets
}

func powerset(array []int, currSet []int, k int, sets *[][]int, seen *map[int]bool) {
	if len(array) == k {
		arr := make([]int, len(currSet))
		copy(arr, currSet)
		*sets = append(*sets, arr)
		return
	}

	currSet = append(currSet, array[k])
	powerset(array, currSet, k+1, sets, seen)
	currSet = currSet[:len(currSet)-1]
	powerset(array, currSet, k+1, sets, seen)
}

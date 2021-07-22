package recursion

func GetPermutations(array []int) [][]int {
	permutations := [][]int{}
	if len(array) == 0 {
		return permutations
	}

	getPermutations(array, []int{}, &permutations, map[int]bool{})
	return permutations
}

func getPermutations(array []int, currPermutation []int, permutations *[][]int, seen map[int]bool) {
	if len(array) == len(currPermutation) {
		arr := make([]int, len(currPermutation))
		copy(arr, currPermutation)
		*permutations = append(*permutations, arr)
		return
	}

	for _, v := range array {
		if isSeen := seen[v]; isSeen {
			continue
		}

		seen[v] = true
		currPermutation = append(currPermutation, v)
		getPermutations(array, currPermutation, permutations, seen)
		currPermutation = currPermutation[:len(currPermutation)-1]
		seen[v] = false
	}
}

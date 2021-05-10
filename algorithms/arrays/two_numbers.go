package main


import "fmt"

func main() {
	a := []int{3, 5, -4, 8, 11, 1, -1, 6}
	hashMap := make(map[int]bool)
	targetSum := 10

	for _, n := range a  {
		hashMap[n] = true
	}

	result := make([]int, 0)
	for _, n := range a  {
		diff := targetSum - n
		hashMap[n] = false
		if val, ok := hashMap[diff]; ok && val {
			result = append(result, n)
		}
		hashMap[n] = true
	}

	for _, n := range result {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()
}

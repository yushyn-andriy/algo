package sorting_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/sorting"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{2, 0, 1, 3}
	dup := make([]int, len(arr))
	copy(dup, arr)

	exp := []int{0, 1, 2, 3}
	sorting.InsertionSort(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != exp[i] {
			t.Errorf("insertionSort(%v): got %v expected %v", dup, arr, exp)
			break
		}
	}
}

func TestRecursiveInsertionSort(t *testing.T) {
	arr := []int{2, 0, 1, 3}
	dup := make([]int, len(arr))
	copy(dup, arr)

	exp := []int{0, 1, 2, 3}
	sorting.RecursiveInsertionSort(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != exp[i] {
			t.Errorf("recursiveInsertionSort(%v): got %v expected %v", dup, arr, exp)
			break
		}
	}
}

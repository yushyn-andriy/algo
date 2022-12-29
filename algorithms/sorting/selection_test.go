package sorting_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/sorting"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		tosort   []int
		expected []int
	}{
		{
			[]int{4, 2, 1, 6, -1},
			[]int{-1, 1, 2, 4, 6},
		},
		{
			[]int{8, 5, 2, 9, 5, 6, 3},
			[]int{2, 3, 5, 5, 6, 8, 9},
		},
	}

	for i, test := range tests {
		sorted := sorting.SelectionSort(test.tosort)
		if len(sorted) != len(test.expected) {
			t.Errorf("test(%d) Expected %v got %v", i, test.expected, sorted)
			return
		}
		for i := 0; i < len(test.expected); i++ {
			if sorted[i] != test.expected[i] {
				t.Errorf("test(%d) Expected %v got %v", i, test.expected, sorted)
				return
			}
		}
	}
}

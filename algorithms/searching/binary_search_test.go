package searching_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/searching"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr                  []int
		value, expectedIndex int
		expectedFound        bool
	}{
		{[]int{1, 2}, 1, 0, true},
		{[]int{5, 6, 2}, 6, 2, true},
		{[]int{5, 6, 2, -2, 10}, 100, 0, false},
		{[]int{}, 100, 0, false},
		{[]int{5, 6, 2, -2, 10}, -2, 0, true},
		{[]int{85, 36, 25, 55, 22, 15}, 36, 3, true},
	}
	for i, test := range tests {
		index, found := searching.BinarySearch(test.value, test.arr)
		if index != test.expectedIndex {
			t.Errorf("test(%d) error expected index=%v, got=%v", i, test.expectedIndex, index)
			return
		}
		if found != test.expectedFound {
			t.Errorf("test(%d) error expected found=%v, found=%v", i, test.expectedFound, found)
			return
		}
	}
}

func TestBinarySearchAlgoSolution1(t *testing.T) {
	tests := []struct {
		arr                  []int
		value, expectedIndex int
	}{
		{[]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 70, -1},
	}
	for i, test := range tests {
		idx := searching.BinarySearchAlgoSolution1(test.arr, test.value)
		if idx != test.expectedIndex {
			t.Errorf("test(%d) error expected index=%v, got=%v", i, test.expectedIndex, idx)
		}
	}
}

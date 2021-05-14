package arrays_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestTwoNumbers(t *testing.T) {
	s := [][]int{
		{3, 5, -4, 8, 11, 1, -1},
		{3, 2, 8},
		{3, 4, 3, 7},
		{-18, 4, 3, 7, 28},
	}
	expNumbers := [][]int{
		{11, -1},
		{2, 8},
		{3, 7},
		{-18, 28},
	}

	for index, arr := range s {
		targetSum := 10

		result := arrays.TwoNumbers(arr, targetSum)
		if len(result) != len(expNumbers[index]) {
			t.Errorf("Expected equal array sizes, got len(result)=%d, len(expNumbers)=%d", len(result), len(expNumbers))
		}
		for i, v := range result {
			if expNumbers[index][i] != v {
				t.Errorf("Expected value is %d, got %d", expNumbers[index][i], v)
			}
		}
	}
}

func BenchmarkTwoNumbers(b *testing.B) {
	arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10
	for i := 0; i < b.N; i++ {
		arrays.TwoNumbers(arr, targetSum)
	}
}

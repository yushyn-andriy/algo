package trees_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestRightSmallerThan(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{8, 5, 11, -1, 3, 4, 2},
			[]int{5, 4, 4, 0, 1, 1, 0},
		},
	}

	for i, test := range tests {
		out := trees.RightSmallerThan1(test.in)
		if !IsEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}

func TestRightSmallerThanSolution2(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{8, 5, 11, -1, 3, 4, 2},
			[]int{5, 4, 4, 0, 1, 1, 0},
		},
		{
			[]int{0, 1, 1, 2, 3, 5, 8, 13},
			[]int{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1, 0},
			[]int{1, 0},
		},
		{
			[]int{0, 1},
			[]int{0, 0},
		},
	}

	for i, test := range tests {
		out := trees.RightSmallerThan2(test.in)
		if !IsEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}

func IsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

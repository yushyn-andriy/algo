package arrays_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestValidateSubsequante(t *testing.T) {
	testCases := []struct {
		sequent      []int
		subSequent   []int
		expectResult bool
	}{
		{
			[]int{2, 5, 4, -1, 6, 8},
			[]int{5, -1, 8},
			true,
		},
		{
			[]int{5, 1, 22, 25, 6, -1, 8, 10},
			[]int{1, 6, -1, 10},
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{2, 3, 5, 6},
			false,
		},
	}
	for _, testCase := range testCases {
		if res := arrays.ValidateSubsequent(testCase.sequent, testCase.subSequent); res != testCase.expectResult {
			t.Errorf("Expected %v, got %v", testCase.expectResult, res)
		}
	}

}

func BenchmarkValidateSubsequante(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrays.ValidateSubsequent([]int{2, 5, 4, -1, 6, 8}, []int{5, -1, 8})
	}
}

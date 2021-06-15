package recursion_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/recursion"
)

func TestFactors(t *testing.T) {
	tests := []struct {
		fn []int
		cs [][][]int
	}{
		{
			[]int{1, 2, 3},
			[][][]int{
				{
					{3},
					{2},
					{1},
				},
				{
					{3},
					{2, 1},
				},
				{
					{3, 2},
					{1},
				},
				{
					{3, 2, 1},
				},
			},
		},
		{
			[]int{1, 2},
			[][][]int{
				{
					{2},
					{1},
				},
				{
					{2, 1},
				},
			},
		},
		{
			[]int{2, 2, 17},
			[][][]int{
				{
					{17},
					{2},
					{2},
				},
				{
					{17},
					{2, 2},
				},
				{
					{17, 2},
					{2},
				},
				{
					{17, 2, 2},
				},
			},
		},
	}

	for i, test := range tests {
		factors := recursion.Factors(test.fn)
		if len(factors) != len(test.cs) {
			t.Errorf("test(%d) Expected len %d got %d", i, len(factors), len(test.cs))
			continue
		}
		for j, l := range factors {
			if !reflect.DeepEqual(test.cs[j], l) {
				t.Errorf("test(%d) subtest(%d) expected %v got %v", i, j, test.cs[j], l)
			}
		}
	}
}

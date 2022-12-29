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

func TestFlattenFactorCombinations(t *testing.T) {
	tests := []struct {
		cs [][][]int
		cm [][]int
	}{
		{
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
			[][]int{
				{3, 2, 1},
				{3, 2},
				{6, 1},
				{6},
			},
		},

		{
			[][][]int{
				{
					{2},
					{1},
				},
				{
					{2, 1},
				},
			},
			[][]int{
				{2, 1},
				{2},
			},
		},
		{
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
			[][]int{
				{17, 2, 2},
				{17, 4},
				{34, 2},
				{68},
			},
		},
	}

	for i, test := range tests {
		cm := recursion.FlattenFactorCombinations(test.cs)
		if len(cm) != len(test.cm) {
			t.Errorf("test(%d) Expected %v got %v", i, cm, test.cm)
			continue
		}
		for j, l := range cm {
			if !reflect.DeepEqual(test.cm[j], l) {
				t.Errorf("test(%d) subtest(%d) expected %v got %v", i, j, test.cm[j], l)
			}
		}
	}
}

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		cm [][]int
		ns []int
	}{
		{
			[][]int{
				{3, 2},
				{6},
			},
			[]int{12, 32},
		},
		{
			[][]int{
				{11, 3, 2},
				{33, 2},
				{11, 6},
				{65},
			},
			[]int{46080, 12884901888, 248832},
		},
	}

	for i, test := range tests {
		ns := recursion.FindNumbers(test.cm)
		if !reflect.DeepEqual(ns, test.ns) {
			t.Errorf("test(%d) Expected %v got %v", i, test.ns, ns)
		}
	}
}

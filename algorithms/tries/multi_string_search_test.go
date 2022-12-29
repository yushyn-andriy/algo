package tries

import (
	"reflect"
	"testing"
)

var cases []struct {
	bigString    string
	smallStrings []string
	expected     []bool
} = []struct {
	bigString    string
	smallStrings []string
	expected     []bool
}{
	{
		"this is a big string",
		[]string{"this", "yo", "is", "a", "bigger", "string", "kappa"},
		[]bool{true, false, true, true, false, true, false},
	},
	{
		"abcdefghijklmnopqrstuvwxyz",
		[]string{"abc", "mnopqr", "wyz", "no", "e", "tuuv"},
		[]bool{true, true, false, true, true, false},
	},
	{
		"abcdefghijklmnopqrstuvwxyz",
		[]string{"abcdefghijklmnopqrstuvwxyz", "abc", "j", "mnopqr", "pqrstuvwxyz", "xyzz", "defh"},
		[]bool{true, true, true, true, true, false, false},
	},
	{
		"hj!)!%Hj1jh8f1985n!)51",
		[]string{"%Hj7", "8f198", "!)5", "!)!", "!!", "jh81", "j181hf"},
		[]bool{false, true, true, true, false, false, false},
	},
}

func TestMultiStringSearch1(t *testing.T) {
	for i, testCase := range cases {
		out := MultiStringSearch1(testCase.bigString, testCase.smallStrings)
		if !reflect.DeepEqual(out, testCase.expected) {
			t.Errorf("%d error expected %v but got %v", i, testCase.expected, out)
		}
	}
}

func TestMultiStringSearch2(t *testing.T) {
	for i, testCase := range cases {
		out := MultiStringSearch2(testCase.bigString, testCase.smallStrings)
		if !reflect.DeepEqual(out, testCase.expected) {
			t.Errorf("%d error expected %v but got %v", i, testCase.expected, out)
		}
	}
}

func TestMultiStringSearch3(t *testing.T) {
	for i, testCase := range cases {
		out := MultiStringSearch3(testCase.bigString, testCase.smallStrings)
		if !reflect.DeepEqual(out, testCase.expected) {
			t.Errorf("%d error expected %v but got %v", i, testCase.expected, out)
		}
	}
}

func TestSuffixTrie1(t *testing.T) {
	tests := []struct {
		populateStr, containsStr string
		contains                 bool
	}{
		{
			"acdc",
			"dc",
			true,
		},
		{
			"acdc",
			"da",
			false,
		},
		{
			"algo and datastructure",
			"datastructure",
			true,
		},
		{
			"this is a big string",
			"this",
			true,
		},
	}

	for i, test := range tests {
		trie := NewSuffixTrie1()
		trie.Populate(test.populateStr)
		out := trie.Contains(test.containsStr)
		if out != test.contains {
			t.Errorf("test(%d) expected %v got %v", i, test.contains, out)
		}
	}

}


func TestSortingInPlace(t *testing.T) {
	cases := []struct{
		arr []int
		expected []int
	}{
		{
			[]int{5, 4, 7, 1, 2},
			[]int{1, 2, 4, 5, 7},
		},
		{
			[]int{5, 4, 7, 1, 2, 3},
			[]int{1, 2, 3, 4, 5, 7},
		},
	}

	for _, c := range cases {
		sortingInPlace(c.arr)
		if !reflect.DeepEqual(c.arr, c.expected) {
			t.Errorf("expected %v got %v", c.expected, c.arr)
		}
	}

}
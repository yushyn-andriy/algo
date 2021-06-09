package dynamic_programming_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/dynamic_programming"
)

func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		s1, s2 string
		d      int // Levenshtein Distance
	}{
		{
			"abc",
			"abc",
			0,
		},
		{
			"abc",
			"abf",
			1,
		},
		{
			"yabd",
			"abc",
			2,
		},
		{
			"levenshteinDist(s1, s2, i, j-1),",
			"levenshteinDist(s1, s2, i-1, j),",
			4,
		},
		{
			"pol",
			"exl",
			2,
		},
		{
			"Levenshtein distance is a string metric for measuring the difference between two sequences. Informally, the Levenshtein distance between two words is the minimum number of single-character edits (insertions, deletions or substitutions) required to change one word into the other.",
			"levenshtein istance is a string metric for measuring the difference between two sequences. Informally, the Levenshtein distance between two words is the minimum number of single-character edtts required to change one word into the other.",
			43,
		},
	}

	for i, test := range tests {
		d := dynamic_programming.LevenshteinDistance(test.s1, test.s2)
		if test.d != d {
			t.Errorf("test(%d) expected distance %d but got %d", i, test.d, d)
		}
	}

}

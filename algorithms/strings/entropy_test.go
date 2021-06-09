package strings_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/strings"
)

func TestEntropy(t *testing.T) {
	tests := []struct {
		s string  // input string
		e float64 // expected entropy
	}{
		{
			"AAAAABBCDE",
			1.96,
		},
		{
			"AB",
			1.0,
		},
		{
			"A",
			0.0,
		},
		{
			"AAB",
			0.91,
		},
	}

	for i, test := range tests {
		e := strings.Entropy(test.s)
		if e != test.e {
			t.Errorf("test(%d) expected entropy is %v but got %v", i, test.e, e)
		}
	}
}

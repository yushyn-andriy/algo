package strings

import "testing"

func TestSmallestSubstringContaining(t *testing.T) {
	tests := []struct {
		bigString, smallString, out string
	}{
		{
			"abcd$ef$axb$c$",
			"$$abf",
			"f$axb$",
		},
	}

	for i, test := range tests {
		out := SmallestSubstringContaining(test.bigString, test.smallString)
		if test.out != out {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}

}

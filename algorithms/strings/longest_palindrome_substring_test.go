package strings

import (
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			"abaxyzzyxf",
			"xyzzyx",
		},
	}

	for i, test := range tests {
		out := LongestPalindromicSubstring(test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}

package strings

import "testing"

func TestLongestSubstringWithoutDuplication(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"clementisacap", "mentisac"},
	}

	for i, test := range tests {
		out := LongestSubstringWithoutDuplication(test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %q got %q", i, test.out, out)
		}

	}
}

package strings

import "testing"

func TestFirstNonRepeatingCharacter(t *testing.T) {
	tests := []struct {
		in    string
		index int
	}{
		{"abcdcaf", 1},
	}

	for i, test := range tests {
		out := FirstNonRepeatingCharacterSolution1(test.in)
		if out != test.index {
			t.Errorf("test(%d) expected %v got %v", i, test.index, out)
		}
	}
}

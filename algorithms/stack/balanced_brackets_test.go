package stack

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		str        string
		isBalanced bool
	}{
		{"(a)", true},
		{"([)]", false},
		{"([])", true},
	}

	for i, test := range tests {
		isBalanced := IsBalanced(test.str)
		if isBalanced != test.isBalanced {
			t.Errorf("test(%d) expected %v got %v", i, test.isBalanced, isBalanced)
		}
	}
}

package strings

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		str        string
		isBalanced bool
	}{
		{"()", true},
		{"())", false},
		{"((())()", false},
		{"((()()))", true},
	}

	for i, test := range tests {
		isBalanced := IsBalanced(test.str)
		if isBalanced != test.isBalanced {
			t.Errorf("test(%d) expected %v got %v", i, test.isBalanced, isBalanced)
		}
	}
}

func TestLongestBalancedSubstring(t *testing.T) {
	tests := []struct {
		str        string
		longestMax int
	}{
		{"()", 2},
		{"())", 2},
		{"()())", 4},
		{"((())()", 6},
		{"((()()))", 8},
		{"(()))(", 4},
	}

	for i, test := range tests {
		longestMax := LongestBalancedSubstring(test.str)
		if longestMax != test.longestMax {
			t.Errorf("test(%d) expected %v got %v", i, test.longestMax, longestMax)
		}
	}
}

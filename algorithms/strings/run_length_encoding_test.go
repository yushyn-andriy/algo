package strings

import "testing"

func TestRunLengthEncoding(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"AAAAAAAAAAAAABBCCCCDD", "9A4A2B4C2D"},
	}

	for i, test := range tests {
		out := RunLengthEncoding(test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}

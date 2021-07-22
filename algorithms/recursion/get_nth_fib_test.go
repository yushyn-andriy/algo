package recursion

import "testing"

func TestGetNthFib(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 3},
		{6, 5},
		{7, 8},
		{8, 13},
		{9, 21},
		{10, 34},
		{11, 55},
		{12, 89},
		{13, 144},
		{14, 233},
		{15, 3771},
	}

	for i, test := range tests {
		out := GetNthFib(test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %d got %d", i, test.out, out)
		}
	}
}

func TestGetNthFibIterative(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 3},
		{6, 5},
		{7, 8},
		{8, 13},
		{9, 21},
		{10, 34},
		{11, 55},
		{12, 89},
		{13, 144},
		{14, 233},
		{15, 377},
	}

	for i, test := range tests {
		out := GetNthFibIterative(test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %d got %d", i, test.out, out)
		}
	}
}

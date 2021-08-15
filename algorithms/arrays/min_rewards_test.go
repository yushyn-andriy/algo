package arrays

import "testing"

func TestMergeOverlappingIntervals(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 2, 3, 1}, //1, 2, 3, 1
			7,
		},
	}

	for i, test := range tests {
		out := MinRewards(test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
			return
		}
	}
}

package greedy

import (
	"testing"
)

func TestValidStartingCity(t *testing.T) {
	tests := []struct {
		distances, fuel []int
		mpg, out        int
	}{
		{
			[]int{5, 25, 15, 10, 15},
			[]int{1, 2, 1, 0, 3},
			10,
			4,
		},
	}

	for i, test := range tests {
		out := ValidStartingCity(test.distances, test.fuel, test.mpg)
		if out != test.out {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
		}
	}
}

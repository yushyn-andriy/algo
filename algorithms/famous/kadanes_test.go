package famous_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/famous"
)

func TestKadanes(t *testing.T) {
	tests := []struct {
		x    []int
		best int
	}{
		{
			[]int{3, 5, -9, 1, 3, -2, 3, 4, 7},
			16,
		},
		{
			[]int{3, 5, -5, 1, 1},
			8,
		},
	}

	for i, test := range tests {
		best := famous.KadanesAlgorithm(test.x)
		if test.best != best {
			t.Errorf("test(%d) expect %#v got %#v", i, test.best, best)
		}
	}
}

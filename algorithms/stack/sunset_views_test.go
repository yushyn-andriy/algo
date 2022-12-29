package stack

import (
	"reflect"
	"testing"
)

func TestSunsetViews(t *testing.T) {
	tests := []struct {
		buildings, out []int
		direction      string
	}{
		{
			[]int{3, 5, 4, 4, 3, 1, 3, 2},
			[]int{1, 3, 6, 7},
			"EAST",
		},
	}

	for i, test := range tests {
		out := SunsetViews(test.buildings, test.direction)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v  got %v", i, test.out, out)
		}
	}
}

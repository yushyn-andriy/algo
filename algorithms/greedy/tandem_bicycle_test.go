package greedy

import "testing"

func TestTandemBicycle(t *testing.T) {
	tests := []struct {
		redShirtSpeeds, blueShirtSpeeds []int
		fastest                         bool
		out                             int
	}{
		{
			[]int{},
			[]int{},
			true,
			32,
		},
	}

	for i, test := range tests {
		out := TandemBicycle(test.redShirtSpeeds, test.blueShirtSpeeds, test.fastest)
		if out != test.out {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
		}
	}
}

package greedy

import (
	"testing"
)

func TestClassPhotos(t *testing.T) {
	tests := []struct {
		redShirtHeights, blueShirtHeights []int
		out                               bool
	}{}

	for i, test := range tests {
		out := ClassPhotos(test.redShirtHeights, test.blueShirtHeights)
		if out != test.out {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
		}
	}
}

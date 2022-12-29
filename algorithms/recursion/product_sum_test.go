package recursion_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/recursion"
)

func TestProductSum(t *testing.T) {
	parr := recursion.ProductArr{
		1,
		2,
		recursion.ProductArr{2, 3},
		3,
	}
	sum := recursion.ProductSum(parr, 1)
	if sum != 16 {
		t.Errorf("Expected 16 got=%v", sum)
	}
}

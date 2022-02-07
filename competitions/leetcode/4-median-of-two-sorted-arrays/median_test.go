package median

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays1(t *testing.T) {
	cases := []struct {
		nums1, nums2 []int
		expected     float64
	}{
		{
			[]int{1, 3},
			[]int{2},
			2.00000,
		},
		{
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
		{
			[]int{},
			[]int{},
			0.0,
		},
		{
			[]int{1},
			[]int{},
			1,
		},
		{
			[]int{},
			[]int{2},
			2,
		},
		{
			[]int{1, 3, 4, 5},
			[]int{4},
			4,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, findMedianSortedArrays1(c.nums1, c.nums2))
	}
}

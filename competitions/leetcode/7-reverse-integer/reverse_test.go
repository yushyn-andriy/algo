package reverse

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		x   int
		exp int
	}{
		{123, 321},
		{1, 1},
		{10, 1},
		{math.MaxInt32 + 1, 0},
	}

	for _, c := range cases {
		res := reverse(c.x)
		assert.Equal(t, c.exp, res)
	}
}

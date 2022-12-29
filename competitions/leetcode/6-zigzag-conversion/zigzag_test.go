package zigzagconversion

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		s, expected string
		numRows     int
	}{
		{
			"PAYPALISHIRING",
			"PINALSIGYAHRPI",
			4,
		},
		{
			"PAYPALISHIRING",
			"PAHNAPLSIIGYIR",
			3,
		},
		{
			"A",
			"A",
			1,
		},
		{
			"ABC",
			"ACB",
			2,
		},
		{
			"ABC",
			"ABC",
			3,
		},
		{
			"ABC",
			"ABC",
			1,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, convert(c.s, c.numRows))
	}
}

func BenchmarkConvert(b *testing.B) {
	s := strings.Repeat("PAYPALISHIRING", 100)
	for i := 0; i < b.N; i++ {
		convert(s	, 4)
	}
}

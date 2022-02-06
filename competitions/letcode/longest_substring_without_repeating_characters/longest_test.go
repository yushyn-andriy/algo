package longestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		s      string
		length int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"   ab", 3},
		{" ", 1},
	}

	for _, c := range cases {
		assert.Equal(t, c.length, LengthOfLongestSubstring(c.s))
	}
}

func TestLengthOfLongestSubstring1(t *testing.T) {
	cases := []struct {
		s      string
		length int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"   ab", 3},
		{" ", 1},
	}

	for _, c := range cases {
		assert.Equal(t, c.length, lengthOfLongestSubstring(c.s))
	}
}

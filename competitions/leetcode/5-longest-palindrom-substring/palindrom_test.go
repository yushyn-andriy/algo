package palindrom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrom(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{"abba", true},
		{"abc", false},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, IsPalindrom(c.s))
	}
}

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{"abba", "abba"},
		{"abc", "a"},
		{"cbbd", "bb"},
		{"babad", "bab"},
		{"a", "a"},
		{"", ""},
		{"abcd", "a"},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, longestPalindrome(c.s))
	}
}

func TestLongestPalindrome1(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{"abba", "abba"},
		{"abc", "a"},
		{"cbbd", "bb"},
		{"babad", "bab"},
		{"a", "a"},
		{"", ""},
		{"abcd", "a"},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, longestPalindrome1(c.s))
	}
}

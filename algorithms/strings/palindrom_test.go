package strings_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/strings"
)

func TestIsPalindrom(t *testing.T) {
	tests := []struct {
		word                string
		expectedIsPalindrom bool
	}{
		{"anna", true},
		{"civic", true},
		{"mom", true},
		{"noon", true},
		{"no", false},
		{"civic", true},
		{"kayak", true},
		{"level", true},
		{"madam", true},
		{"mom", true},
		{"noon", true},
		{"racecar", true},
		{"radar", true},
		{"redder", true},
		{"refer", true},
		{"repaper", true},
		{"rotator", true},
		{"rotor", true},
		{"sagas", true},
		{"solos", true},
		{"soloos", false},
		{"stats", true},
		{"tenet", true},
		{"wow", true},
		{"false", false},
		{"palindromes", false},
		{"pets", false},
		{"character", false},
		{"matches", false},
		{"ignored", false},
		{"most", false},
		{"a", true},
		{"", true},
		{"a", true},
		{"abcdcba", true},
	}

	for i, test := range tests {
		isPalindrom := strings.IsPalindrom(test.word)
		if isPalindrom != test.expectedIsPalindrom {
			t.Errorf("test(%d) Expected word=\"%s\", isPalindrome=%v got isPalindrome=%v", i, test.word, test.expectedIsPalindrom, isPalindrom)
			return
		}
	}
}

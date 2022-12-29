package main

import "testing"

func TestNumbers(t *testing.T) {
	tests := []struct {
		s string
		n int
	}{
		{
			"12343",
			5,
		},
		{
			"985454",
			6,
		},
	}

	for _, test := range tests {
		n := Numbers(test.s)
		if n != test.n {
			t.Errorf("Expected %d got %d", test.n, n)
		}
	}
}

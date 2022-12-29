package arrays_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestTournamentWinner(t *testing.T) {
	tests := []struct {
		competitions [][]string
		results      []int
		out          string
	}{
		{
			[][]string{
				{"HTML", "C#"},
				{"C#", "Python"},
				{"Python", "HTML"},
			},
			[]int{0, 0, 1},
			"Python",
		},
	}

	for i, test := range tests {
		out := arrays.TournamentWinner(test.competitions, test.results)
		if out != test.out {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
			return
		}
	}
}

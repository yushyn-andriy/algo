package graph

import (
	"reflect"
	"testing"
)

func TestBoggleBoard(t *testing.T) {
	tests := []struct {
		board      [][]rune
		words, out []string
	}{
		{
			[][]rune{
				{'t', 'h', 'i', 's', 'i', 's', 'a'},
				{'s', 'i', 'm', 'p', 'l', 'e', 'x'},
				{'b', 'x', 'x', 'x', 'x', 'e', 'b'},
				{'x', 'o', 'g', 'g', 'l', 'x', 'o'},
				{'x', 'x', 'x', 'D', 'T', 'r', 'a'},
				{'R', 'E', 'P', 'E', 'A', 'd', 'x'},
				{'x', 'x', 'x', 'x', 'x', 'x', 'x'},
				{'N', 'O', 'T', 'R', 'E', '-', 'P'},
				{'x', 'x', 'D', 'E', 'T', 'A', 'E'},
			},
			[]string{
				"this", "is", "not", "a", "simple", "boggle", "board", "test", "REPEATED", "NOTRE-PEATED",
			},
			[]string{
				"board", "NOTRE-PEATED", "this", "is", "simple", "a", "boggle",
			},
		},
	}

	for i, test := range tests {
		out := BoggleBoard(test.board, test.words)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %#v got %#v", i, test.out, out)
		}
	}
}

package strings

import (
	"reflect"
	"testing"
)

func TestMinimumCharactersForWords(t *testing.T) {
	tests := []struct {
		in, out []string
	}{
		{
			[]string{"yyour", "you", "or", "yo"},
			[]string{"u", "r", "y", "y", "o"},
		},
	}
	for i, test := range tests {
		out := MinimumCharactersForWords(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %q got %q", i, test.out, out)
		}
	}

}

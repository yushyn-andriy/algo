package strings

import (
	"reflect"
	"testing"
)

func TestReverseWordsInString(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			"4  whit",
			"whit  4",
		},
		{
			"Algo is the best",
			"best the is Algo",
		},
	}
	for i, test := range tests {
		out := ReverseWordsInString(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %q got %q", i, test.out, out)
		}
	}

}

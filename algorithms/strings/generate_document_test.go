package strings_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/strings"
)

func TestSameBsts(t *testing.T) {
	tests := []struct {
		characters string
		document   string
		out        bool
	}{
		{
			"Bste!hetsi ogEAxpelrt x ",
			"AlgoExpert is the Best!",
			true,
		},
		{
			"helloworld ",
			"hello wOrld",
			false,
		},
	}

	for i, test := range tests {
		out := strings.GenerateDocument(test.characters, test.document)
		if !reflect.DeepEqual(test.out, out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}

	}
}

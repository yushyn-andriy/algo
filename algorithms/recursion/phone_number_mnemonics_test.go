package recursion

import (
	"reflect"
	"testing"
)

func TestPhoneNumberMnemonics(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			"1905",
			[]string{
				"1w0j",
				"1w0k",
				"1w0l",
				"1x0j",
				"1x0k",
				"1x0l",
				"1y0j",
				"1y0k",
				"1y0l",
				"1z0j",
				"1z0k",
				"1z0l",
			},
		},
	}

	for i, test := range tests {
		out := PhoneNumberMnemonics(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}

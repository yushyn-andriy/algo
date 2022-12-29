package strings

import (
	"reflect"
	"testing"
)

func TestValidIPAddress(t *testing.T) {
	tests := []struct {
		str       string
		addresses []string
	}{
		{
			"1921680",
			[]string{
				"1.9.216.80",
				"1.92.16.80",
				"1.92.168.0",
				"19.2.16.80",
				"19.2.168.0",
				"19.21.6.80",
				"19.21.68.0",
				"19.216.8.0",
				"192.1.6.80",
				"192.1.68.0",
				"192.16.8.0",
			},
		},
	}
	for i, test := range tests {
		out := ValidIPAddresses(test.str)
		if !reflect.DeepEqual(out, test.addresses) {
			t.Errorf("test(%d) expected %v got %v", i, test.addresses, out)
		}
	}

}

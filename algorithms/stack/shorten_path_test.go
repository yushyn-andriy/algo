package stack

import (
	"reflect"
	"testing"
)

func TestShortenPath(t *testing.T) {
	tests := []struct {
		path, out string
	}{
		{
			"/foo/../test/../test/../foo//bar/./baz",
			"/foo/bar/baz",
		},
	}

	for i, test := range tests {
		out := ShortenPath(test.path)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %q got %q", i, test.out, out)
		}
	}
}

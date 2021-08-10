package strings

import (
	"reflect"
	"testing"
)

func TestPatternMatcher(t *testing.T) {
	tests := []struct {
		pattern, str string
		out          []string
	}{
		{"xyx", "stolstulstol", []string{"stol", "stul"}},
		{"xxxyy", "moremoremoresolenoesolenoe", []string{"more", "solenoe"}},
		{"xxyxy", "andrewandrewsvetlanaandrewsvetlana", []string{"andrew", "svetlana"}},
		{"xxy", "gogope", []string{"go", "pe"}},
		{"xxyxxy", "gogopowerrangergogopowerranger", []string{"go", "powerranger"}},
	}
	for i, test := range tests {
		out := PatternMatcher(test.pattern, test.str)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %q got %q", i, test.out, out)
		}
	}

}

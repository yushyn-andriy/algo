package strings

import "testing"

func TestUnderscorifySubstring(t *testing.T) {
	tests := []struct {
		str, substring, out string
	}{
		{
			"testthis is a testtest to see if testestest it works",
			"test",
			"_test_this is a _testtest_ to see if _testestest_ it works",
		},
		{
			"ttttttttttttttbtttttctatawtatttttastvb",
			"ttt",
			"_tttttttttttttt_b_ttttt_ctatawta_ttttt_astvb",
		},

	}

	for i, test := range tests {
		out := UnderscorifySubstring(test.str, test.substring)
		if out != test.out {
			t.Errorf("test(%d) expected %q got %q", i, test.out, out)
		}

	}
}

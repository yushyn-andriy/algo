package strings

import "testing"

func TestCaesarCipherEncryptort(t *testing.T) {
	tests := []struct {
		in  string // input string
		key int
		out string
	}{
		{
			"xyz",
			2,
			"zab",
		},
	}

	for i, test := range tests {
		out := CaesarCipherEncryptor(test.in, test.key)
		if out != test.out {
			t.Errorf("test(%d) expected entropy is %v but got %v", i, test.out, out)
		}
	}
}

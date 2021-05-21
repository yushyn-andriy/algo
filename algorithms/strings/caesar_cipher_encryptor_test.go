package strings_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/strings"
)

func TestCaesarCipherEncrypt(t *testing.T) {
	tests := []struct {
		toencrypt, encrypted string
		key                    int
	}{
		{"xyz", "z a", 2},
		{"xyz", "yz ", 1},
		{"andrew and sve", "cpftgybcpfbuxg", 2},
	}

	for i, test := range tests {
		encrypted, err := strings.CaesarCipherEncrypt(test.toencrypt, test.key)
		if err != nil {
			t.Errorf("test(%d), Expected error is nil got %v", i, err)
			return
		}

		if encrypted != test.encrypted {
			t.Errorf("test(%d), Expected encrypted '%s', got '%s'", i, test.encrypted, encrypted)
			return
		}
	}
}

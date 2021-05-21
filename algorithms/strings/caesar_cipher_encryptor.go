package strings

import (
	"fmt"
)

const (
	ASCII = "abcdefghijklmnopqrstuvwxyz "
)

var (
	ASCIIToNum map[rune]int
	NumToASCII map[int]rune
)

func init() {
	ASCIIToNum = make(map[rune]int, len(ASCII))
	NumToASCII = make(map[int]rune, len(ASCII))
	for i, r := range ASCII {
		ASCIIToNum[r] = i
		NumToASCII[i] = r
	}
}

func CaesarCipherEncrypt(s string, key int) (string, error) {
	encrypted := ""
	for _, r := range s {
		num, ok := ASCIIToNum[r]
		if !ok {
			return "", fmt.Errorf("Rune out of range ASCII symbols '%v'", r)
		}
		shiftedNum := (num + key) % len(ASCII)
		shiftedRune, ok := NumToASCII[shiftedNum]
		if !ok {
			return "", fmt.Errorf("Shifted num out of range '%v'", shiftedNum)
		}

		encrypted += string(shiftedRune)
	}

	return encrypted, nil
}

func CaesarCipherDecrypt(s string, n int) (string, error) {
	return "", nil
}

package strings

import "strings"

func CaesarCipherEncryptor(str string, key int) string {
	runes := []rune(str)
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i, char := range runes {
		index := strings.Index(alphabet, string(char))
		if index == -1 {
			return ""
		}
		newindex := (index + key) % 26
		runes[i] = rune(alphabet[newindex])
	}

	return string(runes)
}
